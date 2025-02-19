package i

import "github.com/google/uuid"

// Authenticator an interface for authenticating the client token
type Authenticator interface {
	Authenticate([]byte) (clientID uuid.UUID, error error)
}
