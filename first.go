package main

import (
	"fmt"
	"log"
	"github.com/jroimartin/gocui"
	"math/rand"
    "time"
)

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

var d Dungeon
var laura Actor = Actor{x: 2, y: 2, alive: true, health: 10, attack: 5}

func (d *Dungeon)printFloor(v *gocui.View){
	d.floor[laura.y][laura.x] = "@"
	for i := 0; i < d.floorVert; i++ {
		for j := 0; j < d.floorHoriz; j++ {
			fmt.Fprint(v, d.floor[i][j])
		}
		fmt.Fprint(v, "\n")
		
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
	d.genFloor()
	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		log.Panicln(err)
	}
	defer g.Close()

	g.SetManagerFunc(layout)
	
	if err := keybindings(g); err != nil {
		log.Panicln(err)
	}
	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}


}


func layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()
	v, err := g.SetView("Toilet", 0, 0, maxX-1, maxY-1)
	if err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
	}
	v.Clear()
	d.printFloor(v)
	return nil
}


func keybindings(g * gocui.Gui) error {
	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		return err
	}
	if err := g.SetKeybinding("", gocui.KeyArrowUp, gocui.ModNone, laura.moveW); err != nil {
		return err
	}
	if err := g.SetKeybinding("", gocui.KeyArrowLeft, gocui.ModNone, laura.moveA); err != nil {
		return err
	}
	if err := g.SetKeybinding("", gocui.KeyArrowDown, gocui.ModNone, laura.moveS); err != nil {
		return err
	}
	if err := g.SetKeybinding("", gocui.KeyArrowRight, gocui.ModNone, laura.moveD); err != nil {
		return err
	}
	return nil
}

func quit(g * gocui.Gui, v * gocui.View) error {
	return gocui.ErrQuit
}

func (l *Actor)moveW(g *gocui.Gui, v *gocui.View) error{
	if d.floor[laura.y - 1][laura.x] == "." {
		d.floor[laura.y][laura.x] = "."
		l.y -= 1
	}
	return nil
}
func (l *Actor)moveA(g *gocui.Gui, v *gocui.View) error{
	if d.floor[laura.y][laura.x - 1] == "." {
		d.floor[laura.y][laura.x] = "."
		l.x -= 1
	}
	return nil
}
func (l *Actor)moveS(g *gocui.Gui, v *gocui.View) error{
	if d.floor[laura.y + 1][laura.x] == "." {
		d.floor[laura.y][laura.x] = "."
		l.y += 1
	}
	return nil
}
func (l *Actor)moveD(g *gocui.Gui, v *gocui.View) error{
	if d.floor[laura.y][laura.x + 1] == "." {
		d.floor[laura.y][laura.x] = "."
		l.x += 1
	}
	return nil
}


