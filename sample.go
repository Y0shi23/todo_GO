package main

import (
	"Y0shi23/todo/configure"
	"Y0shi23/todo/models"
	"context"
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/volatiletech/sqlboiler/boil"
)

//　設定ファイルへの実行ファイルからの相対パスを指定
const confDir = "./configure/env/"

func main() {
	var err error

	// Get configuration
	conf, err := configs.NewConfig(confDir) // 引数に渡す
	if err != nil {
		panic(err.Error())
	}

	fmt.Println(conf)

	fmt.Println("Hello World")
	ctx := context.Background()

	// Open handle to database like normal
	s := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true", conf.DB.User, conf.DB.Password, conf.DB.Host, conf.DB.Port, conf.DB.Name)
	db, _ := sql.Open("mysql", s)

	fmt.Println(db)

	boil.SetDB(db)

	r := gin.Default()
	r.LoadHTMLFiles("./web/html/index.html")
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.GET("/todo", func(c *gin.Context) {
		todos, _ := models.Todos().All(ctx, db)
		// if err != nil {
		// 	fmt.Errorf("Get todo error: %v", err)
		// }

		c.HTML(http.StatusOK, "index.html", map[string]interface{}{
			"todo": todos,
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
