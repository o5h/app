package keyboard

import "fmt"

type Direction byte

const (
	Repeat  Direction = 0
	Press   Direction = 1
	Release Direction = 2
)

//Event is a keyboard event
type Event struct {
	Direction Direction
	Code      Code
	Rune      rune
}

func (e *Event) String() string {
	return fmt.Sprintf("direct:=%v:code=%v:rune:=%v", e.Direction, e.Code, string(e.Rune))
}
