package main

import (
	"fmt"
	"math/rand"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db, _ = createDB()
)

// 房间服务模型
type RoomServerModel struct {
	gorm.Model
	Name    string `desc:"名称" json:"name" form:"name" gorm:"type:varchar(255);not null;comment:昵称"`                                                      // 昵称
	IP      string `desc:"IP" json:"ip" form:"ip" gorm:"type:varchar(255); not null;comment:IP地址"`                                                       // 玩法类型
	Port    uint16 `desc:"端口" json:"port" form:"port" gorm:"type:smallint unsigned;size:5; not null;comment:端口"`                                         // 端口
	WanType uint16 `desc:"玩法类型" json:"wan_type" form:"wan_type" gorm:"type:smallint unsigned;size:5;not null;comment:玩法类型"`                              // 玩法类型
	State   uint8  `desc:"服务状态 1:正常,2:手动停止,3:挂掉了" json:"state" form:"state" gorm:"type:tinyint unsigned;size:1;not null;comment:服务状态 1:正常,2:手动停止,3:挂掉了"` // 服务状态
}

func (RoomServerModel) TableName() string {
	return "room_server"
}

// 分页查询结果
type PageResult struct {
	Items      []*RoomServerModel `json:"items"`
	TotalCount int64              `json:"total_count"`
}

// 创建并初始化数据库连接
func createDB() (*gorm.DB, error) {
	dsn := "root:123456@tcp(127.0.0.1:3306)/lobby?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		return nil, err
	}
	// 自动迁移数据库结构
	db.AutoMigrate(&RoomServerModel{})
	return db, nil
}

// 插入数据
func insertRoomServer(db *gorm.DB) {
	// roomServer := RoomServerModel{
	// 	Name:    "红中血战",
	// 	IP:      "192.168.0.219",
	// 	Port:    6666,
	// 	WanType: 4,
	// 	State:   1,
	// }
	// db.Create(&roomServer)
}

// 查询数据
func queryRoomServer(db *gorm.DB, id uint) {
	var roomServer RoomServerModel
	db.First(&roomServer, id)
	fmt.Printf("RoomServer %d: %+v\n", id, roomServer)
}

// 更新数据
func updateRoomServer(db *gorm.DB, id uint) {
	var roomServer RoomServerModel
	db.First(&roomServer, id)
	roomServer.State = 2
	db.Save(&roomServer)
}

// 删除数据
func deleteRoomServer(db *gorm.DB, id uint) {
	var roomServer RoomServerModel
	db.First(&roomServer, id)
	db.Delete(&roomServer)
}

// 分页查询
func listRoomServers(db *gorm.DB, pageIndex, pageSize int) (*PageResult, error) {
	var total int64
	result := &PageResult{
		Items:      []*RoomServerModel{},
		TotalCount: 0,
	}
	db.Model(&RoomServerModel{}).Count(&total)
	db.Offset((pageIndex - 1) * pageSize).Limit(pageSize).Order("id desc").Find(&result.Items)
	result.TotalCount = total

	return result, nil
}

func FindRoomServersByType(typ uint16) (result []RoomServerModel) {
	db.Where("wan_type = ? AND state = ?", typ, 1).Find(&result)
	return
}
func FindRoomServers() (result []RoomServerModel) {
	db.Where("state = ?", 1).Find(&result)
	return
}

func main() {
	// insertRoomServer(db)
	// servers := FindRoomServersByType(1)
	servers := FindRoomServers()
	index := rand.Intn(len(servers))
	server := servers[index]
	fmt.Printf("servers = %#v\n", server)
}

// 主函数
// func main() {
// 	db, err := createDB()
// 	if err != nil {
// 		panic(err)
// 	}
// 	gormDB, err := db.DB()
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer gormDB.Close()

// 	insertRoomServer(db)
// 	queryRoomServer(db, 1)
// 	updateRoomServer(db, 1)
// 	queryRoomServer(db, 1)
// 	deleteRoomServer(db, 1)
// 	_, err = listRoomServers(db, 1, 10)
// 	if err != nil {
// 		panic(err)
// 	}

// 	fmt.Println("Done")
// }
