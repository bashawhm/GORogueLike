package main

import (
	"fmt"
	"log"
	"github.com/jroimartin/gocui"
	"math/rand"
	"time"
	"os"
)
const MONSTERNUM = 2
const NEAR = 5

type Actor struct {
	x int
	y int
	gold uint16
	alive bool
	health int
	attack int
}

type Dungeon struct {
	floor [][]string
	floorHoriz int
	floorVert int
	goldX int
	goldY int
}

var d Dungeon
var laura Actor = Actor{x: 1, y: 1, gold: 0, alive: true, health: 10, attack: 5}
var monster [MONSTERNUM]Actor


func spawnMob(){
	for i := 0; i < MONSTERNUM; i++ {
		monster[i].alive = true
		monster[i].health = 2
		monster[i].attack = 4
	}
}

func (d *Dungeon)printFloor(g *gocui.Gui, v *gocui.View){
	for i := 0; i < d.floorVert; i++ {
		for j := 0; j < d.floorHoriz; j++ {
			fmt.Fprint(v, d.floor[i][j])
		}
		fmt.Fprint(v, "\n")
		
	}
}
func (d *Dungeon)genFloor() {
	spawnMob()
	laura.x = 1
	laura.y = 1
	d.floorHoriz = 10 + (rand.Int()%20)
	d.floorVert = 10 + (rand.Int()%10)
	newFloor := make([][]string, d.floorVert)
	for i := 0; i < d.floorVert; i++ {
		newFloor[i] = make([]string, d.floorHoriz)
	}
	d.floor = newFloor
	d.floor[laura.x][laura.y] = "@"
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
	 	
	d.goldX = (1 + (rand.Int()%(d.floorVert - 2)))
	d.goldY = (1 + (rand.Int()%(d.floorHoriz - 2)))

	 for laura.x == d.goldX || laura.y == d.goldY || d.goldX == 0 || d.goldY == d.floorVert + 1 || d.goldY == 0 || d.goldY == d.floorHoriz + 1 {
	 	d.goldX = (1 + (rand.Int()%(d.floorVert - 2)))
	 	d.goldY = (1 + (rand.Int()%(d.floorHoriz - 2)))
	 }
	 d.floor[d.goldX][d.goldY] = "G"

	 for i := 0 ; i < MONSTERNUM; i++ {
		randX := (1 + (rand.Int()%(d.floorVert - 2)))
		randY := (1 + (rand.Int()%(d.floorHoriz - 2)))
		for d.floor[randX][randY] != "." {
			randX = (1 + (rand.Int()%(d.floorVert - 2)))
			randY = (1 + (rand.Int()%(d.floorHoriz - 2)))
		}
		monster[i].x = randX
		monster[i].y = randY
		d.floor[monster[i].x][monster[i].y] = "M"
	 }

}


func main(){
	rand.Seed(time.Now().UnixNano())
	d.genFloor()
	spawnMob()

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
	if laura.gold == 7 {
		g.Close()
		fmt.Println("You melt down all the gold pieces into a new Golden Toilet\nYou've got your Golden Toilet back!!!\n")
		os.Exit(1)
	}
	if !laura.alive {
		g.Close()
		fmt.Println("You Failed!\nYour toilet is no more")
		os.Exit(1)
	}
	v.Clear()
	d.floor[laura.x][laura.y] = "@"
	MonsterAI()
	d.printFloor(g, v)
	fmt.Fprintln(v, "\nGold:", laura.gold)
	fmt.Fprintln(v, "Health:", laura.health)
	for i := 0; i < MONSTERNUM; i++ {
		if isNear(laura, monster[i]) {
			fmt.Fprintln(v, "Would you like to attack? (space)")
			break
		}
	}
	return nil
}

func isNear(a Actor, b Actor) bool {
	if (a.y <= b.y) && (a.y + NEAR >= b.y) {
		if (a.x <= b.x) && (a.x + NEAR >= b.x) || (a.x >= b.x) && (a.x - NEAR <= b.x){
			return true
		} 
	} else if (a.y >= b.y) && (a.y - NEAR <= b.y) {
		if (a.x <= b.x) && (a.x + NEAR >= b.x) || (a.x >= b.x) && (a.x - NEAR <= b.x) {
			return true
		}
	} else if (a.x <= b.x) && (a.x + NEAR >= b.x) {
		if (a.y <= b.y) && (a.y + NEAR >= b.y) || (a.y >= b.y) && (a.y - NEAR <= b.y) {
			return true
		}
	} else if (a.x >= b.x) && (a.x - NEAR <= b.x) {
		if (a.y <= b.y) && (a.y + NEAR >= b.y) || (a.y >= b.y) && (a.y - NEAR <= b.y) {
			return true
		}
	}
	return false
}

