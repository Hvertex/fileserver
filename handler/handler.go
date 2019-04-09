package handler

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func UploadFileHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		// 返回index.html
		data, err := ioutil.ReadFile("./static/view/index.html")
		if err != nil {
			fmt.Printf("read file error, %v", err.Error())
			return
		}
		w.WriteString(data)
	} else {
		// 执行上传
	}
}
