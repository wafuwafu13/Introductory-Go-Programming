package main

import (
	"time"
	"image"
	"log"
)

type RoverDriver struct {
	commandc chan command
}

type command int
const(
	right = command(0)
	left = command(1)
)

// コンストラクタ関数
func NewRoverDriver() *RoverDriver {
	r := &RoverDriver{
		commandc: make(chan command),
	}
	go r.drive()
	return r
}
 
func (r *RoverDriver) Left() {
	r.commandc <- left
}
 
func (r *RoverDriver) Right() {
	r.commandc <- right
}

func (r *RoverDriver) drive() {
	pos := image.Point{X:0, Y:0}
	direction := image.Point{X: 1, Y: 0}
	updateInterval := 250 * time.Millisecond
	nextMove := time.After(updateInterval)
	for {
		select {
		case c := <-r.commandc: // コマンドを待つ
		    switch c {
			case right:
				direction = image.Point{
					X: -direction.Y,
					Y: direction.X,
				}
			case left:
				direction = image.Point{
					X: direction.Y,
					Y: -direction.X,
				}
			}
			log.Printf("new direction %v", direction)
		case <- nextMove:
			pos = pos.Add(direction)
			log.Printf("move to %v", pos)
			nextMove = time.After(updateInterval)
		}
	}
}

// func worker() {
// 	pos := image.Point{X: 10, Y: 10}
// 	direction := image.Point{X: 1, Y: 0}

// 	next := time.After(time.Second) // 最初のタイマチャネルを作る

// 	for {
// 		select {
// 		case <- next: // タイマイベントの発火を待つ
// 			pos = pos.Add(direction)
			
// 			fmt.Println("current position is ", pos)
			
// 			next = time.After(time.Second) // 次のイベントを、別のタイマチャネルで待つ
// 		}
// 	}
// }

func main() {
	r := NewRoverDriver()
	time.Sleep(3 * time.Second) // time.Sleepがなかったら1方向づつにしか進まない
	r.Left()
	time.Sleep(3 * time.Second)
	r.Right()
	time.Sleep(3 * time.Second)
}

//  2020/06/22 08:30:20 move to (1,0)
//  2020/06/22 08:30:20 move to (2,0)
//  2020/06/22 08:30:20 move to (3,0)
//  2020/06/22 08:30:20 move to (4,0)
//  2020/06/22 08:30:21 move to (5,0)
//  2020/06/22 08:30:21 move to (6,0)
//  2020/06/22 08:30:21 move to (7,0)
//  2020/06/22 08:30:21 move to (8,0)
//  2020/06/22 08:30:22 move to (9,0)
//  2020/06/22 08:30:22 move to (10,0)
//  2020/06/22 08:30:22 move to (11,0)
//  2020/06/22 08:30:22 new direction (0,-1)
//  2020/06/22 08:30:22 move to (11,-1)
//  2020/06/22 08:30:23 move to (11,-2)
//  2020/06/22 08:30:23 move to (11,-3)
//  2020/06/22 08:30:23 move to (11,-4)
//  2020/06/22 08:30:23 move to (11,-5)
//  2020/06/22 08:30:24 move to (11,-6)
//  2020/06/22 08:30:24 move to (11,-7)
//  2020/06/22 08:30:24 move to (11,-8)
//  2020/06/22 08:30:24 move to (11,-9)
//  2020/06/22 08:30:25 move to (11,-10)
//  2020/06/22 08:30:25 move to (11,-11)
//  2020/06/22 08:30:25 move to (11,-12)
//  2020/06/22 08:30:25 new direction (1,0)
//  2020/06/22 08:30:25 move to (12,-12)
//  2020/06/22 08:30:26 move to (13,-12)
//  2020/06/22 08:30:26 move to (14,-12)
//  2020/06/22 08:30:26 move to (15,-12)
//  2020/06/22 08:30:26 move to (16,-12)
//  2020/06/22 08:30:27 move to (17,-12)
//  2020/06/22 08:30:27 move to (18,-12)
//  2020/06/22 08:30:27 move to (19,-12)
//  2020/06/22 08:30:27 move to (20,-12)
//  2020/06/22 08:30:28 move to (21,-12)
//  2020/06/22 08:30:28 move to (22,-12)
//  2020/06/22 08:30:28 move to (23,-12)
