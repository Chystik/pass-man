package cli

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	vault "github.com/Chystik/pass-man/internal/vault/entities"
)

func (c *cli) card(ctx context.Context, r *bufio.Reader) {
	cleanScr()
	label := `
		Card operations:
		
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
			c, err := c.api.ListCard(ctx)
			if err != nil {
				fmt.Fprintln(os.Stderr, err.Error())
			} else {
				printCardList(c)
			}

			fmt.Fprintf(os.Stdout, "Press Enter to continue")
			r.ReadString('\n')
		case '2':
			reset(r)
			var card vault.Card

			fmt.Fprintf(os.Stdout, "Description: ")
			card.Meta, _ = r.ReadString('\n')
			fmt.Fprintf(os.Stdout, "Number: ")
			card.Number, _ = r.ReadString('\n')

			card.Number = strings.TrimSuffix(card.Number, "\n")
			card.Number = strings.TrimSuffix(card.Number, "\r")
			_, err := strconv.Atoi(card.Number)
			if err != nil {
				if err != nil {
					fmt.Fprintln(os.Stderr, errors.New("only digits allowed in number"))
					break
				}
			}

			if len(card.Number) != 16 {
				fmt.Fprintln(os.Stderr, len(card.Number), card.Number)
				fmt.Fprintln(os.Stderr, errors.New("card number != 16 digits"))
				break
			}

			fmt.Fprintf(os.Stdout, "ValidThru \"MM/YY\": ")
			vthStr, _ := r.ReadString('\n')
			vthStr = strings.TrimSuffix(vthStr, "\n")
			vthStr = strings.TrimSuffix(vthStr, "\r")

			err = card.ValidThru.Parse(vthStr)
			if err != nil {
				fmt.Fprintln(os.Stderr, err, errors.New("date format: \"MM/YY\""))
				break
			}

			fmt.Fprintf(os.Stdout, "Holder Name: ")
			card.Holder, _ = r.ReadString('\n')

			fmt.Fprintf(os.Stdout, "CVV: ")
			cvvStr, _ := r.ReadString('\n')
			cvvStr = strings.TrimSuffix(cvvStr, "\n")
			cvvStr = strings.TrimSuffix(cvvStr, "\r")
			card.CVV, err = strconv.Atoi(cvvStr)
			if err != nil {
				fmt.Fprintln(os.Stderr, errors.New("only digits allowed in CVV"))
				break
			}

			err = c.api.AddCard(ctx, card)
			if err != nil {
				fmt.Fprintln(os.Stdout, err)
			}

			fmt.Fprintln(os.Stdout, "The card was added")
		case '3':
			reset(r)
			fmt.Fprintf(os.Stdout, "Description: ")
			meta, _ := r.ReadString('\n')

			card, err := c.api.GetCard(ctx, meta)
			if err != nil {
				fmt.Fprintln(os.Stderr, err.Error())
			} else {
				fmt.Fprintf(os.Stdout,
					"Number: %s\nValidThru: %s\nHolder Name:%s\nCVV: %d\n",
					card.Number, card.ValidThru, card.Holder, card.CVV)
			}

			fmt.Fprintf(os.Stdout, "Press Enter to continue")
			r.ReadString('\n')
		case '4':
			reset(r)
			fmt.Fprintf(os.Stdout, "Description: ")
			meta, _ := r.ReadString('\n')

			err := c.api.DeleteCard(ctx, meta)
			if err != nil {
				fmt.Fprintln(os.Stderr, err.Error())
			} else {
				fmt.Fprintln(os.Stdout, "The card was removed")
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

func printCardList(c []vault.Card) {
	for i := range c {
		fmt.Fprintf(os.Stdout,
			"(%d): Description: %s\nNumber: %s\nValidThru: %s\nHolder: %s\nCVV: %d\n",
			i+1, c[i].Meta, c[i].Number, c[i].ValidThru, c[i].Holder, c[i].CVV)
	}
}
