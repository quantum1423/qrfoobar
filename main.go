package main

import "net/http"

import "log"

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.LUTC | log.Lshortfile)

	http.HandleFunc("/funcs/gen-qr", handGenQr)
	http.Handle("/", http.FileServer(http.Dir("./html/")))
	log.Fatalln(http.ListenAndServe("127.0.0.1:8081", nil))
}
