package storage

import (
	. "boosty/internal/lib"
	"fmt"
	"io"
	"strings"
)

type Storage interface {
	Get(key string) string
	Set(key string, value string)
	Delete(key string)
}

type StorageCloser interface {
	Storage
	io.Closer
}

type storage struct {
	file  File
	items map[string]string
}

func New(settingPath string) (StorageCloser, error) {
	file := File(settingPath).ExpandEnv()
	items := make(map[string]string)

	if file.Exists() {
		lines, err := file.ReadAllLines()
		if err != nil {
			return nil, fmt.Errorf("ReadAllLines(): %v", err)
		}
		for _, line := range lines {
			parts := strings.SplitN(line, ":", 2)
			key := strings.TrimSpace(parts[0])
			val := strings.TrimSpace(parts[1])
			items[key] = val
		}
	}

	return &storage{
		file:  file,
		items: items,
	}, nil
}

func (s *storage) Close() error {
	err := s.file.CreateDirectoryIfNotExist()
	if err != nil {
		return fmt.Errorf("CreateDirectoryIfNotExist(): %v", err)
	}

	file, err := s.file.CreateFileForWrite()
	if err != nil {
		return fmt.Errorf("failed to create file: %v", err)
	}
	defer file.Close()

	for k, v := range s.items {
		fmt.Fprintf(file, "%s: %s\n", k, v)
	}

	return nil
}

func (s *storage) Get(key string) string {
	if val, ok := s.items[key]; ok {
		return val
	}
	return ""
}

func (s *storage) Delete(key string) {
	delete(s.items, key)
}

func (s *storage) Set(key string, value string) {
	s.items[key] = value
}
