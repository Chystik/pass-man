package cli

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/Chystik/pass-man/internal/vault/file/entities"
)

func (c *cli) file(ctx context.Context, r *bufio.Reader) {
	cleanScr()

	label := `
		Files operations:
		
		(1) List
		(2) Upload
		(3) Download
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
			f, err := c.api.ListFiles(ctx)
			if err != nil {
				fmt.Fprintln(os.Stderr, err.Error())
			} else {
				printFileList(f)
			}

			fmt.Fprintf(os.Stdout, "Press Enter to continue")
			r.ReadString('\n')
		case '2':
			reset(r)
			f := &entities.File{}

			fmt.Fprintf(os.Stdout, "Description: ")
			f.Meta, _ = r.ReadString('\n')
			fmt.Fprintf(os.Stdout, "FilePath: ")
			filePath, _ := r.ReadString('\n')

			filePath = strings.TrimSuffix(filePath, "\n")
			filePath = strings.TrimSuffix(filePath, "\r")

			err := c.api.UploadFile(ctx, f, filePath)
			if err != nil {
				fmt.Fprintln(os.Stdout, err)
				break
			}

			fmt.Fprintln(os.Stdout, "The file was uploaded")
		case '3':
			reset(r)
			f := &entities.File{}

			fmt.Fprintf(os.Stdout, "FileID: ")
			idStr, _ := r.ReadString('\n')

			idStr = strings.TrimSuffix(idStr, "\n")
			idStr = strings.TrimSuffix(idStr, "\r")

			id, err := strconv.Atoi(idStr)
			if err != nil {
				fmt.Fprintln(os.Stdout, err)
				break
			}

			f.ID = uint32(id)

			fmt.Fprintf(os.Stdout, "Save to file: ")
			fPath, _ := r.ReadString('\n')

			fPath = strings.TrimSuffix(fPath, "\n")
			fPath = strings.TrimSuffix(fPath, "\r")

			err = c.api.DownloadFile(ctx, f, fPath)
			if err != nil {
				fmt.Fprintln(os.Stderr, err.Error())
			} else {
				fmt.Fprintln(os.Stdout, "The file was downloaded")
			}

			fmt.Fprintf(os.Stdout, "Press Enter to continue")
			r.ReadString('\n')
		case '4':
			reset(r)
			f := &entities.File{}

			fmt.Fprintf(os.Stdout, "FileID: ")
			idStr, _ := r.ReadString('\n')

			idStr = strings.TrimSuffix(idStr, "\n")
			idStr = strings.TrimSuffix(idStr, "\r")

			id, err := strconv.Atoi(idStr)
			if err != nil {
				fmt.Fprintln(os.Stdout, err)
			}

			f.ID = uint32(id)

			err = c.api.DeleteFile(ctx, f)
			if err != nil {
				fmt.Fprintln(os.Stderr, err.Error())
			} else {
				fmt.Fprintln(os.Stdout, "The file was removed")
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

func printFileList(f []*entities.File) {
	for i := range f {
		fmt.Fprintf(os.Stdout,
			"(%d): FileID:%d Description: %sFilename: %s\n",
			i+1, f[i].ID, f[i].Meta, f[i].Name)
	}
}
