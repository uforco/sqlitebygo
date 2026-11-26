package readsqlfile

import (
	"errors"
	"os"
	"path/filepath"
	"strings"
)

func ReadSqlFile(path string) (string, error) {
	if strings.TrimSpace(path) == "" {
		return "", errors.New("path not defined")
	}

	rootPath := "sql"
	path = filepath.Clean(path)

	// Avoid adding the root path twice.
	if !filepath.IsAbs(path) {
		path = strings.TrimPrefix(path, "."+string(os.PathSeparator))

		if path != rootPath &&
			!strings.HasPrefix(path, rootPath+string(os.PathSeparator)) {
			path = filepath.Join(rootPath, path)
		}
	}

	//if not exist,
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return "", err
	}

	file, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(file), nil
}
