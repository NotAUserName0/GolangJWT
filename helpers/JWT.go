package helpers

import (
	"GoJWT/models"
	"fmt"
	"os"

	"github.com/golang-jwt/jwt/v5"
)

 func GenerateToken(user models.User) (string, error) {
	tokenBuilder := jwt.NewWithClaims(jwt.SigningMethodHS256 , jwt.MapClaims{
		"user":user.Email,
		"nombre":user.Nombre,
		"apellido":user.Apellido,
		"edad":fmt.Sprint(user.Edad),
		"genero":user.Genero,
		"rol":user.Rol,
	})

	tokenString, err := tokenBuilder.SignedString([]byte(os.Getenv("SECRET")))

	return tokenString,err
	
}