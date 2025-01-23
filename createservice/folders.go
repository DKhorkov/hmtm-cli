package createservice

import "os"

func createFolder(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err = os.MkdirAll(path, 0755)
		if err != nil {
			return err
		}
	}

	return nil
}

func createFolders(paths []string) error {
	for _, path := range paths {
		err := createFolder(path)
		if err != nil {
			return err
		}
	}

	return nil
}
