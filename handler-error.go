package main

import "net/http"

func handlerError(w http.ResponseWriter, r *http.Request){
	respondwithError(w,400,"Something Went Wrong !")
}