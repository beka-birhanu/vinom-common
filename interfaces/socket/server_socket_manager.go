// Package i defines interfaces for cryptographic operations, socket communication,
// and message encoding used in a secure server-client architecture.
//
// It includes interfaces for symmetric and asymmetric cryptography, HMAC authentication,
// server socket management, handshake processes, and encoding/decoding mechanisms.
package i

import "github.com/google/uuid"

// ServerSocketManager defines an interface for managing server-side socket communication
// and handling client interactions.
type ServerSocketManager interface {
	// SetClientRequestHandler sets a handler function for processing client requests.
	SetClientRequestHandler(func(clientID uuid.UUID, requestType byte, requestData []byte))

	// SetClientRegisterHandler sets a handler function for registering new clients.
	SetClientRegisterHandler(func(clientID uuid.UUID))

	// Stop gracefully shuts down the server.
	Stop()

	// Serve starts the server and begins listening for client connections.
	Serve()

	// SetClientAuthenticator sets an authenticator to verify client identities.
	SetClientAuthenticator(Authenticator)

	// BroadcastToClients sends a message to multiple clients.
	BroadcastToClients(clientIDs []uuid.UUID, messageType byte, message []byte)

	// GetPublicKey returns the server's public key for secure communication.
	GetPublicKey() []byte

	// GetAddr returns the server's socket address.
	GetAddr() string
}
