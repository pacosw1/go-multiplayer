package message

import (
	"sockets/types"
	"time"
)

//Connect request to create new Player on user joined
type Connect struct {
	Username string
	ClientID uint32
}

//ConnectMessage creates a new message of type Connect
func ConnectMessage(username string, ID uint32) *Connect {
	return &Connect{
		Username: username,
		ClientID: ID,
	}
}

//Disconnect request to disconnect player from server
type Disconnect struct {
	ClientID uint32
}

//DisconnectMessage creates a new message of type Disconnect
func DisconnectMessage(ID uint32) *Disconnect {
	return &Disconnect{
		ClientID: ID,
	}
}

//StateMessage broadcast message send to clients
type StateMessage struct {
	timestamp time.Time
}

//SendState message constructor
func SendState() *StateMessage {
	return &StateMessage{
		timestamp: time.Now(),
	}
}

//SendInput wtf
func SendInput(m NetworkInput) *NetworkInput {
	return &NetworkInput{
		IsShooting: m.IsShooting,
		Direction:  m.Direction,
		SequenceID: m.SequenceID,
		ID:         m.ID,
	}
}

//UserInput request to update player state
type UserInput struct {
	IsShooting bool
	*types.Vector
	SequenceID int16
	Aim        *types.Vector
	ID         int
}

//SpawnProjectile message to spawn a projectile
type SpawnProjectile struct {
	ID        int
	Position  *types.Vector
	Direction *types.Vector
	PlayerID  int
}

//KillPlayer event trigger on delete
type KillPlayer struct {
	ID int
}

//ProjectileHit message to destroy projectile
type ProjectileHit struct {
	ID int16
}
