package lib

import (
	"bufio"
	"crypto/md5"
	"crypto/sha256"
	"fmt"
	"io"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"
)

type File string

func (f File) String() string {
	return string(f)
}

//////////////////////////////////////////////////////////////////////
// PUBLIC

func (f File) CreateFile() error {
	file, err := os.Create(string(f))
	if err != nil {
		return err
	}
	file.Close()
	return nil
}

func (f File) DeleteFile() error {
	return os.Remove(string(f))
}

func (f File) ReadAllLines() ([]string, error) {
	file, err := os.Open(string(f))
	if err != nil {
		return nil, fmt.Errorf("os.Open(): %v", err)
	}
	defer file.Close()

	lines := make([]string, 0)

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading from file: %v", err)
	}

	return lines, nil
}

func (f File) ReadAllBytes() ([]byte, error) {
	bytes, err := os.ReadFile(string(f))
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

func (f File) ReadAllText() (string, error) {
	bytes, err := f.ReadAllBytes()
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func (f File) WriteAllBytes(bytes []byte) error {
	return os.WriteFile(string(f), bytes, 0777)
}

func (f File) WriteAllText(text string) error {
	return f.WriteAllBytes([]byte(text))
}

func (f File) AppendAllText(text string) error {
	file, err := os.OpenFile(f.String(), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("os.OpenFile(): %v", err)
	}
	defer file.Close()
	if _, err := file.WriteString(text); err != nil {
		return fmt.Errorf("file.WriteString(): %v", err)
	}
	return nil
}

func (f File) CreateFileForWrite() (io.WriteCloser, error) {
	file, err := os.Create(os.ExpandEnv(string(f)))
	if err != nil {
		return nil, fmt.Errorf("os.Create(): %v", err)
	}
	return file, nil
}

func (f File) OpenFileForRead() (io.ReadCloser, error) {
	file, err := os.Open(string(f))
	if err != nil {
		return nil, fmt.Errorf("os.Open(): %v", err)
	}
	return file, nil
}

func (f File) CreateDirectoryIfNotExist() error {
	baseDir := path.Dir(f.String())
	info, err := os.Stat(baseDir)
	if err == nil && info.IsDir() {
		return nil
	}
	return os.MkdirAll(baseDir, 0755)
}

func (f File) ExpandEnv() File {
	filePath := strings.TrimSpace(f.String())
	if strings.HasPrefix(filePath, "~/") {
		dirname, _ := os.UserHomeDir()
		filePath = filepath.Join(dirname, filePath[2:])
	}

	return File(filePath)
}

//////////////////////////////////////////////////////////////////////

func (f File) AppendPath(elems ...string) File {
	path := append([]string{string(f)}, elems...)
	return File(filepath.Join(path...))
}

// Exists determines whether the specified file exists.
func (f File) Exists() bool {
	stat, err := os.Stat(string(f))
	if os.IsNotExist(err) {
		return false
	}
	return !stat.IsDir()
}

func (f File) GetFileSize() (int64, error) {
	stat, err := os.Stat(string(f))
	if os.IsNotExist(err) {
		return 0, err
	}
	return stat.Size(), nil
}

func (f File) GetFileModifiedTime() (time.Time, error) {
	stat, err := os.Stat(string(f))
	if os.IsNotExist(err) {
		return time.Time{}, fmt.Errorf("os.Stat(): %v", err)
	}
	return stat.ModTime(), nil
}

func (f File) IsEmptyFile() (bool, error) {
	size, err := f.GetFileSize()
	if err != nil {
		return false, fmt.Errorf("GetFileSize(): %v", err)
	}
	return size == 0, nil
}

func (f File) GetFileSha256() (string, error) {
	reader, err := f.OpenFileForRead()
	if err != nil {
		return "", fmt.Errorf("OpenFileForRead(): %v", err)
	}
	defer reader.Close()

	hasher := sha256.New()
	if _, err := io.Copy(hasher, reader); err != nil {
		return "", fmt.Errorf("io.Copy(hasher, reader): %v", err)
	}

	return fmt.Sprintf("%x", hasher.Sum(nil)), nil
}

func (f File) GetFileMd5() (string, error) {
	reader, err := f.OpenFileForRead()
	if err != nil {
		return "", fmt.Errorf("OpenFileForRead(): %v", err)
	}
	defer reader.Close()

	hasher := md5.New()
	if _, err := io.Copy(hasher, reader); err != nil {
		return "", fmt.Errorf("io.Copy(hasher, reader): %v", err)
	}

	return fmt.Sprintf("%x", hasher.Sum(nil)), nil
}
