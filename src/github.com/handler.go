package main

import (
	"./domain/model"
	"github.com/labstack/echo"
	"net/http"
	"github.com/jinzhu/gorm"
	"time"
	"fmt"
)

// -------------------------------------------------------------------------------------
//                                                                            Definition
//                                                                            ----------
/**
 * Routing
 */
func setRoute(e *echo.Echo) {
	e.GET("/hello", handleIndex)
	e.GET("/select", handleTalentSelectByID)
	e.GET("/setup", handleTalentSetup)
}

/**
 * DBに接続します
 */
func connect() *gorm.DB {
	db, err := gorm.Open("mysql", "root@/tmgs?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}

	return db
}

// -------------------------------------------------------------------------------------
//                                                                               Handler
//                                                                               -------

/**
 * IndexのHandler
 */
func handleIndex(c echo.Context) error {
	return c.Render(http.StatusOK, "hello", "world")
}

/**
 * タレントを取得します
 */
func handleTalentSelectByID(c echo.Context) error {
	db := connect()

	// 実行完了後DB接続を閉じる
	defer db.Close()

	// ログ出力を有効にする
	db.LogMode(true)

	/** SELECT */
	var talent model.Talent
	talent.Talent_Id = 3
	db.First(&talent)
	fmt.Printf("mail_address: " + talent.Email_Address + "\n")

	return c.Render(http.StatusOK, "hello", "Select")
}

/**
 * タレントを登録します
 */
func handleTalentSetup(c echo.Context) error {
	db := connect()

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

	/** PRINT */
	fmt.Printf("****************************************************************************************\n")
	fmt.Printf("Hello, World!\n")

	tx.Commit()

	return c.Render(http.StatusOK, "hello", "Setup")
}