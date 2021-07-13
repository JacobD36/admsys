package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/jacobd39/admsys/middlew"
	"github.com/jacobd39/admsys/routers"
	"github.com/rs/cors"
)

//BackendHandlers setean mi puerto, el handler y luego pongo a escuchar al servidor
func BackendHandlers() {
	router := mux.NewRouter()

	router.HandleFunc("/newUser", middlew.CheckDB(middlew.CheckJWT(routers.CreateUser))).Methods("POST")
	router.HandleFunc("/usersList", middlew.CheckDB(middlew.CheckJWT(routers.UsersList))).Methods("GET")
	router.HandleFunc("/getUser", middlew.CheckDB(middlew.CheckJWT(routers.GetUser))).Methods("GET")
	router.HandleFunc("/updateUser", middlew.CheckDB(middlew.CheckJWT(routers.UpdateUser))).Methods("PUT")
	router.HandleFunc("/seekUser", middlew.CheckDB(middlew.CheckJWT(routers.SeekProfile))).Methods("GET")
	router.HandleFunc("/login", middlew.CheckDB(routers.Login)).Methods("POST")
	router.HandleFunc("/getCampaigns", middlew.CheckDB(middlew.CheckJWT(routers.ListCampaign))).Methods("GET")
	router.HandleFunc("/newCampaign", middlew.CheckDB(middlew.CheckJWT(routers.CreateCampaign))).Methods("POST")
	router.HandleFunc("/newProfile", middlew.CheckDB(middlew.CheckJWT(routers.CreateProfile))).Methods("POST")
	router.HandleFunc("/getProfiles", middlew.CheckDB(middlew.CheckJWT(routers.ListProfiles))).Methods("GET")
	router.HandleFunc("/seekProfile", middlew.CheckDB(middlew.CheckJWT(routers.SeekDBProfile))).Methods("GET")
	router.HandleFunc("/newPosition", middlew.CheckDB(middlew.CheckJWT(routers.CreatePosition))).Methods("POST")
	router.HandleFunc("/getPositions", middlew.CheckDB(middlew.CheckJWT(routers.ListPositions))).Methods("GET")

	PORT := os.Getenv("PORT")

	handler := cors.AllowAll().Handler(router)

	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
