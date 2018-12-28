package login

import (
	"crypto/md5"
)

type LoginService struct {

}


func (self *LoginService) LoginProcess(username, password string) bool{
	if username != "caijiang" {
		return false
	}
	md5 := md5.New()
	cryptoData := md5.Sum([]byte(password))
	successPassword := md5.Sum([]byte("caijiang"))
	if string(cryptoData) != string(successPassword) {
		return false
	}
	return true
}

