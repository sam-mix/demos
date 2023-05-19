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

// 玩法模型
type WanFaModel struct {
	// 基础数据
	Model
	WanId       uint16 `desc:"主玩法ID" json:"wan_id" form:"wan_id" gorm:"type:smallint unsigned;size:5;uniqueIndex;not null;comment:主玩法ID"`      // 主玩法ID
	Name        string `desc:"主玩法名称" json:"name" form:"name" gorm:"type:varchar(255);not null;comment:主玩法名称"`                                  // 主玩法名称
	Sort        uint16 `desc:"序号" json:"sort" form:"sort" gorm:"type:smallint unsigned;size:5;uniqueIndex;not null;comment:排序序号"`              // 排序序号
	SecondWanId uint16 `desc:"二级玩法ID" json:"second_wan_id" form:"second_wan_id" gorm:"type:smallint unsigned;size:5;default 0;comment:二级玩法ID"` // 二级玩法ID
	SecondName  string `desc:"二级玩法名称" json:"second_name" form:"second_name" gorm:"type:varchar(255);comment:二级玩法名称"`                           // 二级玩法名称
	ThirdWanId  uint16 `desc:"三级玩法ID" json:"third_wan_id" form:"third_wan_id" gorm:"type:smallint unsigned;size:5;default 0;comment:三级玩法ID"`   // 二级玩法ID
	ThirdName   string `desc:"三级玩法名称" json:"third_name" form:"third_name" gorm:"type:varchar(255);comment:三级玩法名称"`                             // 二级玩法名称
}

// gorm 自定义表名
func (WanFaModel) TableName() string {
	return "wan_fa"
}
func main() {
	// Connect to MySQL database
	dsn := "root:123456@tcp(127.0.0.1:3306)/lobby?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	}

	// Auto-migrate the "my_users" table
	db.AutoMigrate(&WanFaModel{})
	// db.Create(&WanFaModel{
	// 	WanId:       13,
	// 	Name:        "东北麻将",
	// 	Sort:        7,
	// 	SecondWanId: 14,
	// 	SecondName:  "东北麻将不洗牌",
	// 	ThirdWanId:  0,
	// 	ThirdName:   "",
	// })
	wanFas := FindAllWanFa(db)
	fmt.Println(wanFas)
}

func FindAllWanFa(db *gorm.DB) (result []WanFaModel) {
	db.Find(&result)
	return
}
