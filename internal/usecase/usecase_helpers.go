package usecase

import (
	"crypto/hmac"
	"crypto/sha256"
	"errors"
	"fmt"

	"solver/internal/model"
)

func findMissingNumbers(task []uint32) ([]uint32, error) {
	existedNumbers := make([]bool, len(task))
	result := make([]uint32, 0)
	for _, v := range task {
		existedNumbers[v] = true
	}

	for i, v := range existedNumbers {
		if v == false {
			result = append(result, uint32(i+1))
		}
	}

	return result, nil
}

func validateUser(user *model.User) error {
	if user.FirstName == "" {
		return errors.New("invalid first name length")
	}

	if user.SecondName == "" {
		return errors.New("invalid second name length")
	}

	if user.FatherName == "" {
		return errors.New("invalid father name length")
	}

	if user.GroupName == "" {
		return errors.New("invalid group name length")
	}

	if user.Password == "" {
		return errors.New("invalid password length")
	}

	return nil
}

func getPasswordHash(password string) string {
	src := []byte(password)

	h := hmac.New(sha256.New, []byte(""))
	h.Write(src)
	dst := fmt.Sprintf("%x", h.Sum(nil))

	return dst
}
