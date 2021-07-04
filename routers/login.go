package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/jacobd39/admsys/bd"
	"github.com/jacobd39/admsys/jwt"
	model "github.com/jacobd39/admsys/models"
)

//Login realiza el login
func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	var t model.User

	err := json.NewDecoder(r.Body).Decode(&t)

	if err != nil {
		http.Error(w, "PETICIÓN INVÁLIDA", http.StatusBadRequest)
		return
	}

	if len(t.Dni) == 0 {
		http.Error(w, "NO SE INGRESO DNI", http.StatusBadRequest)
		return
	}

	document, exists, res := bd.LoginIntent(t.Dni, t.Password)

	if exists == false {
		if res == 1 {
			http.Error(w, "USUARIO NO ENCONTRADO", http.StatusBadRequest)
			return
		}

		if res == 2 {
			http.Error(w, "PASSWORD ERRONEO", http.StatusBadRequest)
			return
		}
	}

	jwtKey, err := jwt.GenerateJWT(document)

	if err != nil {
		http.Error(w, "ERROR EN TOKEN", http.StatusBadRequest)
		return
	}

	resp := model.LoginReturn{
		Token: jwtKey,
		ID:    document.ID.Hex(),
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)

	expirationTime := time.Now().Add(time.Hour)
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   jwtKey,
		Expires: expirationTime,
	})
}
