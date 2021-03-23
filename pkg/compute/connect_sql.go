package compute

import (
	"Y0shi23/todo/configs"
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

//　設定ファイルへの実行ファイルからの相対パスを指定
const confDir = "./configs/env/"

func ConnectSQL() *sql.DB {
	var err error

	// Get configuration
	conf, err := configs.NewConfig(confDir) // 引数に渡す
	if err != nil {
		panic(err.Error())
	}

	fmt.Println(conf)

	// Open handle to database like normal
	//s := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true", conf.DB.User, conf.DB.Password, conf.DB.Host, conf.DB.Port, conf.DB.Name)
	db, err := sql.Open("mysql", "fumi042:Yoshifumi_01@tcp(127.0.0.1:3306)/TODO?parseTime=true")
	if err != nil {
		panic(err.Error())
	}

	fmt.Println(db)

	return db
}
