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
type ChangCiModel struct {
	Model
	WanId           uint16 `desc:"玩法ID" json:"wan_id" form:"wan_id" gorm:"type:smallint unsigned;size:5;not null;comment:玩法ID"`                     //玩法ID
	ChangId         uint16 `desc:"场次ID" json:"chang_id" form:"chang_id" gorm:"type:smallint unsigned;size:5;not null;comment:场次ID"`                 // 场次ID
	Name            string `desc:"场次名称" json:"name" form:"name" gorm:"type:varchar(255); not null;comment:场次名称"`                                    // 场次名称
	CardId          string `desc:"卡片样式ID" json:"card_id" form:"card_id" gorm:"type:varchar(255); not null; comment:卡片样式ID"`                         // 卡片样式ID
	TagId           string `desc:"展示标签" json:"tag_id" form:"tag_id" gorm:"type:varchar(255); not null;comment:展示标签"`                                //展示标签
	DiFen           string `desc:"底分" json:"di_fen" form:"di_fen" gorm:"type:varchar(255);not null;comment:底分"`                                     //底分
	Fee             string `desc:"门票" json:"Fee" form:"Fee" gorm:"type:varchar(255);not null;comment:门票"`                                           // 门票
	FengDingBeiShu  string `desc:"封顶倍数" json:"FengDingBeiShu" form:"FengDingBeiShu" gorm:"type:varchar(255);not null;comment:封顶倍数"`                 // 封顶倍数
	TuiJianJinChang string `desc:"推荐进场" json:"TuiJianJinChang" form:"TuiJianJinChang" gorm:"type:varchar(255);not null;comment:推荐进场"`               // 推荐进场
	LimitMin        string `desc:"场次下限" json:"LimitMin" form:"LimitMin" gorm:"type:varchar(255);not null;comment:场次下限"`                             // 场次下限
	LimitMax        string `desc:"场次上限" json:"LimitMax" form:"LimitMax" gorm:"type:varchar(255);not null;comment:场次上限"`                             // 场次上限
	PopGift         uint8  `desc:"是否弹出破产礼包" json:"PopGift" form:"PopGift" gorm:"type:tinyint unsigned;size: 1;not null;default 0;comment:是否弹出破产礼包"` // 是否弹出破产礼包
}

