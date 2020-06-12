package file

import (
	"database/sql"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/RepSklvska/socofile_back/glb"
	"github.com/RepSklvska/socofile_back/user"
	"github.com/gin-gonic/gin"
	"io/ioutil"
)

// Handle users' uploads

//var c glb.Config = glb.Read()

func HandleUpload(type_ string, ctx *gin.Context, db *sql.DB) bool {
	file, err := ctx.FormFile("name")
	if err != nil {
		return false
	}
	
	switch type_ {
	case "plain":
		dest := c.FileStore + "/" + glb.MD5Sum(user.GetUser(ctx)) + "/plain"
		if err := ctx.SaveUploadedFile(file, dest); err != nil {
			return true
		} else {
			return false
		}
	case "secret":
		dest := c.FileStore + "/" + glb.MD5Sum(user.GetUser(ctx)) + "/secret/cache"
		err := ctx.SaveUploadedFile(file, dest)
		if err != nil {
			return false
		}
		// 放到cache里面进行加密然后删掉原文件
		
		// 插入数据
	}
	return true
}

func HandleUpload2(jsonBody []byte, db *sql.DB) error {
	body := jsonBody
	
	type_ := glb.GetField(body, "type")
	username := glb.GetField(body, "username")
	userDir := c.FileStore + "/" + glb.MD5Sum(username)
	
	mapData := make(map[string]string)
	json.Unmarshal(body, &mapData)
	//fmt.Println(mapData)
	
	switch type_ {
	case "plain":
		filePath := userDir + "/plain/" + glb.GetField(body, "filename")
		fmt.Println("file saved to:", filePath)
		fileBytes, err := base64.StdEncoding.DecodeString(glb.GetField(body, "base64ed"))
		if err != nil {
			return err
		}
		if err := ioutil.WriteFile(filePath, fileBytes, 0666); err != nil {
			return err
		}
	case "secret":
		filePath := userDir + "/secret/" + glb.GetField(body, "filename")
		_=filePath
	}
	
	return nil
}
