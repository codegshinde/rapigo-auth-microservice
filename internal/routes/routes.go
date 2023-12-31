package routes

import (
	"rapigo/internal/handlers"
)

const (
	RegisterPath = "/register"
	LoginPath    = "/login"
	InfoPath     = "/info"
)

func SetupRoutes() {

	//router.(SpeedPath, handlers.SpeedHandler)
	routes := SetupRouter()
	routes.POST(LoginPath, handlers.Login)
	routes.POST(RegisterPath, handlers.Register)
	routes.GET(InfoPath, handlers.Info)
}
