package pkgHTTP

import (
	"flag"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

func NewServerHTTP() {
	addr := flag.String("addr", ":8181", "Сетевой адрес сервера")
	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)

	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	newRouter := mux.NewRouter()

	newRouter.HandleFunc("/fibonacci", GetFibonacci).Methods("POST")

	infoLog.Printf("Запуск сервера на %s", *addr)
	err := http.ListenAndServe(*addr, newRouter)
	errorLog.Fatal(err)
}
