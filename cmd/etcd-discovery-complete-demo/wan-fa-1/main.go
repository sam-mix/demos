package main

// import (
// 	"context"
// 	"fmt"
// 	"log"

// 	_ "github.com/go-sql-driver/mysql"
// 	"github.com/jmoiron/sqlx"
// 	"github.com/smallnest/rpcx/client"
// 	"github.com/smallnest/rpcx/server"
// )

// type User struct {
// 	ID   int
// 	Name string
// 	Age  int
// }

// type UserService struct {
// 	db *sqlx.DB
// }

// func NewUserService(db *sqlx.DB) *UserService {
// 	return &UserService{
// 		db: db,
// 	}
// }

// func (u *UserService) CreateUser(ctx context.Context, user *User) error {
// 	_, err := u.db.NamedExec(`
//         INSERT INTO users (name, age)
//         VALUES (:name, :age)
//     `, user)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

// func (u *UserService) ListUsers(ctx context.Context) ([]User, error) {
// 	var users []User
// 	err := u.db.Select(&users, "SELECT * FROM users")
// 	if err != nil {
// 		return nil, err
// 	}
// 	return users, nil
// }

// func (u *UserService) UpdateUser(ctx context.Context, id int, user *User) error {
// 	_, err := u.db.NamedExec(`
//         UPDATE users SET name=:name, age=:age WHERE id=:id
//     `, map[string]interface{}{
// 		"name": user.Name,
// 		"age":  user.Age,
// 		"id":   id,
// 	})
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

// func (u *UserService) DeleteUser(ctx context.Context, id int) error {
// 	_, err := u.db.Exec("DELETE FROM users WHERE id=?", id)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

// func main() {
// 	// 连接MySQL数据库
// 	db, err := sqlx.Connect("mysql", "root:123456@tcp(127.0.0.1:3306)/test")
// 	if err != nil {
// 		log.Fatal("Failed to connect to database: ", err)
// 	}
// 	defer db.Close()

// 	// 注册服务并启动服务器
// 	s := server.NewServer()
// 	s.RegisterName("UserService", NewUserService(db), "")
// 	s.Serve("tcp", "127.0.0.1:8888")

// 	// 创建客户端并调用API
// 	xclient := client.NewXClient("UserService", client.Failover, client.RandomSelect, []*client.KVPair{
// 		{Key: "127.0.0.1:8888", Value: ""},
// 	}, client.DefaultOption)
// 	defer xclient.Close()

// 	ctx := context.Background()

// 	// 新增用户
// 	user := &User{
// 		Name: "Alice",
// 		Age:  20,
// 	}
// 	err = xclient.Call(ctx, "CreateUser", user, nil)
// 	if err != nil {
// 		log.Fatal("Failed to create user: ", err)
// 	}

// 	// 获取用户列表
// 	var users []User
// 	err = xclient.Call(ctx, "ListUsers", nil, &users)
// 	if err != nil {
// 		log.Fatal("Failed to list users: ", err)
// 	}
// 	fmt.Println(users)

// 	// 修改用户名为Bob
// 	err = xclient.Call(ctx, "UpdateUser", 1, &User{Name: "Bob"})
// 	if err != nil {
// 		log.Fatal("Failed to update user: ", err)
// 	}

// 	// 删除用户
// 	err = xclient.Call(ctx, "DeleteUser", 1, nil)
// 	if err != nil {
// 		log.Fatal("Failed to delete user: ", err)
// 	}
// }
