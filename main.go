package main

import (
	"github.com/polevpn/webview"
	"github.com/sqweek/dialog"
	"log"
	"net"
	"net/http"
	"testwebview/server"
)

var debug = "true"

// Starts the API server and returns the host name it's listening on.
func startServer() string {
	addr := make(chan string)
	router := server.NewRouter()

	go func() {
		// Let the OS pick an open port
		listener, err := net.Listen("tcp", "127.0.0.1:0")

		if err != nil {
			log.Fatal(err)
		}

		log.Println("API listening on", listener.Addr().String())

		addr <- listener.Addr().String()

		if err := http.Serve(listener, router); err != nil {
			log.Fatal(err)
		}
	}()

	return <-addr
}

type FileResult struct {
	File  string
	Error string
}

func OpenDirDialog(title string, initPath string) FileResult {
	dial := dialog.Directory().Title(title)
	if len(initPath) > 0 {
		dial.SetStartDir(initPath)
	}
	p, e := dial.Browse()
	if e != nil {
		return FileResult{File: "", Error: e.Error()}
	}
	return FileResult{File: p, Error: ""}
}
func OpenFileDialog(title string, save bool, extention []string, initPath string) FileResult {
	dial := dialog.File().Title(title)
	if len(extention) > 0 {
		dial.Filter("", extention...)
	}
	if len(initPath) > 0 {
		dial.SetStartDir(initPath)
	}
	if save {
		p, e := dial.Save()
		if e != nil {
			return FileResult{File: "", Error: e.Error()}
		}
		return FileResult{File: p, Error: ""}
	} else {
		p, e := dial.Load()
		if e != nil {
			return FileResult{File: "", Error: e.Error()}
		}
		return FileResult{File: p, Error: ""}
	}
}
func startWebView(urls string) {
	w := webview.New(800, 600, false, true)
	defer w.Destroy()
	w.Bind("openfile", OpenFileDialog)
	w.Bind("opendir", OpenDirDialog)

	//directory, err := dialog.Directory().Title("Load images").Browse()
	w.SetTitle("demo")
	//w.SetSize(800, 600, webview.HintNone)
	//w.Navigate("http://appuploader.net")
	w.Navigate(urls)
	w.Run()
}
func main() {
	host := startServer()
	startWebView("http://" + host)

	//directory, err := dialog.Directory().Title("Load images").Browse()
	//fmt.Println("%v %v", directory, err)
}
