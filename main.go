package main

import (
	"database/sql"
	"fmt"
	"github.com/RepSklvska/socofile_back/api"
	"github.com/RepSklvska/socofile_back/glb"
	"github.com/RepSklvska/socofile_back/user"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"net/http"
)

var c glb.Config

func main() {
	//conf, err := ioutil.ReadFile("./settings.json")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//if err := json.Unmarshal(conf, &c); err != nil {
	//	log.Fatal(err)
	//}
	c.Read()
	
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		c.Database.Host, c.Database.Port, c.Database.User, c.Database.Password, c.Database.DBname)
	fmt.Println(psqlInfo)
	
	database, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer database.Close()
	
	router := gin.Default()
	
	router.Use(glb.MyDefaultCORS)
	
	fmt.Println()
	
	store := cookie.NewStore([]byte("loginUser"))
	router.Use(sessions.Sessions("loginSession", store))
	
	router.GET("/", func(context *gin.Context) {
		context.String(http.StatusOK, "Root")
	})
	
	v1 := router.Group("/v1")
	
	v1.GET("/", func(context *gin.Context) {
		if context.Query("check") != "" {
			context.String(http.StatusOK, "CHECK")
			fmt.Println(context.Query("check") == "abc")
		}
		context.String(http.StatusOK, "V1 root")
		
	})
	
	v1.GET("/api", func(context *gin.Context) {
		var response gin.H
		fmt.Println(context.Query("method"))
		//switch context.Query("method") {
		//case "checkSession":
		response["loggedIn"] = user.GetSession(context)
		response["loggedAs"] = fmt.Sprint(user.GetUser(context))
		//case "logout":
		//	user.Logout(context)
		//	response["loggedOut"] = true
		//}
		
		fmt.Println(user.GetUser(context))
		
		context.JSON(http.StatusOK, response)
	})
	
	// Main part
	v1.POST("/api", func(context *gin.Context) {
		api.APIv1(context, database)
	})
	
	v1.POST("/api/upload", func(context *gin.Context) {
		api.APIv1Upload(context, database)
	})
	
	router.Run(":3001")
}
