module echodock

go 1.13

require (
	github.com/aws/aws-sdk-go v1.42.20
	github.com/gorilla/sessions v1.2.1
	github.com/jinzhu/now v1.1.4 // indirect
	github.com/joho/godotenv v1.4.0
	github.com/labstack/echo-contrib v0.11.0
	github.com/labstack/echo/v4 v4.6.1
	github.com/labstack/gommon v0.3.1 // indirect
	github.com/mattn/go-colorable v0.1.12 // indirect
	github.com/rakyll/statik v0.1.7
	github.com/stretchr/testify v1.7.0
	golang.org/x/crypto v0.0.0-20211202192323-5770296d904e // indirect
	golang.org/x/net v0.0.0-20211208012354-db4efeb81f4b // indirect
	golang.org/x/sys v0.0.0-20211205182925-97ca703d548d // indirect
	golang.org/x/time v0.0.0-20211116232009-f0f3c7e86c11 // indirect
	gorm.io/driver/mysql v1.2.1
	gorm.io/gorm v1.22.4
)

replace gopkg.in/urfave/cli.v2 => github.com/urfave/cli/v2 v2.1.1
