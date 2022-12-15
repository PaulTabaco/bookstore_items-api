package controllers

import "net/http"

const (
	pongAnswer = "pong \n"
)

var (
	PingController pingControllerInterface = &pingController{}
)

type pingControllerInterface interface {
	Ping(w http.ResponseWriter, r *http.Request)
}

type pingController struct{}

func (c *pingController) Ping(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(pongAnswer))
	// c.String(http.StatusOK, "pong\n")
}
