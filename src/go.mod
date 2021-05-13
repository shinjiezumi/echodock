module github.com/shinjiezumi/echodock/src

go 1.13

require (
	github.com/aws/aws-sdk-go v1.38.39
	github.com/gorilla/sessions v1.2.1
	github.com/joho/godotenv v1.3.0
	github.com/kr/pretty v0.1.0 // indirect
	github.com/labstack/echo-contrib v0.9.0
	github.com/labstack/echo/v4 v4.1.6
	github.com/labstack/gommon v0.3.0 // indirect
	github.com/mattn/go-colorable v0.1.7 // indirect
	github.com/stretchr/testify v1.4.0
	github.com/valyala/fasttemplate v1.2.1 // indirect
	golang.org/x/crypto v0.0.0-20200820211705-5c72a883971a // indirect
	gopkg.in/check.v1 v1.0.0-20180628173108-788fd7840127 // indirect
	gorm.io/driver/mysql v1.0.1
	gorm.io/gorm v1.20.1
)

replace gopkg.in/urfave/cli.v2 => github.com/urfave/cli/v2 v2.1.1
