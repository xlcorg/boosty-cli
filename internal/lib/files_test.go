package lib

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFile_CreateDelete(t *testing.T) {
	t.Run("just works", func(t *testing.T) {
		file := File("empty.txt")

		err := file.CreateFile()
		assert.NoError(t, err)
		assert.True(t, file.Exists())

		err = file.DeleteFile()
		assert.NoError(t, err)
		assert.False(t, file.Exists())
	})

	t.Run("for write", func(t *testing.T) {
		file := File("empty.txt")
		w, err := file.CreateFileForWrite()
		assert.NoError(t, err)
		defer w.Close()
		defer file.DeleteFile()

		_, _ = w.Write([]byte("test"))
		text, err := file.ReadAllText()
		assert.NoError(t, err)

		assert.Equal(t, "test", text)
	})
}

func TestReadLines(t *testing.T) {
	t.Run("read all lines", func(t *testing.T) {
		file := File("test.txt")
		defer file.DeleteFile()

		err := file.WriteAllText("Random text here\n123")
		assert.NoError(t, err)

		lines, err := file.ReadAllLines()
		assert.NoError(t, err)
		assert.NotEmpty(t, lines)
	})

	t.Run("read into string", func(t *testing.T) {
		file := File("test.txt")
		defer file.DeleteFile()

		err := file.WriteAllText("Random text here")
		assert.NoError(t, err)

		text, err := file.ReadAllText()
		assert.NoError(t, err)
		assert.NotEmpty(t, text)
	})

	t.Run("append path", func(t *testing.T) {
		file := File("/Users/xlcode")

		actual := file.AppendPath("code", "go", "playground")

		expected := filepath.Join("/Users/xlcode/code/go/playground")
		assert.Equal(t, expected, actual.String())
	})

	t.Run("exists", func(t *testing.T) {
		file := File("test.txt")
		err := file.WriteAllText("Random text here")
		defer file.DeleteFile()
		assert.NoError(t, err)

		assert.True(t, file.Exists())

		file = "/Users/xlcode/code/go/playground/.pfff"
		assert.False(t, file.Exists())

		// TODO: Create Dir
		//file = "/Users/xlcode/code/go/playground"
		//assert.False(t, file.Exists())
	})

	t.Run("size", func(t *testing.T) {
		file := File("test.txt")
		defer file.DeleteFile()

		err := file.WriteAllText("123")
		assert.NoError(t, err)

		size, err := file.GetFileSize()
		assert.NoError(t, err)
		assert.True(t, size > 0)

		// TODO: Create Dir
		//file = File("/Users/xlcode/code/go/playground")
		//assert.NoError(t, err)
		//size, err = file.GetFileSize()
		//assert.True(t, size > 0)
	})
}

func TestFile_Write(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		file := File("test.txt")
		defer file.DeleteFile()

		err := file.WriteAllText("")
		assert.NoError(t, err)
		assert.True(t, file.Exists())

		empty, err := file.IsEmptyFile()
		assert.NoError(t, err)
		assert.True(t, empty)
	})

	t.Run("append", func(t *testing.T) {
		file := File("test.txt")
		defer file.DeleteFile()

		err := file.AppendAllText("hello")
		assert.NoError(t, err)
		assert.True(t, file.Exists())

		text, err := file.ReadAllText()
		assert.NoError(t, err)
		assert.Equal(t, "hello", text)

		err = file.AppendAllText("world")
		assert.NoError(t, err)
		assert.True(t, file.Exists())

		text, err = file.ReadAllText()
		assert.NoError(t, err)
		assert.Equal(t, "helloworld", text)
	})

	t.Run("just works", func(t *testing.T) {
		expected := "Hello, World!"

		file := File("test.txt")
		defer file.DeleteFile()

		err := file.WriteAllText(expected)
		assert.NoError(t, err)
		assert.True(t, file.Exists())

		empty, err := file.IsEmptyFile()
		assert.NoError(t, err)
		assert.False(t, empty)

		actual, err := file.ReadAllText()
		assert.NoError(t, err)
		assert.Equal(t, expected, actual)
	})
}

func TestFile_GetFileHash(t *testing.T) {
	t.Run("Sha256", func(t *testing.T) {
		file := File("test.txt")
		content := "Hello, World!"

		err := file.WriteAllText(content)
		assert.NoError(t, err)
		defer file.DeleteFile()

		hash, err := file.GetFileSha256()
		assert.NoError(t, err)

		assert.Equal(t, "dffd6021bb2bd5b0af676290809ec3a53191dd81c7f70a4b28688a362182986f", hash)
	})

	t.Run("md5", func(t *testing.T) {
		file := File("test.txt")
		content := "Hello, World!"

		err := file.WriteAllText(content)
		assert.NoError(t, err)
		defer file.DeleteFile()

		hash, err := file.GetFileMd5()
		assert.NoError(t, err)

		assert.Equal(t, "65a8e27d8879283831b664bd8b7f0ad4", hash)
	})

}
