package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/jacobd39/admsys/bd"
	model "github.com/jacobd39/admsys/models"
)

//CreatePosition registra un nuevo puesto en la base de datos
func CreatePosition(w http.ResponseWriter, r *http.Request) {
	var t model.Positions

	err := json.NewDecoder(r.Body).Decode(&t)

	if err != nil {
		http.Error(w, "Error en la data recibida "+err.Error(), http.StatusBadRequest)
		return
	}

	if len(t.Title) == 0 {
		http.Error(w, "El campo título es obligatorio", http.StatusBadRequest)
		return
	}

	t.CreateAt = time.Now()
	t.ModifiedAt = time.Now()

	idCamp, status, err := bd.InsertNewPosition(t)

	if err != nil {
		http.Error(w, "Ocurrió un error cuando se intentó registrar la campaña "+err.Error(), http.StatusBadRequest)
		return
	}

	if !status {
		http.Error(w, "El registro del puesto no puede ser insertado", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(idCamp)
}
