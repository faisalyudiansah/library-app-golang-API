package helpers

import "golang.org/x/crypto/bcrypt"

type Bcrypt interface {
	HashPassword(string, int) ([]byte, error)
	CheckPassword(string, []byte) (bool, error)
}

type BcryptStruct struct {
}

func NewBcryptStruct() *BcryptStruct {
	return &BcryptStruct{}
}

func (bc *BcryptStruct) HashPassword(pwd string, cost int) ([]byte, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), cost)
	if err != nil {
		return nil, err
	}
	return hash, nil
}

func (bc *BcryptStruct) CheckPassword(pwd string, hash []byte) (bool, error) {
	err := bcrypt.CompareHashAndPassword(hash, []byte(pwd))
	if err != nil {
		return false, err
	}
	return true, nil
}
