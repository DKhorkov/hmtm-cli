package createservice

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateFolder(t *testing.T) {
	t.Run("should create folder", func(t *testing.T) {
		path := "test_folder"
		err := createFolder(path)
		require.NoError(t, err)

		_, err = os.Stat(path)
		require.NoError(t, err)

		if err = os.RemoveAll(path); err != nil {
			t.Fatal(err)
		}
	})

	t.Run("folder already exists", func(t *testing.T) {
		path := "content"
		_, err := os.Stat(path)
		require.NoError(t, err)

		err = createFolder(path)
		require.NoError(t, err)
	})
}

func TestCreateFolders(t *testing.T) {
	t.Run("should create folders", func(t *testing.T) {
		paths := []string{
			"test_folder",
			"another_test_folder",
		}

		err := createFolders(paths)
		require.NoError(t, err)

		for _, path := range paths {
			_, err = os.Stat(path)
			require.NoError(t, err)

			if err = os.RemoveAll(path); err != nil {
				t.Fatal(err)
			}
		}
	})
}
