package file

import (
	"github.com/RepSklvska/socofile_back/glb"
	"io/ioutil"
	"os"
	"strings"
)

type File struct {
	Filename string
	Size     string
	ModDate  string
	DLLink   string
}

//var c glb.Config = glb.Read()

func ListDir(type_ string, user string) []map[string]string {
	userDir := c.FileStore + "/" + glb.MD5Sum(user)
	glb.ReplaceRept(&userDir, "/")
	userDir = strings.TrimRight(userDir, "/")
	
	//fmt.Println(userDir)
	if glb.FileExist(userDir) {
		os.Remove(userDir)
	}
	if !glb.DirExist(userDir) {
		os.MkdirAll(userDir+"/plain", 0775)
		os.MkdirAll(userDir+"/secret/cache", 0775)
		os.MkdirAll(userDir+"/plain_share", 0775)
		os.MkdirAll(userDir+"/secret_share", 0775)
	}
	
	var (
		fileList  []map[string]string
		dirToRead string
	)
	switch type_ {
	case "plain":
		dirToRead = userDir + "/plain"
		list, _ := ioutil.ReadDir(dirToRead)
		for _, file := range list {
			fileList = append(fileList, map[string]string{
				"name": file.Name(),
				"size": glb.ComputeSize(file.Size()),
				"date": file.ModTime().Format("2006-01-02 15:04:05"),
				"link": "",
			})
		}
		return fileList
	case "plain_share":
		dirToRead = userDir + "/plain_share"
		list, _ := ioutil.ReadDir(dirToRead)
		for _, file := range list {
			fileList = append(fileList, map[string]string{
				"name": file.Name(),
				"size": glb.ComputeSize(file.Size()),
				"date": file.ModTime().Format("2006-01-02 15:04:05"),
				"link": "",
			})
		}
		return fileList
	case "secret":
		dirToRead = userDir + "/secret"
		list, _ := ioutil.ReadDir(dirToRead)
		for _, file := range list {
			fileList = append(fileList, map[string]string{
				"name": file.Name(),
				"size": glb.ComputeSize(file.Size()),
				"date": file.ModTime().Format("2006-01-02 15:04:05"),
				"link": "",
			})
		}
		return fileList
	
	case "secret_share":
		return nil
	}
	return nil
}
