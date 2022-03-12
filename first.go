package main

import "net/http"

func yourHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World//"))
}

/*func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", yourHandler)
	panic(http.ListenAndServe(":8080", r))
}*/
