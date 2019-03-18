package main

import (
	"fmt"
	"github.com/nsf/termbox-go"
	"github.com/qianlnk/log"
	"math/rand"
	"strconv"
	"time"
)

const (
	GSIZE = 4
)

//4*4数组代表十六个格子
type GameCells [GSIZE][GSIZE]int

//创建游戏，每执行一次返回数组实例
func NewGame() *GameCells {
	termbox.Init()                                            //初始化termbox
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault) //清楚缓冲区
	termbox.SetOutputMode(termbox.Output256)                  //设置输出模式
	termbox.SetInputMode(termbox.InputEsc)                    //设置输入模式
	game := new(GameCells)                                    //分配实例内存
	//随机初始化两个元素
	game.Generate() //随机填充一个空白元素
	game.Generate()

	log.SetOutputPath("./game.log")
	log.SetFormatter("logstash")

	log.Info("init", *game)
	game.Draw(false) //绘制填充边框及元素
	return game      //返回实例
}

//上
func (g *GameCells) Up() bool {
	var ok bool
	//自上而下相邻两个元素相同则相加
	for col := 0; col < GSIZE; col++ {
		for row := 0; row < GSIZE-1; {
			add := false
			if g[row][col] != 0 {
				for r := row + 1; r < GSIZE; r++ {
					if g[r][col] == 0 {
						continue
					}
					if g[row][col] == g[r][col] {
						g[row][col] += g[r][col]
						g[r][col] = 0
						ok = true
						row = r + 1
						add = true
					}

					break
				}
			}

			if !add {
				row++
			}
		}
	}
	//纵向跟进填充空白
	for col := 0; col < GSIZE; col++ {
		for row := 0; row < GSIZE-1; row++ {
			if g[row][col] == 0 {
				for r := row + 1; r < GSIZE; r++ {
					if g[r][col] == 0 {
						continue
					}
					g[row][col] = g[r][col]
					g[r][col] = 0
					ok = true
					break
				}
			}
		}
	}

	return ok
}

func (g *GameCells) Down() bool {
	var ok bool
	//自上而下相邻两个元素相同则相加
	for col := GSIZE - 1; col >= 0; col-- {
		for row := GSIZE - 1; row >= 1; {
			add := false
			if g[row][col] != 0 {
				for r := row - 1; r >= 0; r-- {
					if g[r][col] == 0 {
						continue
					}

					if g[row][col] == g[r][col] {
						g[row][col] += g[r][col]
						g[r][col] = 0
						ok = true
						row = r - 1
						add = true
					}

					break
				}
			}

			if !add {
				row--
			}
		}
	}
	//填补空白处
	for col := GSIZE - 1; col >= 0; col-- {
		for row := GSIZE - 1; row >= 1; row-- {
			if g[row][col] == 0 {
				for r := row - 1; r >= 0; r-- {
					if g[r][col] == 0 {
						continue
					}
					g[row][col] = g[r][col]
					g[r][col] = 0
					ok = true
					break
				}
			}
		}
	}

	return ok
}

func (g *GameCells) Left() bool {
	var ok bool
	//自左而右相邻两个元素相同则相加
	for row := 0; row < GSIZE; row++ {
		for col := 0; col < GSIZE-1; {
			add := false
			if g[row][col] != 0 {
				for c := col + 1; c < GSIZE; c++ {
					if g[row][c] == 0 {
						continue
					}

					if g[row][col] == g[row][c] {
						g[row][col] += g[row][c]
						g[row][c] = 0
						ok = true
						col = c + 1
						add = true
					}

					break
				}
			}

			if !add {
				col++
			}
		}
	}
	//填充空白元素
	for row := 0; row < GSIZE; row++ {
		for col := 0; col < GSIZE; col++ {
			if g[row][col] == 0 {
				for c := col + 1; c < GSIZE; c++ {
					if g[row][c] == 0 {
						continue
					}
					g[row][col] = g[row][c]
					g[row][c] = 0
					ok = true
					break
				}
			}
		}
	}

	return ok
}

