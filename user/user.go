package user

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gin-contrib/sessions"
	
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

type User struct {
	Email       string `json:"email"`
	PasswordMD5 string `json:"passwordMD5"`
	Nickname    string `json:"nickname"`
}

func GetSession(ctx *gin.Context) bool {
	//_, err := ctx.Cookie("user")
	//if err != nil {
	//	cookie = "not set"
	//}
	session := sessions.Default(ctx)
	loginUser := session.Get("loginUser")
	fmt.Println("login user:", loginUser)
	return loginUser != nil
}

func GetUser(ctx *gin.Context) string {
	loginUser := sessions.Default(ctx).Get("loginUser")
	switch loginUser.(type) {
	case string:
		return fmt.Sprint(loginUser)
	}
	return ""
}

func Logout(ctx *gin.Context) {
	session := sessions.Default(ctx)
	session.Delete("loginUser")
	session.Save()
	fmt.Println("logout")
}

func Login(jsonBody []byte, db *sql.DB, ctx *gin.Context) (gin.H, error) {
	var (
		user User
		resp gin.H
		err  error = nil
	)
	if err = json.Unmarshal(jsonBody, &user); err != nil {
		resp["error"] = err.Error()
	}
	sqlSelect := fmt.Sprintf(
		`SELECT COUNT(email) email FROM public."user"
	WHERE email='%s' AND password_md5='%s';`,
		user.Email, user.PasswordMD5)
	
	var result string
	if err = db.QueryRow(sqlSelect).Scan(&result); err != nil {
		resp["error"] = err.Error()
	}
	if result == "1" {
		//给cookie
		fmt.Println("password correct, set session")
		//ctx.SetCookie("user", user.Email, 3600, "/", "localhost", false, true)
		session := sessions.Default(ctx)
		session.Set("loginUser", user.Email)
		session.Save()
	} else {
		fmt.Println("password incorrect")
		resp["error"] = "user unexist or duplicated"
	}
	
	return resp, err
}

func Register(jsonBody []byte, db *sql.DB, ctx *gin.Context) (gin.H, error) {
	var (
		user User
		resp gin.H
		err  error = nil
	)
	if err = json.Unmarshal(jsonBody, &user); err != nil {
		resp["error"] = err.Error()
	}
	sqlInsert := fmt.Sprintf(
		`INSERT INTO public.user_detail(
	email, password_md5, nickname)
	VALUES ('%s', '%s', '%s');`,
		user.Email, user.PasswordMD5, user.Nickname)
	
	_, err = db.Exec(sqlInsert)
	if err != nil {
		resp["error"] = err.Error()
		
		
	}
	
	return resp, err
}
