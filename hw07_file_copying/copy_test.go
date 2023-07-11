package main

import (
	"errors"
	"os"
	"testing"
)

func TestCopy(t *testing.T) {
	fromPath := "testdata/input.txt"
	toPath := "testdata/output.txt"
	t.Run("TestCopyWithLimit0", func(t *testing.T) {
		offset := int64(0)
		limit := int64(0)
		expectedError := (error)(nil)

		err := Copy(fromPath, toPath, offset, limit)

		if !errors.Is(err, expectedError) {
			t.Errorf("Expected error: %v, but got: %v", expectedError, err)
		}
		inputContent, err := os.ReadFile(fromPath)
		if err != nil {
			t.Errorf("Error reading input file: %v", err)
		}
		outputContent, err := os.ReadFile(toPath)
		if err != nil {
			t.Errorf("Error reading output file: %v", err)
			return
		}
		if string(outputContent) != string(inputContent) {
			t.Errorf("Expected copied content: %s, but got: %s", string(inputContent), string(outputContent))
		}
		defer os.Remove(toPath)
	})

	t.Run("TestCopyWithLimit0", func(t *testing.T) {
		offset := int64(0)
		limit := int64(10)
		expectedError := (error)(nil)

		err := Copy(fromPath, toPath, offset, limit)
		if !errors.Is(err, expectedError) {
			t.Errorf("Expected error: %v, but got: %v", expectedError, err)
		}
		inputContent, err := os.ReadFile(fromPath)
		if err != nil {
			t.Errorf("Error reading input file: %v", err)
			return
		}
		outputContent, err := os.ReadFile(toPath)
		if err != nil {
			t.Errorf("Error reading output file: %v", err)
			return
		}
		expectedContent := inputContent[:limit]
		if string(outputContent) != string(expectedContent) {
			t.Errorf("Expected copied content: %s, but got: %s", string(expectedContent), string(outputContent))
		}
		defer os.Remove(toPath)
	})

	t.Run("TestCopyError", func(t *testing.T) {
		fromPath := "out_offset0_limit0.txt"
		toPath := "out_offset0_limit0.txt"
		err := Copy(fromPath, toPath, 0, 0)
		if !errors.Is(err, ErrCopyToYouself) {
			t.Error("An error was expected")
		}
	})

	t.Run("TestCopyReadingFile", func(t *testing.T) {
		fromPath := "input.txt"
		toPath := "output.txt"
		if err := Copy(fromPath, toPath, 0, 0); err == nil {
			t.Error(err)
		}
	})

	t.Run("TestCopyOffsetOversized", func(t *testing.T) {
		fromPath := "testdata/input.txt"
		toPath := "testdata/output.txt"
		offset := int64(10000)
		limit := int64(0)
		if err := Copy(fromPath, toPath, offset, limit); err != nil {
			t.Error(err)
		}
	})
}
