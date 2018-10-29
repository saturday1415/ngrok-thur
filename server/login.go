package server

import "errors"

//token验证
func TokenVerification(token string) (err error) {
	if token != "fff33d74a0e53e98cc9051d99c326cd1" {
		return errors.New("Token is incorrect")
	}
	return
}

//隧道验证
func TunnelVerification() {

}
