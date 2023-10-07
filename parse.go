package assistant

import (
	"embed"
	"encoding/json"
	"fmt"
)

// ParseTests - parse test data from JSON.
func ParseTests[T any](raw []byte) (tests []T, err error) {
	if len(raw) == 0 {
		return nil, ErrNotFoundTestData
	}

	if err = json.Unmarshal(raw, &tests); err != nil {
		return nil, fmt.Errorf("json.Unmarshal: %w", err)
	}

	if len(tests) == 0 {
		return nil, ErrNotFoundTests
	}

	return tests, nil
}

// MustLoadTests - must parse test data from JSON.
func MustLoadTests[T any](raw []byte) (tests []T) {
	var err error

	if tests, err = ParseTests[T](raw); err != nil {
		panic(err)
	}

	return tests
}

// ParseTestsFromFS - parse test data from JSON by path.
func ParseTestsFromFS[T any](fs embed.FS, path string) (tests []T, err error) {
	file, err := fs.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("fs.ReadFile: %w", err)
	}

	tests, err = ParseTests[T](file)
	if err != nil {
		return nil, err
	}

	return tests, nil
}

// MustLoadTestsFromFS - must parse test data from JSON by path.
func MustLoadTestsFromFS[T any](fs embed.FS, path string) (tests []T) {
	var err error

	tests, err = ParseTestsFromFS[T](fs, path)
	if err != nil {
		panic(err)
	}

	return tests
}
