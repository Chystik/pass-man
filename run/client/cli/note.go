package cli

import (
	"bufio"
	"context"
	"fmt"
	"os"

	"github.com/Chystik/pass-man/internal/vault/note/entities"
)

func (c *cli) note(ctx context.Context, r *bufio.Reader) {
	cleanScr()
	label := `
		Note operations:
		
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
			n, err := c.api.ListNote(ctx)
			if err != nil {
				fmt.Fprintln(os.Stderr, err.Error())
			} else {
				printNoteList(n)
			}

			fmt.Fprintf(os.Stdout, "Press Enter to continue")
			r.ReadString('\n')
		case '2':
			reset(r)
			var n entities.Note

			fmt.Fprintf(os.Stdout, "Description: ")
			n.Meta, _ = r.ReadString('\n')
			fmt.Fprintf(os.Stdout, "Note: ")
			n.Note, _ = r.ReadString('\n')

			err := c.api.AddNote(ctx, n)
			if err != nil {
				fmt.Fprintln(os.Stdout, err)
				break
			}

			fmt.Fprintln(os.Stdout, "The note was added")
		case '3':
			reset(r)
			fmt.Fprintf(os.Stdout, "Description: ")
			meta, _ := r.ReadString('\n')

			n, err := c.api.GetNote(ctx, meta)
			if err != nil {
				fmt.Fprintln(os.Stderr, err.Error())
			} else {
				fmt.Fprintf(os.Stdout, "Username: %sPassword: %s\n", n.Meta, n.Note)
			}

			fmt.Fprintf(os.Stdout, "Press Enter to continue")
			r.ReadString('\n')
		case '4':
			reset(r)
			fmt.Fprintf(os.Stdout, "Description: ")
			meta, _ := r.ReadString('\n')

			err := c.api.DeleteNote(ctx, meta)
			if err != nil {
				fmt.Fprintln(os.Stderr, err.Error())
			} else {
				fmt.Fprintln(os.Stdout, "The note was removed")
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

func printNoteList(p []entities.Note) {
	for i := range p {
		fmt.Fprintf(os.Stdout,
			"(%d): Description: %sUsername: %s\n",
			i+1, p[i].Meta, p[i].Note)
	}
}
