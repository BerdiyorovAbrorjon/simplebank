package token

import "time"

// Maker is an interface for managing tokens
type Maker interface {
	//CreateToken is function for create token string
	CreateToken(username string, duration time.Duration) (string, *Payload, error)

	//VerifyToken is function for check token to valid or not
	VerifyToken(token string) (*Payload, error)
}
