package hw04lrucache

import (
	"math/rand"
	"strconv"
	"sync"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCache(t *testing.T) {
	t.Run("empty cache", func(t *testing.T) {
		c := NewCache(10)

		_, ok := c.Get("aaa")
		require.False(t, ok)

		_, ok = c.Get("bbb")
		require.False(t, ok)
	})

	t.Run("simple", func(t *testing.T) {
		c := NewCache(5)
		wasInCache := c.Set("aaa", 100)
		require.Falsef(t, wasInCache, "1")

		wasInCache = c.Set("bbb", 200)
		require.Falsef(t, wasInCache, "2")

		val, ok := c.Get("aaa")
		require.True(t, ok)
		require.Equalf(t, 100, val, "3")

		val, ok = c.Get("bbb")
		require.True(t, ok)
		require.Equalf(t, 200, val, "4")

		wasInCache = c.Set("aaa", 300)
		require.Truef(t, wasInCache, "5")

		val, ok = c.Get("aaa")
		require.True(t, ok)
		require.Equalf(t, 300, val, "6")

		val, ok = c.Get("ccc")
		require.Falsef(t, ok, "7")
		require.Nil(t, val, "8")
	})

	t.Run("purge logic", func(t *testing.T) {
		c := NewCache(2)
		c.Set("aaa", 100)
		c.Set("bbb", 200)
		c.Set("ccc", 300)

		val, ok := c.Get("aaa")
		require.Nil(t, val)
		require.False(t, ok)

		val, ok = c.Get("bbb")
		require.NotNil(t, val)
		require.True(t, ok)

		val, ok = c.Get("ccc")
		require.NotNil(t, val)
		require.True(t, ok)
	})
	t.Run("least recently used", func(t *testing.T) {
		c := NewCache(3)
		c.Set("aaa", 100)
		c.Set("bbb", 200)
		c.Set("ccc", 300)

		// Обращаемся к "aaa", чтобы обновить его положение в LRU кэше
		c.Get("aaa")

		c.Set("ddd", 400)

		// "bbb" должен быть вытолкнут, так как он был затронут наименее недавно
		_, ok := c.Get("bbb")
		require.False(t, ok)

		val, ok := c.Get("aaa")
		require.NotNil(t, val)
		require.True(t, ok)

		val, ok = c.Get("ccc")
		require.NotNil(t, val)
		require.True(t, ok)

		val, ok = c.Get("ddd")
		require.NotNil(t, val)
		require.True(t, ok)
	})
	t.Run("queue size limit", func(t *testing.T) {
		c := NewCache(3)
		c.Set("aaa", 100)
		c.Set("bbb", 200)
		c.Set("ccc", 300)
		c.Set("ddd", 400)

		// "aaa" должен быть вытолкнут, так как кэш заполнен до максимальной емкости
		_, ok := c.Get("aaa")
		require.False(t, ok)

		val, ok := c.Get("bbb")
		require.NotNil(t, val)
		require.True(t, ok)

		val, ok = c.Get("ccc")
		require.NotNil(t, val)
		require.True(t, ok)

		val, ok = c.Get("ddd")
		require.NotNil(t, val)
		require.True(t, ok)
	})
	t.Run("cache capacity", func(t *testing.T) {
		c := NewCache(3)
		c.Set("100", 1)
		c.Set("200", 2)
		c.Set("300", 3)
		ok := c.Set("400", 4)
		require.False(t, ok)

		val, ok := c.Get("100")
		require.Nil(t, val)
		require.False(t, ok)

		val, ok = c.Get("200")
		require.NotNil(t, val)
		require.True(t, ok)

		val, ok = c.Get("300")
		require.NotNil(t, val)
		require.True(t, ok)

		val, ok = c.Get("400")
		require.NotNil(t, val)
		require.True(t, ok)
	})
}

func TestCacheMultithreading(t *testing.T) {
	t.Skip() // Remove me if task with asterisk completed.

	c := NewCache(10)
	wg := &sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		for i := 0; i < 1_000_000; i++ {
			c.Set(Key(strconv.Itoa(i)), i)
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 1_000_000; i++ {
			c.Get(Key(strconv.Itoa(rand.Intn(1_000_000))))
		}
	}()

	wg.Wait()
}
