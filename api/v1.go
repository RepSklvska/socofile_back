package api

import (
	"database/sql"
	"fmt"
	"github.com/RepSklvska/socofile_back/file"
	"github.com/RepSklvska/socofile_back/glb"
	"github.com/RepSklvska/socofile_back/user"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

func APIv1(context *gin.Context, database *sql.DB) {
	var (
		response gin.H = gin.H{}
		err      error
	)
	body, _ := ioutil.ReadAll(context.Request.Body)
	fmt.Println(context.GetHeader("content-type"))
	fmt.Println(string(body))
	switch glb.Method(body) {
	case "login":
		response, err = user.Login(body, database, context)
		if err != nil {
			fmt.Println(err)
		}
	case "register":
		response, err = user.Register(body, database, context)
		if err != nil {
			fmt.Println(err)
		}
	case "logout":
		user.Logout(context)
		response["loggedOut"] = true
	case "checkSession":
		response["loggedIn"] = user.GetSession(context)
		response["loggedAs"] = fmt.Sprint(user.GetUser(context))
	case "getFileList":
		//fmt.Println(string(body))
		response["fileList"] = file.ListDir(
			glb.GetField(body, "type"),
			glb.GetField(body, "username"),
		)
	//case "upload":
	//	_, _ = context.FormFile("file")
	//	response["uploaded"] = file.HandleUpload(
	//		glb.GetField(body, "type"),
	//		context,
	//		database,
	//	)
	case "download":
	
	case "share":
		if err := file.Share(
			glb.GetField(body, "filename"),
			glb.GetField(body, "type"),
			glb.GetField(body, "username"),
		); err != nil {
			response["error"] = err
		} else {
			response["ok"] = "ok"
		}
	case "unshare":
		if err := file.Unshare(
			glb.GetField(body, "filename"),
			glb.GetField(body, "type"),
			glb.GetField(body, "username"),
		); err != nil {
			response["error"] = err
		} else {
			response["ok"] = "ok"
		}
	case "delete":
	
	case "upload":
		if err := file.HandleUpload2(body, database); err != nil {
			response["error"] = err
		} else {
			response["ok"] = "ok"
		}
	}
	
	//if cookie, err := context.Cookie("user"); err != nil {
	//	context.JSONP(http.StatusOK, gin.H{
	//		"error":         "not logged in",
	//		"cookieContent": cookie,
	//	})
	//}
	//response["loggedIn"] = string(user.GetUser(context))
	
	//fmt.Println(user.GetUser(context))
	//fmt.Println(user.GetSession(context))
	fmt.Println("response:", response)
	context.JSONP(http.StatusOK, response)
}

func APIv1Upload(context *gin.Context, database *sql.DB) {
	body, _ := ioutil.ReadAll(context.Request.Body)
	fmt.Println(string(body))
	_, _, _ = context.Request.FormFile("file")
	fmt.Println(context.GetHeader("content-type"))
	
}
