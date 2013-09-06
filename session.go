package socketio

import (
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

type Session struct {
	Id                 string
	HeartbeatTimeout   int
	ConnectionTimeout  int
	SupportedProtocols []string
}

func NewSession(url string) (*Session, error) {
	// Initiate the session via http request
	response, err := http.Get("http://" + url+"/socket.io/1")
	if err != nil {
		return nil, err
	}

	// Read the response
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	response.Body.Close()

	// Extract the session configs from the response
	sessionVars := strings.Split(string(body), ":")
	id := sessionVars[0]
	heartbeatTimeout, _ := strconv.Atoi(sessionVars[1])
	connectionTimeout, _ := strconv.Atoi(sessionVars[2])
	supportedProtocols := strings.Split(string(sessionVars[3]), ",")

	return &Session{id, heartbeatTimeout, connectionTimeout, supportedProtocols}, nil
}
