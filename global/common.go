package global

import (
	"goweb-qrcode/initialize"

	"go.uber.org/zap"
)

const (
	Version = "1.0.0"
)

var (
	BaseUrl    = ""
	Logger     *zap.Logger
	QrcodeSize = 128
)

func Initialize() {
	BaseUrl = initialize.Application.BaseUrl
	QrcodeSize = initialize.Application.QrcodeSize
}
