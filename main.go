package main

import (
	"fileserver/handler"
	"net/http"
)

func main() {

	// 路由
	http.HandleFunc("/file/upload", handler.UploadFileHandler)
	http.HandleFunc("/file/upload/suc", handler.UploadSuccess)

	http.ListenAndServe(":8080", nil)
}
