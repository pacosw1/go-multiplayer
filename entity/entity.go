package entity

import (
	"sockets/message"
	"sockets/types"
)

// Player Stores state data for a player
type Player struct {
	Health int
	*types.Position
	Aim        *types.Position
	IsShooting bool
	SequenceID int16
	ID         int
}

//NewPlayer create a new player
func NewPlayer(clientID int) *Player {
	return &Player{
		Health: 100,
		Position: &types.Position{
			X: 0,
			Y: 0,
		},
		Aim: &types.Position{
			X: 0,
			Y: 0,
		},
		IsShooting: false,
		SequenceID: 0,
		ID:         clientID,
	}
}

//UpdatePlayer <- updates player based on input
func (p *Player) UpdatePlayer(r *message.UserInput) {
	p.Position = r.Position
	p.SequenceID = r.SequenceID
	p.IsShooting = r.IsShooting
	p.Aim = r.Aim
}

//Projectile stores bullet postion and angle
type Projectile struct {
	Aim      *types.Position
	Position *types.Position
	ID       int
}