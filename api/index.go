package handler

import (
	"net/http"

	"github.com/gofiber/fiber/v2/adaptor"

	"backend/app"
)

var server = app.CreateApp()

func Handler(w http.ResponseWriter, r *http.Request) {
	adaptor.HTTPHandler(server)(w, r)
}