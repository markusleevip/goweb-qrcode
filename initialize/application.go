package initialize

import "github.com/spf13/viper"

type application struct {
	Port       string
	Name       string
	BaseUrl    string
	LogPath    string
	QrcodeSize int
}

func initApplication(cfg *viper.Viper) *application {
	return &application{
		Port:       cfg.GetString("port"),
		Name:       cfg.GetString("name"),
		BaseUrl:    cfg.GetString("baseUrl"),
		QrcodeSize: cfg.GetInt("qrcodeSize"),
	}
}

var Application = new(application)
