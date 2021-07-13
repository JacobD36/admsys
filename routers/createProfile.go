package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/jacobd39/admsys/bd"
	model "github.com/jacobd39/admsys/models"
)

//CreateProfile registra un perfil en la base de datos
func CreateProfile(w http.ResponseWriter, r *http.Request) {
	var t model.Profile

	err := json.NewDecoder(r.Body).Decode(&t)

	if err != nil {
		http.Error(w, "Error en la data recibida "+err.Error(), http.StatusBadRequest)
		return
	}

	if len(t.Title) == 0 {
		http.Error(w, "El nombre del perfil es requerido", http.StatusBadRequest)
		return
	}

	t.CreateAt = time.Now()
	t.ModifiedAt = time.Now()

	idProfile, status, err := bd.InsertNewProfile(t)

	if err != nil {
		http.Error(w, "Ocurrió un error cuando se intentó registrar el perfil "+err.Error(), http.StatusBadRequest)
		return
	}

	if status == false {
		http.Error(w, "El registro del perfil no pudo ser insertado", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(idProfile)
}
