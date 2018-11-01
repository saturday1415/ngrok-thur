package server

import (
	"errors"
	"ngrok-thur/msg"
)

//token验证
func TokenVerification(token string) (permission string,err error) {
	if token != "fff33d74a0e53e98cc9051d99c326cd1" && token != "fff33d74a0e53e98cc9051d99c326cd11" {
		return permission, errors.New("Token is incorrect")
	}
	return permission,err
}

//register the clientid
func GetClientId() (clientid string, err error) {
	// it's a new session, assign an ID
	clientid, err = util.SecureRandId(32);
	return clientid, err
	

	if token == "fff33d74a0e53e98cc9051d99c326cd1" {
		return "123123123", err
	}
	if token == "fff33d74a0e53e98cc9051d99c326cd11" {
		return "123123123123", err
	}
	return clientid, errors.New("clientid is incorrect")
}

//隧道验证
func TunnelVerification(clientid string, rawTunnelReq *msg.ReqTunnel) (err error) {
	if clientid == "123123123" {
		if rawTunnelReq.Subdomain != "test" {
			return errors.New("Hostname is incorrect")
		}
	}
	if clientid == "123123123123" {
		if rawTunnelReq.Subdomain != "www" {
			return errors.New("Hostname is incorrect")
		}
	}

	return err
}
