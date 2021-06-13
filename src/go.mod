module github.com/shinjiezumi/echodock/src

go 1.13

require (
	github.com/aws/aws-sdk-go v1.38.60
	github.com/gorilla/sessions v1.2.1
	github.com/joho/godotenv v1.3.0
	github.com/labstack/echo-contrib v0.11.0
	github.com/labstack/echo/v4 v4.3.0
	github.com/rakyll/statik v0.1.7
	github.com/stretchr/testify v1.7.0
	gorm.io/driver/mysql v1.1.0
	gorm.io/gorm v1.21.10
)

replace gopkg.in/urfave/cli.v2 => github.com/urfave/cli/v2 v2.1.1
