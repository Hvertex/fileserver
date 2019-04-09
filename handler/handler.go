package handler

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

func UploadFileHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		// 返回index.html
		data, err := ioutil.ReadFile("./static/view/index.html")
		if err != nil {
			fmt.Printf("read file error, %v", err.Error())
			return
		}
		io.WriteString(w, string(data))
	} else {
		// 执行上传

		// 接收文件
		file,head,err := r.FormFile("file")
		if err != nil {
			fmt.Printf("file to get data ,err %s\n", err.Error())
			return
		}

		defer file.Close()

		// 创建文件
		newFile, err := os.Create("/tmp/" + head.Filename)

		if err != nil {
			fmt.Printf("Fail to create file, err:%s\n", err.Error())
			return
		}

		defer newFile.Close()

		// 保存文件内容到新文件
		_, err = io.Copy(newFile, file)
		if err != nil {
			fmt.Printf("Fail to save data into new file, err: %s\n", err.Error())
		}

		// 成功 302跳转
		http.Redirect(w, r, "/file/upload/suc", http.StatusFound)
	}

}

func UploadSuccess(w http.ResponseWriter, r *http.Request)  {
	io.WriteString(w, "upload finished!")
}