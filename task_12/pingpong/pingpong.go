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
	player1 player
	player2 player
	server  *player
	ch      chan string
}

func newGame(p1 player, p2 player) *game {
	s := game{
		player1: p1,
		player2: p2,
		ch:      make(chan string),
	}
	return &s
}

func (g *game) chooseServer() {
	if g.randChoose(50) {
		g.server = &g.player1
	} else {
		g.server = &g.player1
	}
}

func (g *game) Serve(hit string) {
	fmt.Printf("Подает игрок -> %v: %s\n", g.server.name, hit)

	if g.randChoose(20) {
		g.server.score++
		fmt.Println("\tУспешная подача!!!", g.server)
	}

	time.Sleep(time.Second)

	if g.server.score > 10 {
		hit = "Stop"
	}

	g.ch <- hit
	g.changeActivity()
}

func (g *game) changeActivity() {
	if *g.server == g.player1 {
		g.server = &g.player2
	} else {
		g.server = &g.player1
	}
}

func (g *game) randChoose(p int) bool {
	n := rand.Intn(100)
	if n < p {
		return true
	}

	return false
}

func (g *game) finalScore() {
	fmt.Println(g)
}

func main() {
	rand.Seed(time.Now().UnixNano())
	wg := sync.WaitGroup{}
	wg.Add(2)

	p1 := player{name: "Player_1"}
	p2 := player{name: "Player_2"}
	game := newGame(p1, p2)

	go play(p1, game, &wg)
	go play(p2, game, &wg)

	game.chooseServer()

	game.ch <- "begin"
	wg.Wait()

	game.finalScore()
}

func play(pl player, game *game, wg *sync.WaitGroup) {
	defer wg.Done()

	for val := range game.ch {
		switch val {
		case "begin":
			game.Serve("Ping")
		case "Ping":
			game.Serve("Pong")
		case "Pong":
			game.Serve("Ping")
		case "Stop":
			close(game.ch)
		}
	}
}
