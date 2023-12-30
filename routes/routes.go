package routes

import (
	"rapigo/handlers"
	"rapigo/pkg/router"
)

const (
	SpeedPath = "/speed"
	LoginPath = "/login"
)

func SetupRoutes() {
	//router.(SpeedPath, handlers.SpeedHandler)
	routes := router.SetupRouter()
	routes.POST(LoginPath, handlers.Login)
}
