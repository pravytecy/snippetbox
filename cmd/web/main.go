package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

func main() {
	//fmt.Println("Persistent wins")
	addr := flag.String("addr", ":4000", "HTTP network address")
	flag.Parse()
	infolog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	mux := http.NewServeMux()
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.HandleFunc("/", home)
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))
	mux.HandleFunc("/jerry/create-snippet", snippetCreate)
	mux.HandleFunc("/jerry/view-snippet", snippetView)

	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errLog,
		Handler:  mux,
	}

	infolog.Printf("Starting server on %s :", *addr)

	err := srv.ListenAndServe()
	errLog.Fatal(err)
}
