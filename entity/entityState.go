package entities

type entityState uint8

const (
	down entityState = iota
	up
	left
	right
	idle
)
