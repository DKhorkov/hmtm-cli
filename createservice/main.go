package createservice

import (
	"hmtm-cli/createservice/content"
	"hmtm-cli/createservice/structure"
)

func Init(name string) error {
	err := createFolders(structure.New(name))
	if err != nil {
		return err
	}

	err = createFiles(content.New(name))

	return nil
}
