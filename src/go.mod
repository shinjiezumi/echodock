module echodock

go 1.15

require (
	github.com/aws/aws-sdk-go v1.42.22
	github.com/gorilla/sessions v1.2.1
	github.com/jinzhu/now v1.1.4 // indirect
	github.com/joho/godotenv v1.4.0
	github.com/labstack/echo-contrib v0.11.0
	github.com/labstack/echo/v4 v4.6.1
	github.com/labstack/gommon v0.3.1 // indirect
	github.com/mattn/go-colorable v0.1.12 // indirect
	github.com/rakyll/statik v0.1.7
	github.com/stretchr/testify v1.7.0
	golang.org/x/crypto v0.0.0-20211209193657-4570a0811e8b // indirect
	golang.org/x/net v0.0.0-20211209124913-491a49abca63 // indirect
	golang.org/x/sys v0.1.0 // indirect
	golang.org/x/time v0.0.0-20211116232009-f0f3c7e86c11 // indirect
	gorm.io/driver/mysql v1.2.1
	gorm.io/gorm v1.22.4
)

replace gopkg.in/urfave/cli.v2 => github.com/urfave/cli/v2 v2.1.1

replace github.com/dgrijalva/jwt-go v3.2.0+incompatible => github.com/golang-jwt/jwt/v4 v4.1.0
