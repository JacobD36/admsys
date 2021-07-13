package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/jacobd39/admsys/bd"
)

//ListPositions retorna los puestos activos
func ListPositions(w http.ResponseWriter, r *http.Request) {
	search := r.URL.Query().Get("search")
	page := r.URL.Query().Get("page")

	pageTemp, err := strconv.Atoi(page)

	if err != nil {
		http.Error(w, "Se debe enviar el parámetro page como un entero mayor a cero", http.StatusBadRequest)
		return
	}

	pag := int64(pageTemp)

	result, status := bd.GetPositions(pag, search)

	if status == false {
		http.Error(w, "Ocurrió un error cuando se intentaba leer el listado de puestos", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(result)
}
