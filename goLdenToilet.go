package main

import "fmt"
import "math/rand"
import "time"


type Dungeon struct {
	floor [][]string
	floorHoriz int
	floorVert int
}

func (d *Dungeon)printFloor(){
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

func main(){
	rand.Seed(time.Now().UnixNano())
	d := Dungeon{}
	d.genFloor()
	//fmt.Println("Horiz:", d.floorHoriz)
	//fmt.Println("Vert:", d.floorVert)
	//fmt.Println(d)
	d.printFloor()
}