func MonsterAI(){
	for i := 0; i < MONSTERNUM; i++ {
		if !monster[i].alive {
			continue
		}
		if isNear(laura, monster[i]) {
			if laura.x > monster[i].x {
				if d.floor[monster[i].x+1][monster[i].y] == "." {
					d.floor[monster[i].x][monster[i].y] = "."
					monster[i].x++
					d.floor[monster[i].x][monster[i].y] = "M"
				}
			} else {
				if d.floor[monster[i].x-1][monster[i].y] == "." {
					d.floor[monster[i].x][monster[i].y] = "."
					monster[i].x--
					d.floor[monster[i].x][monster[i].y] = "M"
				}
			}
			
		} 	
		
		if isNear(laura, monster[i]) {
			if laura.y > monster[i].y {
				if d.floor[monster[i].x][monster[i].y+1] == "." {
					d.floor[monster[i].x][monster[i].y] = "."
					monster[i].y++
					d.floor[monster[i].x][monster[i].y] = "M"
				}
			} else {
				if d.floor[monster[i].x][monster[i].y-1] == "." {
					d.floor[monster[i].x][monster[i].y] = "."
					monster[i].y--
					d.floor[monster[i].x][monster[i].y] = "M"
				}
			}
			continue
		}

		randPath := (rand.Int() % 4)
		if randPath == 0 {
			if d.floor[monster[i].x+1][monster[i].y] == "." {
				d.floor[monster[i].x][monster[i].y] = "."
				monster[i].x++
				d.floor[monster[i].x][monster[i].y] = "M"
			}
			continue
		} 
		if randPath == 1 {
			if d.floor[monster[i].x-1][monster[i].y] == "." {
				d.floor[monster[i].x][monster[i].y] = "."
				monster[i].x--
				d.floor[monster[i].x][monster[i].y] = "M"
			}
			continue
		}
		if randPath == 2 {
			if d.floor[monster[i].x][monster[i].y+1] == "." {
				d.floor[monster[i].x][monster[i].y] = "."
				monster[i].y++
				d.floor[monster[i].x][monster[i].y] = "M"
			}
			continue
		} 
		if d.floor[monster[i].x][monster[i].y-1] == "." {
			d.floor[monster[i].x][monster[i].y] = "."
			monster[i].y--
			d.floor[monster[i].x][monster[i].y] = "M"
		}
		
	}
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
	if err := g.SetKeybinding("", gocui.KeySpace, gocui.ModNone, attack); err != nil {
		return err
	}
	return nil
}

func attack(g *gocui.Gui, v *gocui.View) error {
	for i := 0;i < MONSTERNUM; i++ {
		if isNear(laura, monster[i]) && monster[i].alive {
			monster[i].health -= laura.attack
			if monster[i].health <= 0 {
				monster[i].alive = false
				d.floor[monster[i].x][monster[i].y] = "."
			}
			laura.health -= monster[i].attack
			
			if laura.health <= 0 {
				laura.alive = false
			}
			break
		}
	}

	return nil
}

func quit(g * gocui.Gui, v * gocui.View) error {
	return gocui.ErrQuit
}

func (l *Actor)ifAround(i string) bool {
	if d.floor[l.x][l.y + 1] == i || d.floor[l.x][l.y - 1] == i || d.floor[l.x + 1][l.y] == i || d.floor[l.x - 1][l.y] == i{
		return true
	}
	return false

}

func (l *Actor)moveW(g *gocui.Gui, v *gocui.View) error{
	if laura.ifAround("G") {
		l.gold++
		laura.health += 5
		d.genFloor()
		return nil
	}

	if d.floor[laura.x - 1][laura.y] == "." {
		d.floor[laura.x][laura.y] = "."
		l.x -= 1
	}
	d.floor[laura.x][laura.y] = "@"
	return nil
}
func (l *Actor)moveA(g *gocui.Gui, v *gocui.View) error{
	if laura.ifAround("G") {
		l.gold++
		laura.health += 5
		d.genFloor()
		return nil
	}

	if d.floor[laura.x][laura.y - 1] == "." {
		d.floor[laura.x][laura.y] = "."
		l.y -= 1
	}
	d.floor[laura.x][laura.y] = "@"
	return nil
}
func (l *Actor)moveS(g *gocui.Gui, v *gocui.View) error{
	if laura.ifAround("G") {
		l.gold++
		laura.health += 5
		d.genFloor()
		return nil
	}

	if d.floor[laura.x + 1][laura.y] == "." {
		d.floor[laura.x][laura.y] = "."
		l.x += 1
	}
	d.floor[laura.x][laura.y] = "@"
	return nil
}
func (l *Actor)moveD(g *gocui.Gui, v *gocui.View) error{
	if laura.ifAround("G") {
		l.gold++
		laura.health += 5
		d.genFloor()
		return nil
	}

	if d.floor[laura.x][laura.y + 1] == "." {
		d.floor[laura.x][laura.y] = "."
		l.y += 1
	}
	d.floor[laura.x][laura.y] = "@"
	return nil
}


