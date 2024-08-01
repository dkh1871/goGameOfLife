package main

import (
	"fmt"
)

type Grid struct {
	x_axis int
	y_axis int
}

func (g *Grid) GetHeight() {
	fmt.Print("Please enter the Grid Hight:")
	// look into take input from users and vaildating them
}

func (g Grid) MakeNewGrid() [][]bool {
	gb := make([][]bool, g.y_axis)

	for i := range gb {
		gb[i] = make([]bool, g.x_axis)
	}

	return gb
}

func main() {
	fmt.Println("Fun times lol!!!S")

	x := Grid{}
	x.x_axis = 20
	x.y_axis = 20

	fmt.Println(x.x_axis)
	fmt.Println(x.y_axis)

	gameboard := x.MakeNewGrid()

	fmt.Println(gameboard)

}
