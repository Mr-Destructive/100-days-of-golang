package main

import "fmt"

type Player struct {
	name   string
	health int
}

type Mob struct {
	name       string
	health     int
	is_passive bool
}

type Creature interface {
	intro() string
	//attack() int
	//heal() int
}

func main() {
	player := Player{name: "Steve"}
	mob := Mob{name: "Zombie"}
	fmt.Println(player)
	fmt.Println(mob)
}
