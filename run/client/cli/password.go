package cli

import (
	"bufio"
	"context"
	"fmt"
	"os"

	"github.com/Chystik/pass-man/internal/vault/password/entities"
)

func (c *cli) password(ctx context.Context, r *bufio.Reader) {
	cleanScr()
	label := `
		Password operations:
		
		(1) List
		(2) Add
		(3) Get
		(4) Delete

		(b) Back
		(q) Quit
	`
	for {
		fmt.Printf("%s\n>>>>", label)
		r.Reset(os.Stdin)
		s, _, _ := r.ReadRune()
		switch s {
		case '1':
			reset(r)
			p, err := c.api.ListPassword(ctx)
			if err != nil {
				fmt.Fprintln(os.Stderr, err.Error())
			} else {
				printPasswordList(p)
			}

			fmt.Fprintf(os.Stdout, "Press Enter to continue")
			r.ReadString('\n')
		case '2':
			reset(r)
			var p entities.Password

			fmt.Fprintf(os.Stdout, "Description: ")
			p.Meta, _ = r.ReadString('\n')
			fmt.Fprintf(os.Stdout, "Username: ")
			p.Username, _ = r.ReadString('\n')
			fmt.Fprintf(os.Stdout, "Password: ")
			p.Password, _ = r.ReadString('\n')

			err := c.api.AddPassword(ctx, p)
			if err != nil {
				fmt.Fprintln(os.Stdout, err)
				break
			}

			fmt.Fprintln(os.Stdout, "The password was added")
		case '3':
			reset(r)
			fmt.Fprintf(os.Stdout, "Description: ")
			meta, _ := r.ReadString('\n')

			p, err := c.api.GetPassword(ctx, meta)
			if err != nil {
				fmt.Fprintln(os.Stderr, err.Error())
			} else {
				fmt.Fprintf(os.Stdout, "Username: %sPassword: %s\n", p.Username, p.Password)
			}

			fmt.Fprintf(os.Stdout, "Press Enter to continue")
			r.ReadString('\n')
		case '4':
			reset(r)
			fmt.Fprintf(os.Stdout, "Description: ")
			meta, _ := r.ReadString('\n')

			err := c.api.DeletePassword(ctx, meta)
			if err != nil {
				fmt.Fprintln(os.Stderr, err.Error())
			} else {
				fmt.Fprintln(os.Stdout, "The password was removed")
			}

			fmt.Fprintf(os.Stdout, "Press Enter to continue")
			r.ReadString('\n')

		case 'b':
			cleanScr()
			return
		case 'q':
			cleanScr()
			os.Exit(0)
		default:
			wrongSelect()
		}
	}
}

func printPasswordList(p []entities.Password) {
	for i := range p {
		fmt.Fprintf(os.Stdout,
			"(%d): Description: %sUsername: %sPassword: %s\n",
			i+1, p[i].Meta, p[i].Username, p[i].Password)
	}
}
