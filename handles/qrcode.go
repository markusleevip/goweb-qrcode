package handles

import (
	"fmt"
	"goweb-qrcode/global"
	"goweb-qrcode/kit"

	"github.com/gofiber/fiber/v2"
	qrcode "github.com/skip2/go-qrcode"
)

type Qrcode struct {
}

// http://127.0.0.1:8001/qrcode/1692528
func GetQrCode(ctx *fiber.Ctx) error {
	uid := ctx.Params("uid")
	url_str := ""
	if uid == "" {
		return ctx.Status(fiber.StatusOK).JSON(kit.FailAndMsg("Parameter error."))
	} else {
		url_str = fmt.Sprintf(global.BaseUrl, uid)
	}
	url := ctx.FormValue("url", url_str)
	var png []byte
	png, err := qrcode.Encode(url, qrcode.Medium, global.QrcodeSize)
	if err != nil {
		return ctx.Status(fiber.StatusOK).JSON("Generates a qrcode error.")
	}
	ctx.Response().Header.SetContentType("image/png")
	return ctx.Status(fiber.StatusOK).Send(png)
}
