package mouse

type Button int

const (
	ButtonNone Button = iota
	ButtonLeft
	ButtonMiddle
	ButtonRight
)

type Action int

const (
	ActionNone Action = iota
	ActionPress
	ActionRelease
	ActionMove
)
