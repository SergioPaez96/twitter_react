package bd

import "golang.org/x/crypto/bcrypt"

/*EncriptarPassword es la rutina que me permite encriotar la password*/
func EncriptarPassword(pass string) (string, error) {
	costo := 8
	bytes, err := bcrypt.GenerateFromPassword([]byte(pass), costo)
	return string(bytes), err
}