func (g *GameCells) Right() bool {
	var ok bool
	//自右而左相邻两个元素相同则相加
	for row := GSIZE - 1; row >= 0; row-- {
		for col := GSIZE - 1; col >= 1; {
			add := false
			if g[row][col] != 0 {
				for c := col - 1; c >= 0; c-- {
					if g[row][c] == 0 {
						continue
					}

					if g[row][col] == g[row][c] {
						g[row][col] += g[row][c]
						g[row][c] = 0
						ok = true
						col = c - 1
						add = true
					}

					break
				}
			}

			if !add {
				col--
			}
		}
	}
	//填充空白元素
	for row := GSIZE - 1; row >= 0; row-- {
		for col := GSIZE - 1; col >= 1; col-- {
			if g[row][col] == 0 {
				for c := col - 1; c >= 0; c-- {
					if g[row][c] == 0 {
						continue
					}

					g[row][col] = g[row][c]
					g[row][c] = 0
					ok = true
					break
				}
			}
		}
	}

	return ok
}

//判断游戏是否结束
func (g *GameCells) GameOver() bool {
	//无空白或横向相邻友相等元素，则游戏继续
	for row := 0; row < GSIZE; row++ {
		for col := 0; col < GSIZE-1; col++ {
			//出现2048,游戏结束
			if g[row][col] == 2048 || g[row][col+1] == 2048 {
				return true
			}
			if g[row][col] == 0 || g[row][col+1] == 0 {
				return false
			}
			if g[row][col] == g[row][col+1] {
				return false
			}
		}
	}
	//纵向相邻两个相等，则游戏继续
	for col := 0; col < GSIZE; col++ {
		for row := 0; row < GSIZE-1; row++ {
			if g[row][col] == g[row+1][col] {
				return false
			}
		}
	}

	return true
}

//给随机空白地方放回一个随机的2或4（2的概率大于4）
func (g *GameCells) Generate() {
	//定位元素实体
	type point struct {
		row int
		col int
	}

	//获得空白元素
	var emptyCells []point
	for row := 0; row < GSIZE; row++ {
		for col := 0; col < GSIZE; col++ {
			if g[row][col] == 0 {
				emptyCells = append(emptyCells, point{row, col})
			}
		}
	}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	index := r.Intn(len(emptyCells))
	val := r.Intn(10)
	if val >= 8 {
		val = 4
	} else {
		val = 2
	}
	g[emptyCells[index].row][emptyCells[index].col] = val
}

//整体绘制
func (g *GameCells) Draw(gameOver bool) {
	//绘制填充边框
	drawTable(BorderDouble, 0, 0, 32, 80, 4, 4, termbox.ColorRed, termbox.ColorDefault)
	top := 1

	for row := 0; row < GSIZE; row++ {
		left := 1
		for col := 0; col < GSIZE; col++ {
			if g[row][col] != 0 {
				//填充不为空的元素
				drawCell(left, top, 20, 8, strconv.Itoa(g[row][col]))
			}
			left += 20 + 1
		}
		top += 8 + 1
	}
	//结束时的填充色
	if gameOver {
		g.Close()
		fmt.Println("Game Over")
	}

	termbox.Flush() //同步缓冲
}

//操作游戏
func (g *GameCells) Play() {
	var bug bool
	//循环操作直到中段
	for {
		redraw := false
		//阻塞等待监控事件
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			switch ev.Key {
			//esc键，则游戏退出
			case termbox.KeyEsc:
				g.Close()
				return
			case termbox.KeyArrowUp: //上
				if redraw = g.Up(); redraw {
					log.Info("U", *g)
				}
			case termbox.KeyArrowDown: //下
				if redraw = g.Down(); redraw {
					log.Info("D", *g)
				}
			case termbox.KeyArrowLeft: //左
				if redraw = g.Left(); redraw {
					log.Info("L", *g)
				}
			case termbox.KeyArrowRight: //右
				if redraw = g.Right(); redraw {
					log.Info("R", *g)
				}
			case termbox.KeyEnter:
				bug = true
			}
		}
		if bug == true {
			g.Draw(true) //绘制画布
			break
		}
		//操作有效
		if redraw {
			g.Generate() //随机增加一个元素
			log.Info("G", *g)
			g.Draw(g.GameOver()) //绘制画布

		}
	}
}

//关闭游戏
func (g *GameCells) Close() {
	termbox.Close() //结束服务
}
