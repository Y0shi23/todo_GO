package main

import (
	"Y0shi23/todo/api/router"

        "net"
        "net/http/fcgi"

	"github.com/fvbock/endless"
)

//　設定ファイルへの実行ファイルからの相対パスを指定
const confDir = "./configs/env/"

func main() {
	a := router.SetupRouter()
        l, _ := net.Listen("tcp", "127.0.0.1:9000")
        
        fcgi.Serve(l, nil)

	endless.ListenAndServe(":8080", a)
}
