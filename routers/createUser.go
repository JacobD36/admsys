package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/jacobd39/admsys/bd"
	model "github.com/jacobd39/admsys/models"
)

//CreateUser función que registra un usuario en la BD
func CreateUser(w http.ResponseWriter, r *http.Request) {
	var t model.User

	err := json.NewDecoder(r.Body).Decode(&t)

	if err != nil {
		http.Error(w, "Error en la data recibida "+err.Error(), 400)
		return
	}

	if len(t.Name1) == 0 {
		http.Error(w, "El primer nombre es requerido", 400)
		return
	}

	if len(t.LastName1) == 0 {
		http.Error(w, "El primer apellido es requerido", 400)
		return
	}

	if len(t.LastName2) == 0 {
		http.Error(w, "El segundo apellido es requerido", 400)
		return
	}

	_, userFound, _ := bd.UserVerification(t.Dni)

	if userFound == true {
		http.Error(w, "El usuario ya existe", 400)
		return
	}

	t.CreateAt = time.Now()
	t.ModifiedAt = time.Now()

	idUsr, status, err := bd.InsertNewUser(t)

	if err != nil {
		http.Error(w, "Ocurrió un error cuando se intentó registrar al usuario "+err.Error(), 400)
		return
	}

	if status == false {
		http.Error(w, "El registro del usuario no puede ser insertado", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(idUsr)
}
