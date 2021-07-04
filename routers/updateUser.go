package routers

import (
	"encoding/json"
	"net/http"

	"github.com/jacobd39/admsys/bd"
	model "github.com/jacobd39/admsys/models"
)

//UpdateUser actualiza la información de un usuario en la base de datos
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	var t model.User
	err := json.NewDecoder(r.Body).Decode(&t)

	if err != nil {
		http.Error(w, "Error en la Data "+err.Error(), http.StatusBadRequest)
		return
	}

	var status bool

	status, err = bd.UpdateRecord(t, ID)

	if err != nil {
		http.Error(w, "Ocurrió un error cuando se intentaba actualizar el registro. Por favor, intente nuevamente. "+err.Error(), http.StatusBadRequest)
		return
	}

	if status == false {
		http.Error(w, "El registro del usuario no ha sido actualizado", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
