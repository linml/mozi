package common

import "golang.org/x/crypto/bcrypt"

func CheckBCrypt(str string, hashedStr string) bool {

	err := bcrypt.CompareHashAndPassword([]byte(hashedStr), []byte(str))
	if err != nil {
		return false
	}
	return true
}
