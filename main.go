package main

import (
	"github.com/gorilla/mux"
	"app/Routers"
	"net/http"
	"app/Utils"
	//"fmt"
)

func main() {
		/*
		Routing using mux
		*/

		Utils.GetNumCpu()
		r := mux.NewRouter()
		r.HandleFunc("/set",Routers.SetHandler)
		r.HandleFunc("/get",Routers.GetHandler)
		r.HandleFunc("/map",Routers.MappingHandler)
		http.Handle("/",r);
		http.ListenAndServe(":8000",r)

	}
