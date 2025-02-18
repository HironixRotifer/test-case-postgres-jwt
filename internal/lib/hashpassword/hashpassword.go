package hashpassword

import (
	"crypto/rand"
	"crypto/sha512"
	"fmt"
)

const saltSize = 16

// HashPassword возвращает хэш пароля
func HashPassword(password string, salt []byte) string {
	var hash = sha512.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum(salt))
}

// DoPasswordsMatch сравнивает пароли
func DoPasswordsMatch(hashedPassword, currPassword string, salt []byte) bool {
	var currPasswordHash = HashPassword(currPassword, salt)
	return hashedPassword == currPasswordHash
}

// GenerateRandomSalt генерирует соль для пароля
func GenerateRandomSalt() []byte {
	var salt = make([]byte, saltSize)

	_, err := rand.Read(salt[:])
	if err != nil {
		panic(err)
	}

	return salt
}
