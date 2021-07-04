package routers

import (
	"errors"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/jacobd39/admsys/bd"
	model "github.com/jacobd39/admsys/models"
)

//CodUser es el valor de codUser usado en todos los endpoints
var Dni string

//IDUsuario es el ID devuelto por el modelo que se usará en todos los endpoints
var IDUsuario string

//ProcessToken es el proceso que extrae los valores del token
func ProcessToken(tk string) (*model.Claim, bool, string, error) {
	myKey := []byte("kbjnfqfsfy79")
	claims := &model.Claim{}

	splitToken := strings.Split(tk, "Bearer")

	if len(splitToken) != 2 {
		return claims, false, string(""), errors.New("invalid token format")
	}

	tk = strings.TrimSpace(splitToken[1])

	tkn, err := jwt.ParseWithClaims(tk, claims, func(token *jwt.Token) (interface{}, error) {
		return myKey, nil
	})

	if err == nil {
		_, found, _ := bd.UserVerification(claims.Dni)

		if found == true {
			Dni = claims.Dni
			IDUsuario = claims.ID.Hex()
		}

		return claims, found, IDUsuario, nil
	}

	if !tkn.Valid {
		return claims, false, string(""), errors.New("Invalid Token")
	}

	return claims, false, string(""), err
}
