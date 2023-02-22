package person

import (
	"encoding/json"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func JSONToStruct(value interface{}, out any) error {
	valueByte, err := json.Marshal(value)
	if err != nil {
		return err
	}
	err = json.Unmarshal(valueByte, out)
	if err != nil {
		return err
	}
	return nil
}
