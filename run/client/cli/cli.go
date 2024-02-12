package cli

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"time"

	"github.com/Chystik/pass-man/internal/vault"
)

type cli struct {
	api vault.VaultAPICliet
}

func NewCli(api vault.VaultAPICliet) *cli {
	return &cli{
		api: api,
	}
}

func (c *cli) Main(ctx context.Context) {
	time.Sleep(time.Second)
	cleanScr()
	r := bufio.NewReader(os.Stdin)

	label := `
		Select category:
		
		(1) Passwords
		(2) Bank cards
		(3) Notes
		(4) Files

		(q) Quit
	`

	for {
		fmt.Printf("%s\n>>>>", label)
		r.Reset(os.Stdin)
		s, _, _ := r.ReadRune()
		switch s {
		case '1':
			cleanScr()
			c.password(ctx, r)
		case '2':
			cleanScr()
			c.card(ctx, r)
		case '3':
			cleanScr()
			c.note(ctx, r)
		case '4':
			cleanScr()
			c.file(ctx, r)
		case 'q':
			cleanScr()
			return
		default:
			wrongSelect()
		}

	}
}

func cleanScr() {
	fmt.Fprint(os.Stdout, "\033c")
}

func reset(r *bufio.Reader) {
	cleanScr()
	r.Reset(os.Stdin)
}

func wrongSelect() {
	cleanScr()
	fmt.Fprintln(os.Stderr, "Please select a number")
}
