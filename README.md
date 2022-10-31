# goweb-qrcode
## config.yml
```
application:
  name: WebQrCode
  port: 8001
  baseUrl: http://jobathome.cn/profile/%s
  qrcodeSize: 256   
```
## Build
```
git clone https://github.com/markusleevip/goweb-qrcode.git
cd goweb-qrcode
go run qrcode.go
```
## Snapshot
![demo](/doc/qrcode-demo.png)