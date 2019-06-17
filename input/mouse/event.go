package mouse

import "fmt"

type Event struct {
	Action Action
	X, Y   int
	Button Button
}

func (e *Event) String() string {
	return fmt.Sprintf("action:=%v:x=%v:y=%v:btn:=%v", e.Action, e.X, e.Y, e.Button)
}
