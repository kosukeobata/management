package main

import (
	"./config"
	"html/template"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
)

// -------------------------------------------------------------------------------------
//                                                                            Definition
//                                                                            ----------
// レイアウト適用済のテンプレートを保存するmap
var templates map[string]*template.Template

// -------------------------------------------------------------------------------------
//                                                                                  Main
//                                                                                  ----
func main() {
	t := &Template{
		templates: template.Must(template.ParseGlob("templates/*.html")),
	}

	// echoの新規インスタンス
	e := echo.New()

	// ログの出力レベルを設定
	e.Logger.SetLevel(log.DEBUG)


	e.Renderer = t

	// ルート設定
	setRoute(e)

	e.Logger.Fatal(e.Start(config.Server.Port))
}

// -------------------------------------------------------------------------------------
//                                                                            Initialize
//                                                                            ----------
func init() {
	config.Load()
}