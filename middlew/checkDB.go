package middlew

import (
	"net/http"

	"github.com/jacobd39/admsys/bd"
)

//CheckDB verifica que se ha establecido una conexión previa
func CheckDB(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if bd.CheckDBConnection() == 0 {
			http.Error(w, "Lost connection with the database", 500)
			return
		}
		next.ServeHTTP(w, r)
	}
}
