package poker

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
	"time"
)

const PlayerPrompt = "Please enter the number of players: "

type Game struct {
	alerter BlindAlerter
	store   PlayerStore
}

func NewGame(alerter BlindAlerter, store PlayerStore) *Game {
	return &Game{
		alerter: alerter,
		store:   store,
	}
}

func (g *Game) Start(numberOfPlayers int) {
	blindIncrement := time.Duration(5+numberOfPlayers) * time.Minute

	blinds := []int{100, 200, 300, 400, 500, 600, 800, 1000, 2000, 4000, 8000}
	blindTime := 0 * time.Second
	for _, blind := range blinds {
		g.alerter.ScheduleAlertAt(blindTime, blind)
		blindTime = blindTime + blindIncrement
	}
}

func (g *Game) Finish(winner string) {
	g.store.RecordWin(winner)
}

type CLI struct {
	in   *bufio.Scanner
	out  io.Writer
	game *Game
}

func NewCLI(in io.Reader, out io.Writer, game *Game) *CLI {
	return &CLI{
		in:   bufio.NewScanner(in),
		out:  out,
		game: game,
	}
}

func (cli *CLI) PlayPoker() {
	fmt.Fprint(cli.out, PlayerPrompt)
	numberOfPlayers, _ := strconv.Atoi(strings.Trim(cli.readLine(), "\n"))
	cli.game.Start(numberOfPlayers)
	winnerInput := cli.readLine()
	winner := extractWinner(winnerInput)
	cli.game.Finish(winner)
}

func (cli *CLI) readLine() string {
	cli.in.Scan()
	return cli.in.Text()
}

func extractWinner(userInput string) string {
	return strings.Replace(userInput, " wins", "", 1)
}
