package service

import (
	"net/http"
)


type ServerConfig struct {
	SpiffeID string
	Port     string
}

var Config ServerConfig

func handle(response http.ResponseWriter, request *http.Request) {

	match, _:= MatchCertificateID(Config.SpiffeID, request.Body)

	if !match {
		response.WriteHeader(401)
	}
}
