package api

import (
	"net/http"
	"fmt"
	"time"
	"strconv"
	"../../domain/model"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Route => handler
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!\n")
	})


	// Setup
	e.GET("/setup", setup)

	e.Logger.Fatal(e.Start(":3000"))
}

func setup(c echo.Context) error {
	db, err := gorm.Open("mysql", "root@/tmgs?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}

	// 実行完了後DB接続を閉じる
	defer db.Close()

	// ログ出力を有効にする
	db.LogMode(true)

	/** MIGRATION */
	db.AutoMigrate(&model.Talent{})

	tx := db.Begin()

	/** INSERT */
	if err := tx.Create(&model.Talent{Email_Address: "k_mementomori@yahoo.co.jp", Ins_Datetime: time.Now(), Upd_Datetime: time.Now()}).Error; err != nil {
		tx.Rollback()
	}

	/** SELECT */
	talent := model.Talent{}
	talent.Talent_Id = 1
	fmt.Printf("talent_id:" + strconv.FormatInt(talent.Talent_Id, 10))
	db.First(&talent)


	/** PRINT */
	fmt.Printf("****************************************************************************************\n")
	fmt.Printf("Hello, World!\n")
	fmt.Printf("めーるあどれす:" + talent.Email_Address + "\n")

	tx.Commit()
	return c.Render(http.StatusOK, "setup", nil)
}
