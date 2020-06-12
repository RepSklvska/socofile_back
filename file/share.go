package file

import (
	"github.com/RepSklvska/socofile_back/glb"
	"os"
)

func Share(filename string, type_ string, user string) error {
	//os.Symlink()
	switch type_ {
	case "plain":
		userDir := c.FileStore + "/" + glb.MD5Sum(user)
		filePath := userDir + "/plain/" + filename
		linkTo := userDir + "/plain_share/" + filename
		if err := os.Symlink(filePath, linkTo); err != nil {
			return err
		}
	case "secret":
	
	}
	return nil
}

func Unshare(filename string, type_ string, user string) error {
	switch type_ {
	case "plain_share":
		linkPath := c.FileStore + "/" + glb.MD5Sum(user) + "/plain_share/" + filename
		if err := os.Remove(linkPath); err != nil {
			return err
		}
	}
	return nil
}
