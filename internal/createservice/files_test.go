package createservice

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

func TestCreateFile(t *testing.T) {
	t.Run("should create file", func(t *testing.T) {
		content := "some information about the file"
		path := "test.txt"
		err := createFile(path, content)
		require.NoError(t, err)

		info, err := os.Stat(path)
		require.NoError(t, err)
		assert.Equal(t, len(content), int(info.Size()))

		if err = os.Remove(path); err != nil {
			t.Fatal(err)
		}
	})

	t.Run("file already exists", func(t *testing.T) {
		path := "test.txt"
		content := "some information about the file"
		file, err := os.Create(path)
		require.NoError(t, err)
		defer func() {
			if err = file.Close(); err != nil {
				t.Fatal(err)
			}
		}()

		_, err = file.WriteString(content)
		require.NoError(t, err)

		fileStat, err := file.Stat()
		require.NoError(t, err)
		assert.Equal(t, len(content), int(fileStat.Size()))

		newContent := "new content"
		err = createFile(path, newContent)
		require.NoError(t, err)

		fileStat, err = file.Stat()
		require.NoError(t, err)
		assert.Equal(t, len(newContent), int(fileStat.Size()))

		if err = os.Remove(path); err != nil {
			t.Fatal(err)
		}
	})
}

func TestCreateFiles(t *testing.T) {
	t.Run("should create files", func(t *testing.T) {
		content := "some information about the file"
		pathsContent := map[string]string{
			"test.txt":         content,
			"another_test.txt": content,
		}

		err := createFiles(pathsContent)
		require.NoError(t, err)

		for path := range pathsContent {
			info, err := os.Stat(path)
			require.NoError(t, err)
			assert.Equal(t, len(content), int(info.Size()))

			if err = os.Remove(path); err != nil {
				t.Fatal(err)
			}
		}
	})
}
