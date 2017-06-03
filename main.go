package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Let's slash a monster. Here are the rules:")
	fmt.Println("- The monster has 3 heads and 3 tails.")
	fmt.Println("- You have a SWORD and you can choose between these actions:")
	fmt.Println("1. Slash off ONE HEAD. Result: one head will grow back.")
	fmt.Println("2. Slash off TWO HEADS. Result: monster will loose two heads!")
	fmt.Println("3. Slash off ONE TAIL. Result: TWO tails will grow back.")
	fmt.Println("4. Slash off TWO TAILS. Result: one head will grow back.")
	fmt.Println("Can you slay the monster, and in how many turns?")

	beast := NewMonster()
	beast.showMonster()

	*beast = slayMonster(*beast, ONE_HEAD, 0)
	if (beast.Death()) {
		// we're done :)

	}
}

type TryMove int32

const (
	ONE_HEAD  TryMove = 0
	TWO_HEADS TryMove = 1
	ONE_TAIL  TryMove = 2
	TWO_TAILS TryMove = 3
)

var Move_name = map[int]string{
	0: "ONE_HEAD",
	1: "TWO_HEADS",
	2: "ONE_TAIL",
	3: "TWO_TAILS",
}

func slayMonster(m monster, move TryMove, turns int) monster {
	fmt.Print("Turn:", turns)
	if turns >= 12 || m.MonsterWins() || m.Death() {
		// we either did not succeed in the amount of time allotted,
		// or we lost, or we won. Either way we are at the end of our tries:
		fmt.Print(".")
		return m
	}

	m.ExecuteMove(move, turns)

	m.showMonster()

	if m.Death() {
		fmt.Println("Yes! We slayed the monster!")

		m.PrintMoves()
		os.Exit(0)
		return m
	}
	slayMonster(m, TWO_HEADS, turns+1)
	if m.Death() {
		return m
	}
	slayMonster(m, ONE_TAIL, turns+1)
	if m.Death() {
		return m
	}
	slayMonster(m, TWO_TAILS, turns+1)

	if m.Death() {
		return m
	}
	return m
}

func (m *monster) ExecuteMove(move TryMove, turn int) {
	fmt.Println(Move_name[int(move)])
	switch move {
	case ONE_HEAD:
		if m.heads > 0 {

		}
	case TWO_HEADS:
		if m.heads > 1 {
			m.heads -= 2
			m.trackmoves[turn] = TWO_HEADS
		}
	case ONE_TAIL:
		if m.tails > 0 {
			m.tails++
			m.trackmoves[turn] = ONE_TAIL
		}
	case TWO_TAILS:
		if m.tails > 1 {
			m.tails -= 2
			m.heads++
			m.trackmoves[turn] = TWO_TAILS
		}
	}
}

// Monster holds the number of heads and tails
type monster struct {
	heads      int
	tails      int
	trackmoves []TryMove // slice of tried moves
}

// NewMonster gives you a new monster,
// initialized with 3 heads and 3 tails
func NewMonster() *monster {
	m := new(monster)
	m.heads = 3
	m.tails = 3
	m.trackmoves = make([]TryMove, 12)
	return m
}
func (m *monster) showMonster() {
	fmt.Printf("%+v", m)
}

func (m *monster) Death() bool {
	if m.heads == 0 && m.tails == 0 {
		return true
	}
	return false
}

func (m *monster) MonsterWins() bool {
	if m.heads == 1 && m.tails == 0 {
		return true
	}
	return false
}

func (m *monster) PrintMoves() {
	printmonster := NewMonster()
	for i := 0; i< len(m.trackmoves);i++ {
		printmonster.ExecuteMove(m.trackmoves[i],i)

		fmt.Printf("%+v", printmonster)
	}
}
