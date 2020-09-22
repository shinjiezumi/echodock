package util

import "fmt"

const appName = "echodock"

func GenerateTitle(pageTitle string) string {
	return fmt.Sprintf("%s|%s", pageTitle, appName)
}
