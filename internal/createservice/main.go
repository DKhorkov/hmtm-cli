package createservice

import (
	"hmtm-cli/internal/createservice/content"
	"hmtm-cli/internal/createservice/structure"
)

func Create(name string) error {
	err := createFolders(structure.New(name))
	if err != nil {
		return err
	}

	return createFiles(content.New(name))
}