// gorm 自定义表名
func (ChangCiModel) TableName() string {
	return "chang_ci"
}
func main() {
	// Connect to MySQL database
	dsn := "root:123456@tcp(127.0.0.1:3306)/lobby?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	}

	// Auto-migrate the "my_users" table
	db.AutoMigrate(&ChangCiModel{})

	db.Create(&ChangCiModel{
		WanId:           1,
		ChangId:         101,
		Name:            "初级场",
		CardId:          "chuji",
		TagId:           "",
		DiFen:           "30",
		Fee:             "30",
		FengDingBeiShu:  "512",
		TuiJianJinChang: "100",
		LimitMin:        "100",
		LimitMax:        "5000",
		PopGift:         1,
	})
	db.Create(&ChangCiModel{
		WanId:           1,
		ChangId:         102,
		Name:            "中级场",
		CardId:          "zhongji",
		TagId:           "",
		DiFen:           "100",
		Fee:             "100",
		FengDingBeiShu:  "512",
		TuiJianJinChang: "1000",
		LimitMin:        "1000",
		LimitMax:        "50000",
		PopGift:         1,
	})
	db.Create(&ChangCiModel{
		WanId:           1,
		ChangId:         103,
		Name:            "高级场",
		CardId:          "gaoji",
		TagId:           "",
		DiFen:           "500",
		Fee:             "500",
		FengDingBeiShu:  "512",
		TuiJianJinChang: "10000",
		LimitMin:        "10000",
		LimitMax:        "500000",
		PopGift:         0,
	})
	db.Create(&ChangCiModel{
		WanId:           1,
		ChangId:         104,
		Name:            "大师场",
		CardId:          "dashi",
		TagId:           "",
		DiFen:           "1000",
		Fee:             "1000",
		FengDingBeiShu:  "512",
		TuiJianJinChang: "100000",
		LimitMin:        "100000",
		LimitMax:        "0",
		PopGift:         0,
	})

	db.Create(&ChangCiModel{
		WanId:           2,
		ChangId:         201,
		Name:            "初级场",
		CardId:          "chuji",
		TagId:           "",
		DiFen:           "30",
		Fee:             "30",
		FengDingBeiShu:  "512",
		TuiJianJinChang: "100",
		LimitMin:        "100",
		LimitMax:        "5000",
		PopGift:         1,
	})
	db.Create(&ChangCiModel{
		WanId:           2,
		ChangId:         202,
		Name:            "中级场",
		CardId:          "zhongji",
		TagId:           "",
		DiFen:           "100",
		Fee:             "100",
		FengDingBeiShu:  "512",
		TuiJianJinChang: "1000",
		LimitMin:        "1000",
		LimitMax:        "50000",
		PopGift:         1,
	})
	db.Create(&ChangCiModel{
		WanId:           2,
		ChangId:         203,
		Name:            "高级场",
		CardId:          "gaoji",
		TagId:           "",
		DiFen:           "500",
		Fee:             "500",
		FengDingBeiShu:  "512",
		TuiJianJinChang: "10000",
		LimitMin:        "10000",
		LimitMax:        "500000",
		PopGift:         0,
	})
	db.Create(&ChangCiModel{
		WanId:           2,
		ChangId:         204,
		Name:            "大师场",
		CardId:          "dashi",
		TagId:           "",
		DiFen:           "1000",
		Fee:             "1000",
		FengDingBeiShu:  "512",
		TuiJianJinChang: "100000",
		LimitMin:        "100000",
		LimitMax:        "0",
		PopGift:         0,
	})

	db.Create(&ChangCiModel{
		WanId:           3,
		ChangId:         301,
		Name:            "初级场",
		CardId:          "chuji",
		TagId:           "",
		DiFen:           "30",
		Fee:             "30",
		FengDingBeiShu:  "512",
		TuiJianJinChang: "100",
		LimitMin:        "100",
		LimitMax:        "5000",
		PopGift:         1,
	})
	db.Create(&ChangCiModel{
		WanId:           3,
		ChangId:         302,
		Name:            "中级场",
		CardId:          "zhongji",
		TagId:           "",
		DiFen:           "100",
		Fee:             "100",
		FengDingBeiShu:  "512",
		TuiJianJinChang: "1000",
		LimitMin:        "1000",
		LimitMax:        "50000",
		PopGift:         1,
	})
	db.Create(&ChangCiModel{
		WanId:           3,
		ChangId:         303,
		Name:            "高级场",
		CardId:          "gaoji",
		TagId:           "",
		DiFen:           "500",
		Fee:             "500",
		FengDingBeiShu:  "512",
		TuiJianJinChang: "10000",
		LimitMin:        "10000",
		LimitMax:        "500000",
		PopGift:         0,
	})
	db.Create(&ChangCiModel{
		WanId:           3,
		ChangId:         304,
		Name:            "大师场",
		CardId:          "dashi",
		TagId:           "",
		DiFen:           "1000",
		Fee:             "1000",
		FengDingBeiShu:  "512",
		TuiJianJinChang: "100000",
		LimitMin:        "100000",
		LimitMax:        "0",
		PopGift:         0,
	})

	db.Create(&ChangCiModel{
		WanId:           4,
		ChangId:         401,
		Name:            "初级场",
		CardId:          "chuji",
		TagId:           "",
		DiFen:           "30",
		Fee:             "30",
		FengDingBeiShu:  "512",
		TuiJianJinChang: "100",
		LimitMin:        "100",
		LimitMax:        "5000",
		PopGift:         1,
	})
	db.Create(&ChangCiModel{
		WanId:           4,
		ChangId:         402,
		Name:            "中级场",
		CardId:          "zhongji",
		TagId:           "",
		DiFen:           "100",
		Fee:             "100",
		FengDingBeiShu:  "512",
		TuiJianJinChang: "1000",
		LimitMin:        "1000",
		LimitMax:        "50000",
		PopGift:         1,
	})
	db.Create(&ChangCiModel{
		WanId:           4,
		ChangId:         403,
		Name:            "高级场",
		CardId:          "gaoji",
		TagId:           "",
		DiFen:           "500",
		Fee:             "500",
		FengDingBeiShu:  "512",
		TuiJianJinChang: "10000",
		LimitMin:        "10000",
		LimitMax:        "500000",
		PopGift:         0,
	})
	db.Create(&ChangCiModel{
		WanId:           4,
		ChangId:         404,
		Name:            "大师场",
		CardId:          "dashi",
		TagId:           "",
		DiFen:           "1000",
		Fee:             "1000",
		FengDingBeiShu:  "512",
		TuiJianJinChang: "100000",
		LimitMin:        "100000",
		LimitMax:        "0",
		PopGift:         0,
	})

	db.Create(&ChangCiModel{
		WanId:           5,
		ChangId:         501,
		Name:            "初级场",
		CardId:          "chuji",
		TagId:           "",
		DiFen:           "30",
		Fee:             "30",
		FengDingBeiShu:  "512",
		TuiJianJinChang: "100",
		LimitMin:        "100",
		LimitMax:        "5000",
		PopGift:         1,
	})
	db.Create(&ChangCiModel{
		WanId:           5,
		ChangId:         502,
		Name:            "中级场",
		CardId:          "zhongji",
		TagId:           "",
		DiFen:           "100",
		Fee:             "100",
		FengDingBeiShu:  "512",
		TuiJianJinChang: "1000",
		LimitMin:        "1000",
		LimitMax:        "50000",
		PopGift:         1,
	})
	db.Create(&ChangCiModel{
		WanId:           5,
		ChangId:         503,
		Name:            "高级场",
		CardId:          "gaoji",
		TagId:           "",
		DiFen:           "500",
		Fee:             "500",
		FengDingBeiShu:  "512",
		TuiJianJinChang: "10000",
		LimitMin:        "10000",
		LimitMax:        "500000",
		PopGift:         0,
	})
	db.Create(&ChangCiModel{
		WanId:           5,
		ChangId:         504,
		Name:            "大师场",
		CardId:          "dashi",
		TagId:           "",
		DiFen:           "1000",
		Fee:             "1000",
		FengDingBeiShu:  "512",
		TuiJianJinChang: "100000",
		LimitMin:        "100000",
		LimitMax:        "0",
		PopGift:         0,
	})

	db.Create(&ChangCiModel{
		WanId:           6,
		ChangId:         601,
		Name:            "初级场",
		CardId:          "chuji",
		TagId:           "",
		DiFen:           "30",
		Fee:             "30",
		FengDingBeiShu:  "512",
		TuiJianJinChang: "100",
		LimitMin:        "100",
		LimitMax:        "5000",
		PopGift:         1,
	})
	db.Create(&ChangCiModel{
		WanId:           6,
		ChangId:         602,
		Name:            "中级场",
		CardId:          "zhongji",
		TagId:           "",
		DiFen:           "100",
		Fee:             "100",
		FengDingBeiShu:  "512",
		TuiJianJinChang: "1000",
		LimitMin:        "1000",
		LimitMax:        "50000",
		PopGift:         1,
	})
	db.Create(&ChangCiModel{
		WanId:           6,
		ChangId:         603,
		Name:            "高级场",
		CardId:          "gaoji",
		TagId:           "",
		DiFen:           "500",
		Fee:             "500",
		FengDingBeiShu:  "512",
		TuiJianJinChang: "10000",
		LimitMin:        "10000",
		LimitMax:        "500000",
		PopGift:         0,
	})
	db.Create(&ChangCiModel{
		WanId:           6,
		ChangId:         604,
		Name:            "大师场",
		CardId:          "dashi",
		TagId:           "",
		DiFen:           "1000",
		Fee:             "1000",
		FengDingBeiShu:  "512",
		TuiJianJinChang: "100000",
		LimitMin:        "100000",
		LimitMax:        "0",
		PopGift:         0,
	})

	db.Create(&ChangCiModel{
		WanId:           7,
		ChangId:         701,
		Name:            "初级场",
		CardId:          "chuji",
		TagId:           "",
		DiFen:           "30",
		Fee:             "30",
		FengDingBeiShu:  "512",
		TuiJianJinChang: "100",
		LimitMin:        "100",
		LimitMax:        "5000",
		PopGift:         1,
	})
	db.Create(&ChangCiModel{
		WanId:           7,
		ChangId:         702,
		Name:            "中级场",
		CardId:          "zhongji",
		TagId:           "",
		DiFen:           "100",
		Fee:             "100",
		FengDingBeiShu:  "512",
		TuiJianJinChang: "1000",
		LimitMin:        "1000",
		LimitMax:        "50000",
		PopGift:         1,
	})
	db.Create(&ChangCiModel{
		WanId:           7,
		ChangId:         703,
		Name:            "高级场",
		CardId:          "gaoji",
		TagId:           "",
		DiFen:           "500",
		Fee:             "500",
		FengDingBeiShu:  "512",
		TuiJianJinChang: "10000",
		LimitMin:        "10000",
		LimitMax:        "500000",
		PopGift:         0,
	})
	db.Create(&ChangCiModel{
		WanId:           7,
		ChangId:         704,
		Name:            "大师场",
		CardId:          "dashi",
		TagId:           "",
		DiFen:           "1000",
		Fee:             "1000",
		FengDingBeiShu:  "512",
		TuiJianJinChang: "100000",
		LimitMin:        "100000",
		LimitMax:        "0",
		PopGift:         0,
	})

	db.Create(&ChangCiModel{
		WanId:           8,
		ChangId:         801,
		Name:            "初级场",
		CardId:          "chuji",
		TagId:           "",
		DiFen:           "30",
		Fee:             "30",
		FengDingBeiShu:  "512",
		TuiJianJinChang: "100",
		LimitMin:        "100",
		LimitMax:        "5000",
		PopGift:         1,
	})
	db.Create(&ChangCiModel{
		WanId:           8,
		ChangId:         802,
		Name:            "中级场",
		CardId:          "zhongji",
		TagId:           "",
		DiFen:           "100",
		Fee:             "100",
		FengDingBeiShu:  "512",
		TuiJianJinChang: "1000",
		LimitMin:        "1000",
		LimitMax:        "50000",
		PopGift:         1,
	})
	db.Create(&ChangCiModel{
		WanId:           8,
		ChangId:         803,
		Name:            "高级场",
		CardId:          "gaoji",
		TagId:           "",
		DiFen:           "500",
		Fee:             "500",
		FengDingBeiShu:  "512",
		TuiJianJinChang: "10000",
		LimitMin:        "10000",
		LimitMax:        "500000",
		PopGift:         0,
	})
	db.Create(&ChangCiModel{
		WanId:           8,
		ChangId:         804,
		Name:            "大师场",
		CardId:          "dashi",
		TagId:           "",
		DiFen:           "1000",
		Fee:             "1000",
		FengDingBeiShu:  "512",
		TuiJianJinChang: "100000",
		LimitMin:        "100000",
		LimitMax:        "0",
		PopGift:         0,
	})

	db.Create(&ChangCiModel{
		WanId:           9,
		ChangId:         901,
		Name:            "初级场",
		CardId:          "chuji",
		TagId:           "",
		DiFen:           "30",
		Fee:             "30",
		FengDingBeiShu:  "512",
		TuiJianJinChang: "100",
		LimitMin:        "100",
		LimitMax:        "5000",
		PopGift:         1,
	})
	db.Create(&ChangCiModel{
		WanId:           9,
		ChangId:         902,
		Name:            "中级场",
		CardId:          "zhongji",
		TagId:           "",
		DiFen:           "100",
		Fee:             "100",
		FengDingBeiShu:  "512",
		TuiJianJinChang: "1000",
		LimitMin:        "1000",
		LimitMax:        "50000",
		PopGift:         1,
	})
	db.Create(&ChangCiModel{
		WanId:           9,
		ChangId:         903,
		Name:            "高级场",
		CardId:          "gaoji",
		TagId:           "",
		DiFen:           "500",
		Fee:             "500",
		FengDingBeiShu:  "512",
		TuiJianJinChang: "10000",
		LimitMin:        "10000",
		LimitMax:        "500000",
		PopGift:         0,
	})
	db.Create(&ChangCiModel{
		WanId:           9,
		ChangId:         904,
		Name:            "大师场",
		CardId:          "dashi",
		TagId:           "",
		DiFen:           "1000",
		Fee:             "1000",
		FengDingBeiShu:  "512",
		TuiJianJinChang: "100000",
		LimitMin:        "100000",
		LimitMax:        "0",
		PopGift:         0,
	})

	db.Create(&ChangCiModel{
		WanId:           10,
		ChangId:         1001,
		Name:            "初级场",
		CardId:          "chuji",
		TagId:           "",
		DiFen:           "30",
		Fee:             "30",
		FengDingBeiShu:  "512",
		TuiJianJinChang: "100",
		LimitMin:        "100",
		LimitMax:        "5000",
		PopGift:         1,
	})
	db.Create(&ChangCiModel{
		WanId:           10,
		ChangId:         1002,
		Name:            "中级场",
		CardId:          "zhongji",
		TagId:           "",
		DiFen:           "100",
		Fee:             "100",
		FengDingBeiShu:  "512",
		TuiJianJinChang: "1000",
		LimitMin:        "1000",
		LimitMax:        "50000",
		PopGift:         1,
	})
	db.Create(&ChangCiModel{
		WanId:           10,
		ChangId:         1003,
		Name:            "高级场",
		CardId:          "gaoji",
		TagId:           "",
		DiFen:           "500",
		Fee:             "500",
		FengDingBeiShu:  "512",
		TuiJianJinChang: "10000",
		LimitMin:        "10000",
		LimitMax:        "500000",
		PopGift:         0,
	})
	db.Create(&ChangCiModel{
		WanId:           10,
		ChangId:         1004,
		Name:            "大师场",
		CardId:          "dashi",
		TagId:           "",
		DiFen:           "1000",
		Fee:             "1000",
		FengDingBeiShu:  "512",
		TuiJianJinChang: "100000",
		LimitMin:        "100000",
		LimitMax:        "0",
		PopGift:         0,
	})
	db.Create(&ChangCiModel{
		WanId:           11,
		ChangId:         1101,
		Name:            "初级场",
		CardId:          "chuji",
		TagId:           "",
		DiFen:           "30",
		Fee:             "30",
		FengDingBeiShu:  "512",
		TuiJianJinChang: "100",
		LimitMin:        "100",
		LimitMax:        "5000",
		PopGift:         1,
	})
	db.Create(&ChangCiModel{
		WanId:           11,
		ChangId:         1102,
		Name:            "中级场",
		CardId:          "zhongji",
		TagId:           "",
		DiFen:           "100",
		Fee:             "100",
		FengDingBeiShu:  "512",
		TuiJianJinChang: "1000",
		LimitMin:        "1000",
		LimitMax:        "50000",
		PopGift:         1,
	})
	db.Create(&ChangCiModel{
		WanId:           11,
		ChangId:         1103,
		Name:            "高级场",
		CardId:          "gaoji",
		TagId:           "",
		DiFen:           "500",
		Fee:             "500",
		FengDingBeiShu:  "512",
		TuiJianJinChang: "10000",
		LimitMin:        "10000",
		LimitMax:        "500000",
		PopGift:         0,
	})
	db.Create(&ChangCiModel{
		WanId:           11,
		ChangId:         1104,
		Name:            "大师场",
		CardId:          "dashi",
		TagId:           "",
		DiFen:           "1000",
		Fee:             "1000",
		FengDingBeiShu:  "512",
		TuiJianJinChang: "100000",
		LimitMin:        "100000",
		LimitMax:        "0",
		PopGift:         0,
	})

	db.Create(&ChangCiModel{
		WanId:           12,
		ChangId:         1201,
		Name:            "初级场",
		CardId:          "chuji",
		TagId:           "",
		DiFen:           "30",
		Fee:             "30",
		FengDingBeiShu:  "512",
		TuiJianJinChang: "100",
		LimitMin:        "100",
		LimitMax:        "5000",
		PopGift:         1,
	})
	db.Create(&ChangCiModel{
		WanId:           12,
		ChangId:         1202,
		Name:            "中级场",
		CardId:          "zhongji",
		TagId:           "",
		DiFen:           "100",
		Fee:             "100",
		FengDingBeiShu:  "512",
		TuiJianJinChang: "1000",
		LimitMin:        "1000",
		LimitMax:        "50000",
		PopGift:         1,
	})
	db.Create(&ChangCiModel{
		WanId:           12,
		ChangId:         1203,
		Name:            "高级场",
		CardId:          "gaoji",
		TagId:           "",
		DiFen:           "500",
		Fee:             "500",
		FengDingBeiShu:  "512",
		TuiJianJinChang: "10000",
		LimitMin:        "10000",
		LimitMax:        "500000",
		PopGift:         0,
	})
	db.Create(&ChangCiModel{
		WanId:           12,
		ChangId:         1204,
		Name:            "大师场",
		CardId:          "dashi",
		TagId:           "",
		DiFen:           "1000",
		Fee:             "1000",
		FengDingBeiShu:  "512",
		TuiJianJinChang: "100000",
		LimitMin:        "100000",
		LimitMax:        "0",
		PopGift:         0,
	})

	db.Create(&ChangCiModel{
		WanId:           13,
		ChangId:         1301,
		Name:            "初级场",
		CardId:          "chuji",
		TagId:           "",
		DiFen:           "30",
		Fee:             "30",
		FengDingBeiShu:  "512",
		TuiJianJinChang: "100",
		LimitMin:        "100",
		LimitMax:        "5000",
		PopGift:         1,
	})
	db.Create(&ChangCiModel{
		WanId:           13,
		ChangId:         1302,
		Name:            "中级场",
		CardId:          "zhongji",
		TagId:           "",
		DiFen:           "100",
		Fee:             "100",
		FengDingBeiShu:  "512",
		TuiJianJinChang: "1000",
		LimitMin:        "1000",
		LimitMax:        "50000",
		PopGift:         1,
	})
	db.Create(&ChangCiModel{
		WanId:           13,
		ChangId:         1303,
		Name:            "高级场",
		CardId:          "gaoji",
		TagId:           "",
		DiFen:           "500",
		Fee:             "500",
		FengDingBeiShu:  "512",
		TuiJianJinChang: "10000",
		LimitMin:        "10000",
		LimitMax:        "500000",
		PopGift:         0,
	})
	db.Create(&ChangCiModel{
		WanId:           13,
		ChangId:         1304,
		Name:            "大师场",
		CardId:          "dashi",
		TagId:           "",
		DiFen:           "1000",
		Fee:             "1000",
		FengDingBeiShu:  "512",
		TuiJianJinChang: "100000",
		LimitMin:        "100000",
		LimitMax:        "0",
		PopGift:         0,
	})

	db.Create(&ChangCiModel{
		WanId:           14,
		ChangId:         1401,
		Name:            "初级场",
		CardId:          "chuji",
		TagId:           "",
		DiFen:           "30",
		Fee:             "30",
		FengDingBeiShu:  "512",
		TuiJianJinChang: "100",
		LimitMin:        "100",
		LimitMax:        "5000",
		PopGift:         1,
	})
	db.Create(&ChangCiModel{
		WanId:           14,
		ChangId:         1402,
		Name:            "中级场",
		CardId:          "zhongji",
		TagId:           "",
		DiFen:           "100",
		Fee:             "100",
		FengDingBeiShu:  "512",
		TuiJianJinChang: "1000",
		LimitMin:        "1000",
		LimitMax:        "50000",
		PopGift:         1,
	})
	db.Create(&ChangCiModel{
		WanId:           14,
		ChangId:         1403,
		Name:            "高级场",
		CardId:          "gaoji",
		TagId:           "",
		DiFen:           "500",
		Fee:             "500",
		FengDingBeiShu:  "512",
		TuiJianJinChang: "10000",
		LimitMin:        "10000",
		LimitMax:        "500000",
		PopGift:         0,
	})
	db.Create(&ChangCiModel{
		WanId:           14,
		ChangId:         1404,
		Name:            "大师场",
		CardId:          "dashi",
		TagId:           "",
		DiFen:           "1000",
		Fee:             "1000",
		FengDingBeiShu:  "512",
		TuiJianJinChang: "100000",
		LimitMin:        "100000",
		LimitMax:        "0",
		PopGift:         0,
	})

}
