package routes

import (
	"goweb-qrcode/handles"

	"github.com/gofiber/fiber/v2"
)

func InitRoutes(app *fiber.App) {
	app.Get("/qrcode/:uid", handles.GetQrCode)
}
