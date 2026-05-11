package storage

import (
	"demo/struct/bins"
	"encoding/json"
	"fmt"
	"io/fs"
	"os"
)

// Storage defines the interface for storage operations
type Storage interface {
	SaveBin(bin *bins.Bin) error
	SaveBinList(binList *bins.BinList) error
	LoadBinList() (*bins.BinList, error)
}

// FileSystem defines the interface for file operations
type FileSystem interface {
	WriteFile(name string, data []byte, perm fs.FileMode) error
	ReadFile(name string) ([]byte, error)
}

// FileStorage implements Storage using the filesystem
type FileStorage struct {
	fs FileSystem
}

// NewFileStorage creates a new FileStorage instance
func NewFileStorage(fs FileSystem) *FileStorage {
	return &FileStorage{fs: fs}
}

// SaveBin implements Storage interface
func (s *FileStorage) SaveBin(bin *bins.Bin) error {
	data, err := json.Marshal(bin)
	if err != nil {
		return fmt.Errorf("marshal bin: %w", err)
	}
	return s.fs.WriteFile("bin.json", data, 0644)
}

// SaveBinList implements Storage interface
func (s *FileStorage) SaveBinList(binList *bins.BinList) error {
	data, err := json.Marshal(binList)
	if err != nil {
		return fmt.Errorf("marshal bin list: %w", err)
	}
	return s.fs.WriteFile("binList.json", data, 0644)
}

// LoadBinList implements Storage interface
func (s *FileStorage) LoadBinList() (*bins.BinList, error) {
	data, err := s.fs.ReadFile("binList.json")
	if err != nil {
		return nil, fmt.Errorf("read bin list: %w", err)
	}

	var binList bins.BinList
	if err := json.Unmarshal(data, &binList); err != nil {
		return nil, fmt.Errorf("unmarshal bin list: %w", err)
	}
	return &binList, nil
}

// OSFileSystem implements FileSystem using the OS package
type OSFileSystem struct{}

func (OSFileSystem) WriteFile(name string, data []byte, perm fs.FileMode) error {
	return os.WriteFile(name, data, perm)
}

func (OSFileSystem) ReadFile(name string) ([]byte, error) {
	return os.ReadFile(name)
}
