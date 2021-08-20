package handler

import (
	"encoding/json"
	"io"
	"strconv"

	"golang.org/x/crypto/bcrypt"
)

type Response struct {
	Status  bool        `json:"status"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

type IdResponse struct {
	Id uint64
}

func parseID(idJSON string) (uint64, error) {

	id, err := strconv.ParseUint(idJSON, 10, 8)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func ToJSON(i interface{}, w io.Writer) error {
	encoder := json.NewEncoder(w)
	return encoder.Encode(i)
}

func FromJSON(i interface{}, r io.Reader) error {
	decoder := json.NewDecoder(r)
	return decoder.Decode(i)
}

func HashPassword(pass string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(pass), 8)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil

}

func ComparePassword(hashedPassword, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return err
	}
	return nil
}
