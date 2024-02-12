package entities

import (
	"fmt"
	"strings"
	"time"
)

const timeLayout = "01/06"

type Card struct {
	Meta      string
	Number    string
	ValidThru CardTime
	Holder    string
	CVV       int
}

type CardTime struct {
	time.Time
}

func (f *CardTime) Parse(s string) error {
	s = strings.TrimSuffix(s, "\n")
	t, err := time.Parse(timeLayout, s)
	if err != nil {
		return fmt.Errorf("cant parse time %T to format %s", t, timeLayout)
	}
	f.Time = t
	return nil
}

func (f CardTime) Value() []byte {
	return []byte(fmt.Sprintf("\"%s\"", f.Time.Format(timeLayout)))
}

func (f CardTime) String() string {
	return f.Time.Format(timeLayout)
}
