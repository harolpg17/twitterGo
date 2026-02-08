package bd

import (
	"github.com/harolpg17/twitterGo/models"
	"golang.org/x/crypto/bcrypt"
)

func IntentoLogin(email string, password string) (models.Usuario, bool) {
	usuario, encontrado, _ := ChequeoYaExisteUsuario(email)
	if !encontrado {
		return usuario, false
	}

	passwordBytes := []byte(password)
	passwordDB := []byte(usuario.Password)

	err := bcrypt.CompareHashAndPassword(passwordDB, passwordBytes)
	if err != nil {
		return usuario, false
	}

	return usuario, true
}
