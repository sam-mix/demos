package main

import (
	"database/sql/driver"
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type XTime struct {
	time.Time
}

// 2. 为 Xtime 重写 MarshaJSON 方法，在此方法中实现自定义格式的转换；
func (t XTime) MarshalJSON() ([]byte, error) {
	output := fmt.Sprintf("\"%s\"", t.Format("2006-01-02 15:04:05"))
	return []byte(output), nil
}

// 3. 为 Xtime 实现 Value 方法，写入数据库时会调用该方法将自定义时间类型转换并写入数据库；
func (t XTime) Value() (driver.Value, error) {
	var zeroTime time.Time
	if t.Time.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return t.Time, nil
}

// 4. 为 Xtime 实现 Scan 方法，读取数据库时会调用该方法将时间数据转换成自定义时间类型；
func (t *XTime) Scan(v interface{}) error {
	value, ok := v.(time.Time)
	if ok {
		*t = XTime{Time: value}
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}

type Model struct {
	ID        int            `gorm:"primary_key;comment:gorm通用生成字段,自增id" json:"id" uri:"id"`
	CreatedAt XTime          `json:"created_at" gorm:"autoCreateTime;comment:gorm通用生成字段,创建时间"`
	UpdatedAt XTime          `json:"updated_at" gorm:"autoUpdateTime;comment:gorm通用生成字段,更新时间" `
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index;comment:gorm通用生成字段,删除时间"`
}

// 玩家模型
type PlayerModel struct {
	// 基础数据
	Model
	Name   string `desc:"昵称" json:"name" form:"name" gorm:"type:varchar(255);not null;comment:昵称"`                                        // 昵称
	Icon   string `desc:"头像" json:"icon" form:"icon" gorm:"type:varchar(255);not null;comment:头像"`                                        // 头像
	Gender uint8  `desc:"性别 1:男,2:女,3:保密" json:"gender" form:"gender" gorm:"type:tinyint(1);not null; default 3;comment:性别 1:男,2:女,3:保密"` // 性别 1:男;2:女;3:保密
	Gold   string `desc:"金币" json:"gold" form:"gold" gorm:"type:varchar(255);not null; default '';comment:金币"`                            // 金币
	Dia    string `desc:"钻石" json:"dia" form:"gold" gorm:"type:varchar(255);not null; default '';comment:钻石"`                             // 钻石
	Acc    string `desc:"账号 (account 简写)" json:"acc" form:"acc" gorm:"type:varchar(255); not null; uniqueIndex; comment:账号 (account 简写)"` // 账号 (account 简写)
}

// gorm 自定义表名
func (PlayerModel) TableName() string {
	return "player"
}

func main() {
	// Connect to MySQL database
	dsn := "root:123456@tcp(127.0.0.1:3306)/test1?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	}

	// Auto-migrate the "my_users" table
	db.AutoMigrate(&PlayerModel{})
}
