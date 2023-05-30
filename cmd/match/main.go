// 要实现四个玩家的游戏匹配，我们可以使用 Go 语言的协程和通道（goroutine 和 channel）。具体实现步骤如下：

// 1. 首先定义一个 `player` 结构体，包含玩家的 ID 和当前状态（匹配中、游戏中、等待中等）。

// 2. 定义一个名为 `matchPlayers` 的函数，该函数接收玩家列表作为参数，并使用协程（goroutine）和通道（channel）实现玩家匹配功能。

// 3. 在 `matchPlayers` 函数中使用 for 循环遍历玩家列表，每次遍历时向通道（channel）中发送该玩家信息，表示需要匹配对手。

// 4. 接着使用 select 语句从通道中读取两个玩家，表示匹配成功。如果通道中没有足够的玩家，则休眠一段时间后继续读取，以免进入死循环。

// 5. 匹配成功后，将两个玩家从列表中删除，并将他们的状态改为游戏中，然后启动一个协程来模拟两个玩家的游戏过程。

// 下面是实现这个功能的示例代码：

// ```go
// package main

// import (
//     "fmt"
//     "math/rand"
//     "time"
// )

// // 玩家结构体，包含玩家ID和当前状态
// type player struct {
//     id     int
//     status int // 0:等待中, 1:匹配中, 2:游戏中
// }

// func matchPlayers(players []*player) {
//     matched := make(chan *player)
//     for _, p := range players {
//         p.status = 1 // 修改状态为匹配中
//         go func(p1 *player) {
//             for _, p2 := range players {
//                 if p1 != p2 && p2.status == 1 {
//                     p1.status = 2 // 修改状态为游戏中
//                     p2.status = 2
//                     fmt.Printf("Matched player %d and player %d.\n", p1.id, p2.id)
//                     time.Sleep(2 * time.Second) // 模拟游戏过程，这里休眠2秒钟
//                     fmt.Printf("Player %d and player %d finished the game.\n", p1.id, p2.id)
//                     matched <- p1
//                     matched <- p2
//                     return
//                 }
//             }
//             p1.status = 1 // 恢复状态为匹配中
//         }(p)
//     }

//     for {
//         select {
//         case p1 := <-matched:
//             // 将玩家从列表中删除
//             for i, p2 := range players {
//                 if p2 == p1 {
//                     players = append(players[:i], players[i+1:]...)
//                     break
//                 }
//             }
//         case <-time.After(500 * time.Millisecond):
//             // 在没有匹配成功的情况下，每次等待0.5秒钟再继续匹配
//         }

//         // 如果只剩下一名玩家，结束匹配
//         if len(players) == 1 {
//             fmt.Printf("Player %d has no opponent and quit the game.\n", players[0].id)
//             return
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

// 在上面的代码中，我们通过使用 go 关键字启动一个匿名协程，实现玩家匹配功能。在匹配过程中，使用通道（channel）来传递匹配结果，并在匹配成功后启动一个协程模拟两个玩家的游戏过程。如果只剩下一名玩家，则结束匹配。

package main

import (
	"fmt"
	"time"
)

// 玩家结构体，包含玩家ID和当前状态
type player struct {
	id     int
	status int // 0:等待中, 1:匹配中, 2:游戏中
}

func matchPlayers(players []*player) {
	matched := make(chan *player)
	for _, p := range players {
		p.status = 1 // 修改状态为匹配中
		go func(p1 *player) {
			for _, p2 := range players {
				if p1 != p2 && p2.status == 1 {
					p1.status = 2 // 修改状态为游戏中
					p2.status = 2
					fmt.Printf("Matched player %d and player %d.\n", p1.id, p2.id)
					time.Sleep(2 * time.Second) // 模拟游戏过程，这里休眠2秒钟
					fmt.Printf("Player %d and player %d finished the game.\n", p1.id, p2.id)
					matched <- p1
					matched <- p2
					return
				}
			}
			p1.status = 1 // 恢复状态为匹配中
		}(p)
	}

	for {
		select {
		case p1 := <-matched:
			// 将玩家从列表中删除
			for i, p2 := range players {
				if p2 == p1 {
					players = append(players[:i], players[i+1:]...)
					break
				}
			}
		case <-time.After(500 * time.Millisecond):
			// 在没有匹配成功的情况下，每次等待0.5秒钟再继续匹配
		}

		// 如果只剩下一名玩家，结束匹配
		if len(players) == 1 {
			fmt.Printf("Player %d has no opponent and quit the game.\n", players[0].id)
			return
		}
	}
}

func main() {

	// 创建4个玩家
	var players []*player
	for i := 1; i <= 4; i++ {
		p := &player{id: i, status: 0}
		players = append(players, p)
	}

	// 开始匹配
	matchPlayers(players)
}
