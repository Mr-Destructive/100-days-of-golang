package main

import (
	"fmt"
	"strconv"
)

type Creature interface {
	intro() string
	attack(*int) int
	//heal() int
}

type Player struct {
	name   string
	health int
}

type Mob struct {
	name     string
	health   int
	category bool
}

func (p Player) intro() string {
	fmt.Println("Player has spawned")
	return p.name
}

func (p Player) attack(m_health *int) int {
	fmt.Println("Player has attacked!")
	*m_health = *m_health - 50
	return *m_health
}

func (m Mob) intro() string {
	var name string
	if m.name != "" {
		name = m.name
	} else {
		name = "Mob"
	}
	fmt.Printf("A wild %s has appeared!\n", name)
	return m.name
}

func (m Mob) attack(p_health *int) int {
	fmt.Printf("%s has attacked you! -%d\n", m.name, 30)
	*p_health = *p_health - 30
	return *p_health
}

func parse_int(n interface{}) int {
	switch n.(type) {
	case int:
		return (n).(int) * (n).(int)
	case string:
		s, _ := strconv.Atoi(n.(string))
		return s
	case float64:
		return int(n.(float64))
	default:
		return n.(int)
	}
}

func main() {
	player := Player{name: "Steve", health: 100}
	mob := Mob{name: "Zombie", health: 140}
	fmt.Println(player.intro())
	fmt.Println(mob.intro())
	fmt.Println(mob)
	fmt.Println(player)
	fmt.Println(player.attack(&mob.health))
	fmt.Println(mob.attack(&player.health))
	fmt.Println(mob)
	fmt.Println(player)
	num := parse_int(4)
	fmt.Println(num)
	num = parse_int("4")
	fmt.Println(num)
	num = parse_int(4.1243)
	fmt.Println(num)

	entity := []Creature{Player{}, Mob{}, Mob{}, Player{}}

	for _, obj := range entity {
		fmt.Println(obj.intro())
	}
}
