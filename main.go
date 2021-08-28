package main

import (
	"log"
	"net/http"

	"alejandro/encryptSampleTest/conmons"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	port := ":8888"
	router := mux.NewRouter()
	conmons.SetCodeRoutes(router)

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:4200"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowCredentials: true,
		AllowedHeaders:   []string{"Accept", "Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization"},
	})
	go conmons.Scheduler()
	log.Printf("ejecutansose servidor en el puerto %v", port)
	http.ListenAndServe(port, c.Handler(router))
}

/**

func main() {

	text := "My Super Secret Code Stuff"
	key := []byte("passphrasewhichneedstobe32bytes!")

	mensaje := Encrypting(key, text)
	decryptingMensaje := Decrypting(key, mensaje)

	fmt.Println(decryptingMensaje)
}

/api/valor





*/
