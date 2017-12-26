package main

import (
	"fmt"
	//"os"
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

	fmt.Println("start.")
	for maxturns := 1; maxturns < 11; maxturns++ {
		fmt.Println("Maxturns depth", maxturns)
		beast := NewMonster()
		//beast.showMonster()
		*beast = slayMonster(*beast, 0, maxturns)
		if beast.Death() {
			// we're done :)
			fmt.Println("Yes! We slayed the monster!")
			beast.PrintMoves(maxturns)
			break // break out of the loop
		}
	}
}

// TryMove is the type definining the four possible moves of the player
type TryMove int32

const (
	ONE_HEAD  TryMove = 0
	TWO_HEADS TryMove = 1
	ONE_TAIL  TryMove = 2
	TWO_TAILS TryMove = 3
)

var Move_name = map[TryMove]string{
	ONE_HEAD:  "ONE_HEAD",
	TWO_HEADS: "TWO_HEADS",
	ONE_TAIL:  "ONE_TAIL",
	TWO_TAILS: "TWO_TAILS",
}

func slayMonster(m Monster, turns int, maxturns int) Monster {
	//fmt.Print("Turn:", turns)
	// m.showMonster()
	if turns > maxturns || m.MonsterWins() || m.Death() {
		// we either did not succeed in the amount of time allotted,
		// or we lost, or we won. Either way we are at the end of our tries:
		//fmt.Print(".")
		return m
	}
	for idx, _ := range Move_name {
		beast := m
		beast.ExecuteMove(idx, turns)
		result := slayMonster(beast, turns+1, maxturns)
		if result.Death() {
			return result
		}
	}

	return m
}

// ExecuteMove slashes into the monster with move TryMove and
// ensures the moves is tracked in the trackmoves list at place 'turn'
func (m *Monster) ExecuteMove(move TryMove, turn int) {
	// fmt.Println(Move_name[move])
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
type Monster struct {
	heads      int
	tails      int
	trackmoves []TryMove // slice of tried moves
}

// NewMonster gives you a new monster,
// initialized with 3 heads and 3 tails
func NewMonster() *Monster {
	m := new(Monster)
	m.heads = 3
	m.tails = 3
	m.trackmoves = make([]TryMove, 12)
	return m
}
func (m *Monster) showMonster() {
	fmt.Printf("%+v", m)
}

// Death tells if the monster has been slain
func (m *Monster) Death() bool {
	if m.heads == 0 && m.tails == 0 {
		return true
	}
	return false
}

// MonsterWins tells if the monster has become invincible
func (m *Monster) MonsterWins() bool {
	if m.heads == 1 && m.tails == 0 {
		return true
	}
	return false
}

// PrintMoves is to be called when monster is slain to show what moves let to this
// the program stores earlier moves in the trackmoves list
func (m *Monster) PrintMoves(turns int) {
	printmonster := NewMonster()
	for i := 0; i <= turns; i++ {
		fmt.Printf("%+v", *printmonster)
		fmt.Println(Move_name[TryMove(m.trackmoves[i])])
		printmonster.ExecuteMove(m.trackmoves[i], i)

	}
}
