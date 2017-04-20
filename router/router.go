package router

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/haritsE/test-iap/config"
)

func CreateRouter(config config.Config) *mux.Router {
	r := mux.NewRouter()
	initiateRouter(r, config)

	return r
}

func initiateRouter(router *mux.Router, config config.Config) {

	// Accessory
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		t := template.New("index.html")               // Create a template.
		t, err := t.ParseFiles("template/index.html") // Parse template file.
		if err != nil {
			fmt.Println("Error: %v", err)
			return
		}
		err = t.Execute(w, nil) // merge.
		if err != nil {
			fmt.Println("Error: %v", err)
			return
		}
	}).Methods("GET")

	fmt.Println("test-iap started in Port: " + config.Port)
	err := http.ListenAndServe(":"+config.Port, router)
	if err != nil {
		fmt.Println("%v", err)
	}
}
