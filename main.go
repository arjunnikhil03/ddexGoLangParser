package main

import (
	"net/http"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/moodleexpert/ddexGoLangParser/controllers"
	"github.com/moodleexpert/ddexGoLangParser/utils"
)

func main() {
	router := mux.NewRouter()

	host := "localhost"
	port := "3000"
	url := host + ":" + port

	utils.Log.Info("Server Running On Port: " + port)

	//CORS Options
	headers := handlers.AllowedHeaders([]string{"Access-Control-Allow-Origin", "Authorization", "Content-Type", "Accept", "Accept-Language", "Accept-Encoding", "Access-Control-Request-Headers", "Access-Control-Request-Method"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "OPTIONS"})
	origins := handlers.AllowedOrigins([]string{"*"})

	router.HandleFunc("/auth/user/login", controllers.CreateLogin).Methods("POST")
	router.HandleFunc("/ddex/process", controllers.Process).Methods("POST")

	w := utils.Log.Writer()
	defer w.Close()
	srv := &http.Server{
		ReadTimeout:  120 * time.Second,
		WriteTimeout: 120 * time.Second,
		IdleTimeout:  120 * time.Second,
		Addr:         url,
		Handler:      utils.New(handlers.CORS(headers, methods, origins)(router), utils.Log),
	}
	err := srv.ListenAndServe()

	if err != nil {
		utils.Log.Panic(err)
	}
}
