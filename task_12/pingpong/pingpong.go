package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type player struct {
	name  string
	score int
}

type game struct {
	player1 *player
	player2 *player
	ch      chan string
}

func newGame(p1 *player, p2 *player) *game {
	s := game{
		player1: p1,
		player2: p2,
		ch:      make(chan string),
	}
	return &s
}

func (g *game) Serve(pl *player, hit string) {
	fmt.Printf("Игрок %s -> %s\n", pl.name, hit)

	if g.randChoose(20) {
		fmt.Println("\tУспешная подача!!!")
		pl.score++
	}

	if pl.score > 10 {
		hit = "Stop"
	}

	// time.Sleep(time.Millisecond * 200)

	g.ch <- hit
}

func (g *game) randChoose(p int) bool {
	n := rand.Intn(100)
	if n < p {
		return true
	}

	return false
}

func main() {
	rand.Seed(time.Now().UnixNano())
	wg := sync.WaitGroup{}
	wg.Add(2)

	p1 := player{name: "Player_1"}
	p2 := player{name: "Player_2"}
	game := newGame(&p1, &p2)

	go play(&p1, game, &wg)
	go play(&p2, game, &wg)

	game.ch <- "begin"
	wg.Wait()

	game.finalScore()
}

func play(pl *player, game *game, wg *sync.WaitGroup) {
	defer wg.Done()

	for val := range game.ch {
		switch val {
		case "begin":
			game.Serve(pl, "Ping")
		case "Ping":
			game.Serve(pl, "Pong")
		case "Pong":
			game.Serve(pl, "Ping")
		case "Stop":
			close(game.ch)
		}
	}
}

func (g *game) finalScore() {
	fmt.Println("\tScore")
	fmt.Printf("%s(%v) : %s(%v)\n", g.player1.name, g.player1.score, g.player2.name, g.player2.score)
}
