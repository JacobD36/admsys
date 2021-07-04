package routers

import (
	"encoding/json"
	"net/http"

	"github.com/jacobd39/admsys/bd"
)

//GetUser llama a la función seekUser para buscar un registro en la base de datos
func GetUser(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")

	if len(ID) < 1 {
		http.Error(w, "Parámetro ID requerido", http.StatusBadRequest)
		return
	}

	usrData, err := bd.SeekUser(ID)

	if err != nil {
		http.Error(w, "Ocurrió un error cuando se intentaba buscar el registro "+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Context-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(usrData)
}
