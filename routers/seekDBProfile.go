package routers

import (
	"encoding/json"
	"net/http"

	"github.com/jacobd39/admsys/bd"
)

//SeekDBProfile permite buscar un perfil determinado en la base de datos
func SeekDBProfile(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")

	if len(ID) < 1 {
		http.Error(w, "Debe enviar el parámetro ID", http.StatusBadRequest)
		return
	}

	profile, err := bd.SeekProfile(ID)

	if err != nil {
		http.Error(w, "Ocurrió un error al intentar buscar el registro "+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(profile)
}
