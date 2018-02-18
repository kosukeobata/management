package model

import "time"

// ---------------------------------------------------------------
//                                                Table Definition
//                                                ----------------
/** タレントテーブル定義 */
type Talent struct {
	Talent_Id int64 `gorm:"primary_key"`
	Last_Name string
	First_Name string
	Last_Name_Kana string
	First_Name_Kana string
	Email_Address string `sql:"not null;"`
	Password string
	Birth_Date time.Time
	Gender_Cd string
	Ins_Datetime time.Time
	Upd_Datetime time.Time
	Ins_Trace string
	Upd_Trace string
}

// ---------------------------------------------------------------
//                                                      Table Name
//                                                      ----------
/** テーブル名を単数形に変更 */
func (t *Talent) TableName() string {
	return "talent"
}
