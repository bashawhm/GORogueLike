package main

import (
	"fmt"
	"math/rand"
	"time"
	"os"
	"bufio"
	"strings"
)

const numMonsters = 2

//Global info
var d Dungeon
var laura Actor = Actor{x: 1, y: 1, alive: true, health: 10, attack: 5}

type Actor struct {
	x int
	y int
	alive bool
	health uint16
	attack uint16
}

type Dungeon struct {
	floor [][]string
	floorHoriz int
	floorVert int
}

func (d *Dungeon)printFloor(){
	d.floor[laura.y][laura.x] = "@"
	for i := 0; i < d.floorVert; i++ {
		for j := 0; j < d.floorHoriz; j++ {
			fmt.Print(d.floor[i][j])
		}
		fmt.Print("\n")
	}
}

func (d *Dungeon)genFloor() {
	d.floorHoriz = 5 + (rand.Int()%20)
	d.floorVert = 5 + (rand.Int()%10)
	newFloor := make([][]string, d.floorVert)
	for i := 0; i < d.floorVert; i++ {
		newFloor[i] = make([]string, d.floorHoriz)
	}
	d.floor = newFloor
	for i := 0; i < d.floorVert; i++ {
         for j := 0; j < d.floorHoriz; j++ {
			d.floor[i][j] = "."
            if i == 0 {
				d.floor[i][j] = "#"
			}
			if j == 0 {
				d.floor[i][j] = "#"
			}
			if j == (d.floorHoriz-1){
				d.floor[i][j] = "#"
			}
			if i == (d.floorVert-1){
				d.floor[i][j] = "#"
			}
         }
     }


}

func move(input string, thing *Actor){
	switch input {
	case "w":
		if d.floor[(thing.y - 1)][thing.x] == "."{
			d.floor[thing.y][thing.x] = "."
			thing.y -= 1
		} else {
			fmt.Println("Cannot move up")
		}
	case "s":
		if d.floor[(thing.y + 1)][thing.x] == "."{
			d.floor[thing.y][thing.x] = "."
			thing.y += 1
		} else {
			fmt.Println("Cannot move down")
		}
	case "a":
		if d.floor[thing.y][(thing.x - 1)] == "."{
			d.floor[thing.y][thing.x] = "."
			thing.x -= 1
		} else {
			fmt.Println("Cannot move left")
		}
	case "d":
		if d.floor[thing.y][(thing.x + 1)] == "."{
			d.floor[thing.y][thing.x] = "."
			thing.x += 1
		} else {
			fmt.Println("Cannot move right")
		}
	}

}

func parse(input string){
	input = strings.TrimRight(input, "\n")
	switch input {
	case "exit":
		os.Exit(1)
	case "w":
		move(input, &laura)
	case "s":
		move(input, &laura)
	case "a":
		move(input, &laura)
	case "d":
		move(input, &laura)
	default:
		fmt.Println("Failed to parse: ", input)
	}

}

func main(){
	monsters := make([]Actor, numMonsters)
	for i := 0; i < numMonsters; i++{
		monsters[i].alive = true
		monsters[i].health = 10
		monsters[i].attack = 2
	}
	rand.Seed(time.Now().UnixNano())
	reader := bufio.NewReader(os.Stdin)
	d.genFloor()
	for laura.alive {
		d.printFloor()
		text, _ := reader.ReadString('\n')
		parse(text)
	}

}
