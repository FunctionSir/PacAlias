/*
 * @Author: FunctionSir
 * @License: AGPLv3
 * @Date: 2023-10-27 17:44:49
 * @LastEditTime: 2023-10-28 15:09:14
 * @LastEditors: FunctionSir
 * @Description: PacAlias Version 0.1-alpha (HitoriGotoh)
 * @FilePath: /PacAlias/pacalias.go
 */
package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/go-ini/ini"
)

var ConfPath string = "pacalias.conf"
var Conf *ini.File = nil
var NoLog bool = false
var ListenAddr string = "127.0.0.1"
var ListenPort string = "2090"

func args_parser() {
	if len(os.Args) >= 2 {
		for i := 1; i < len(os.Args); i++ {
			switch os.Args[i] {
			case "-c", "--conf":
				ConfPath = os.Args[i+1]
				i++
			case "-p", "--port":
				ListenPort = os.Args[i+1]
				i++
			case "-a", "--addr":
				ListenAddr = os.Args[i+1]
				i++
			case "-q", "--quiet":
				NoLog = true
				i++
			}
		}
	}
}

func http_handler(w http.ResponseWriter, r *http.Request) {
	repo := r.URL.Query().Get("repo")
	arch := r.URL.Query().Get("arch")
	file := r.URL.Query().Get("file")
	if !NoLog {
		fmt.Println(time.Now().String() + " <- repo=" + repo + ",arch=" + arch + ",file=" + file[1:])
	}
	dest := Conf.Section(repo).Key("Dest").String()
	if file == "/"+repo+".db" {
		file = strings.ReplaceAll(file, repo, Conf.Section(repo).Key("OName").String())
	}
	location := strings.ReplaceAll(dest, "$arch", arch) + file
	w.Header().Set("Location", location)
	w.WriteHeader(301)
	if !NoLog {
		fmt.Println(time.Now().String() + " -> " + location)
	}
}

func welcome() {
	if !NoLog {
		fmt.Println("PacAlias Ver 0.1-alpha (HitoriGotoh)")
		fmt.Println("This is a free (libre) software under AGPLv3")
		fmt.Println("Conf File: " + ConfPath)
		fmt.Println("Listen Addr: " + ListenAddr)
		fmt.Println("Listen Port: " + ListenPort)
	}
}

func main() {
	args_parser()
	welcome()
	var err error
	Conf, err = ini.Load(ConfPath)
	if err != nil {
		panic(err)
	}
	http.HandleFunc("/", http_handler)
	http.ListenAndServe(ListenAddr+":"+ListenPort, nil)
}
