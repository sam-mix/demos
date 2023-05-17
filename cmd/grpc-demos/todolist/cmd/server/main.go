package main

// import (
// 	"context"
// 	"net"

// 	"github.com/jinzhu/gorm"
// 	"github.com/redis/go-redis/v9"
// 	"google.golang.org/grpc"
// 	"google.golang.org/grpc/reflection"

// 	pb "demos/grpc-demos/todolist"
// )

// var db *gorm.DB
// var redisClient *redis.Client

// type server struct{}

// func (s *server) AddItem(ctx context.Context, in *pb.Item) (*pb.ItemId, error) {
// 	item := &Item{
// 		Title:       in.Title,
// 		Description: in.Description,
// 		IsDone:      in.IsDone,
// 	}
// 	db.Create(item)
// 	return &pb.ItemId{Value: int32(item.ID)}, nil
// }

// func (s *server) GetItem(ctx context.Context, in *pb.ItemId) (*pb.Item, error) {
// 	item := &Item{}
// 	db.First(item, in.Value)
// 	return &pb.Item{
// 		Title:       item.Title,
// 		Description: item.Description,
// 		IsDone:      item.IsDone,
// 	}, nil
// }

// func (s *server) GetAllItems(ctx context.Context, in *pb.Empty) (*pb.ItemList, error) {
// 	items := []Item{}
// 	db.Find(&items)
// 	item_list := &pb.ItemList{}
// 	for _, item := range items {
// 		item_list.Items = append(item_list.Items, &pb.Item{
// 			Title:       item.Title,
// 			Description: item.Description,
// 			IsDone:      item.IsDone,
// 		})
// 	}
// 	return item_list, nil
// }

// func (s *server) DeleteItem(ctx context.Context, in *pb.ItemId) (*pb.Empty, error) {
// 	item := &Item{}
// 	db.First(item, in.Value)
// 	db.Delete(item)
// 	return &pb.Empty{}, nil
// }

// func main() {
// 	// Connect to Redis
// 	redisClient = redis.NewClient(&redis.Options{
// 		Addr: "localhost:6379",
// 	})

// 	// Connect to database
// 	db, err := gorm.Open("sqlite3", "todolist.db")
// 	if err != nil {
// 		panic("failed to connect database")
// 	}
// 	defer db.Close()

// 	// Auto Migrate database
// 	db.AutoMigrate(&Item{})

// 	// Create GRPC server
// 	lis, err := net.Listen("tcp", ":50051")
// 	if err != nil {
// 		panic("failed to listen")
// 	}
// 	s := grpc.NewServer()
// 	pb.RegisterTodoListServer(s, &server{})
// 	reflection.Register(s)

// 	if err := s.Serve(lis); err != nil {
// 		panic("failed to serve")
// 	}
// }
