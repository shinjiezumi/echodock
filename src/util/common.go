package util

import "fmt"

// アプリ名
const appName = "echodock"

// GeneratePageTitle はページタイトルを生成します
func GeneratePageTitle(pageName string) string {
	return fmt.Sprintf("%s|%s", pageName, appName)
}
