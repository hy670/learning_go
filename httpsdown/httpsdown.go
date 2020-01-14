package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

const DEFAULTPORT int = 7788

func GetDir(pwd string) []os.FileInfo {
	files, _ := ioutil.ReadDir(pwd)
	return files

}

func handelGetFile(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() //解析参数，默认是不会解析的
	log.Println("Recv:", r.RemoteAddr)
	pwd, _ := os.Getwd()
	des := pwd + string(os.PathSeparator) + r.URL.Path[1:len(r.URL.Path)]
	desStat, err := os.Stat(des)
	if err != nil {
		log.Println("File Not Exit", des)
		http.NotFoundHandler().ServeHTTP(w, r)
	} else if desStat.IsDir() {
		files := GetDir(des)
		var filleting string
		filleting = "<h1>文件下载清单</h1>"
		for _, file := range files {
			if r.URL.Path == "/" {
				filleting = filleting + "<br><a href=\"" + r.URL.Path[1:len(r.URL.Path)] + "/" + file.Name() + "\">" + file.Name() + "</a>\r\n"
			} else {
				filleting = filleting + "<br><a href=\"/" + r.URL.Path[1:len(r.URL.Path)] + "/" + file.Name() + "\">" + file.Name() + "</a>\r\n"
			}

		}
		log.Println("File Is Dir", des)
		fmt.Fprintln(w, filleting)
	} else {
		fileData, err := ioutil.ReadFile(des)
		if err != nil {
			log.Println("Read File Err:", err.Error())
		} else {
			log.Println("Send File:", des)
			w.Write(fileData)
		}
	}
}

func main() {
	port := flag.Int("p", DEFAULTPORT, "Set The Http Port")
	flag.Parse()
	pwd, _ := os.Getwd()
	log.Printf("Listen On Port:%d pwd:%s\n", *port, pwd)

	http.HandleFunc("/", handelGetFile)
	//err := http.ListenAndServe(":"+strconv.Itoa(*port), nil)
	err := http.ListenAndServeTLS(":7788", "E:\\sourcecode\\go\\src\\learning_go\\httpsdown\\server.crt", "E:\\sourcecode\\go\\src\\learning_go\\httpsdown\\server.key", nil)
	if nil != err {
		log.Fatalln("Get Dir Err", err.Error())
	}
}
