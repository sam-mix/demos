// 实现四个玩家的游戏匹配，可以用 Go 语言的通道（channel）来实现等待队列。

// 下面是使用通道实现玩家等待队列匹配出四个玩家进行游戏的示例代码：

// ```go
// package main

// import (
//     "fmt"
//     "math/rand"
//     "time"
// )

// type player struct {
//     id     int
//     status int // 0:等待中, 1:匹配中, 2:游戏中
// }

// func matchPlayers(players []*player) {
//     waitQueue := make(chan *player, len(players))
//     for _, p := range players {
//         waitQueue <- p // 将玩家放入等待队列
//     }

//     matchQueue := make(chan []*player, 1) // 匹配结果队列
//     go func() {
//         for {
//             // 每次从等待队列中取出四名玩家进行匹配
//             matchedPlayers := make([]*player, 0, 4)
//             for i := 0; i < 4; i++ {
//                 matchedPlayer := <-waitQueue
//                 matchedPlayer.status = 1 // 修改状态为匹配中
//                 matchedPlayers = append(matchedPlayers, matchedPlayer)
//             }

//             time.Sleep(2 * time.Second) // 模拟游戏过程，这里休眠2秒钟

//             for _, p := range matchedPlayers {
//                 p.status = 2 // 修改状态为游戏中
//             }

//             matchQueue <- matchedPlayers // 将匹配结果放入队列中
//         }
//     }()

//     for {
//         select {
//         case matchedPlayers := <-matchQueue:
//             // 处理匹配结果
//             fmt.Printf("Matched players %d, %d, %d, %d.\n", matchedPlayers[0].id, matchedPlayers[1].id, matchedPlayers[2].id, matchedPlayers[3].id)
//         }
//     }
// }

// func main() {
//     rand.Seed(time.Now().UnixNano())

//     // 创建4个玩家
//     var players []*player
//     for i := 1; i <= 4; i++ {
//         p := &player{id: i, status: 0}
//         players = append(players, p)
//     }

//     // 开始匹配
//     matchPlayers(players)
// }
// ```

// 在上面的代码中，我们使用一个容量为玩家数量的通道（channel）作为等待队列 `waitQueue`。使用 for 循环将所有玩家放入等待队列中。然后，我们创建一个可以容纳四名玩家的切片 `matchedPlayers`，每次从等待队列中取出四名玩家进行匹配。

// 在游戏过程中，我们休眠 2 秒钟模拟游戏进行过程。匹配成功后，将四名玩家的状态都修改为游戏中，并将匹配结果放入 `matchQueue` 队列中。

// 在主函数中，我们创建了四个玩家，并调用 `matchPlayers` 函数开始匹配。在匹配过程中，使用 select 语句从 `matchQueue` 中读取匹配结果。由于这是一个无限循环，所以一旦匹配成功，就会一直处理匹配结果。

package main

import (
	"fmt"
	"time"
)

type player struct {
	id     int
	status int // 0:等待中, 1:匹配中, 2:游戏中
}

var (
	waitQueue = make(chan *player, 200)
)

func matchPlayers() {
	matchQueue := make(chan []*player, 1) // 匹配结果队列
	defer close(matchQueue)
	go func() {
		for {
			// 每次从等待队列中取出四名玩家进行匹配
			matchedPlayers := make([]*player, 0, 4)
			for i := 0; i < 4; i++ {
				matchedPlayer := <-waitQueue
				matchedPlayer.status = 1 // 修改状态为匹配中
				matchedPlayers = append(matchedPlayers, matchedPlayer)
			}

			time.Sleep(2 * time.Second) // 模拟游戏过程，这里休眠2秒钟

			for _, p := range matchedPlayers {
				p.status = 2 // 修改状态为游戏中
			}

			matchQueue <- matchedPlayers // 将匹配结果放入队列中
		}
	}()

	for matchedPlayers := range matchQueue {
		// 处理匹配结果
		fmt.Printf("Matched players %d, %d, %d, %d.\n", matchedPlayers[0].id, matchedPlayers[1].id, matchedPlayers[2].id, matchedPlayers[3].id)
	}
}

func main() {

	// 创建4个玩家
	go func() {
		for i := 1; ; i++ {
			waitQueue <- &player{id: i, status: 0}

		}
	}()

	// 开始匹配
	matchPlayers()

	select {}
}
