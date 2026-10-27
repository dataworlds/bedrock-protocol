package types

type InputMode uint32

const (
	InputModeUndefined InputMode = 0
	InputModeMouse     InputMode = 1
	InputModeTouch     InputMode = 2
	InputModeGamePad   InputMode = 3
	InputModeCount     InputMode = 5
)
