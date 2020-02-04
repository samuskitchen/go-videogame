package internal

type Direction int

const (
	Idle Direction = iota
	LeftDirection
	RightDirection
	FrontDirection
	BackDirection
)