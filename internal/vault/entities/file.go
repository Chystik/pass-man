package entities

import "io"

type File struct {
	ID   uint32
	Meta string
	Name string
	Data io.ReadWriter
}
