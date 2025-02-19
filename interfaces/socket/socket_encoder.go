package i

// HandshakeRecord represents a record used during the handshake process.
type HandshakeRecord interface {
	// GetSessionID returns the session ID.
	GetSessionID() []byte
	// SetSessionID sets the session ID.
	SetSessionID(sessionID []byte)

	// GetRandom returns the random value used in the handshake.
	GetRandom() []byte
	// SetRandom sets the random value.
	SetRandom(random []byte)

	// GetCookie returns the cookie value.
	GetCookie() []byte
	// SetCookie sets the cookie value.
	SetCookie(cookie []byte)

	// GetToken returns the authentication token.
	GetToken() []byte
	// SetToken sets the authentication token.
	SetToken(token []byte)

	// GetKey returns the cryptographic key.
	GetKey() []byte
	// SetKey sets the cryptographic key.
	SetKey(key []byte)

	// GetTimestamp returns the timestamp of the handshake.
	GetTimestamp() int64
	// SetTimestamp sets the timestamp of the handshake.
	SetTimestamp(timestamp int64)
}

// PingRecord represents a record used for tracking ping messages.
type PingRecord interface {
	// GetSentAt returns the timestamp when the ping was sent.
	GetSentAt() int64
	// SetSentAt sets the timestamp when the ping was sent.
	SetSentAt(sentAt int64)
}

// PongRecord represents a record used for tracking pong messages.
type PongRecord interface {
	// GetPingSentAt returns the timestamp when the associated ping was sent.
	GetPingSentAt() int64
	// SetPingSentAt sets the timestamp when the associated ping was sent.
	SetPingSentAt(pingSentAt int64)

	// GetReceivedAt returns the timestamp when the pong was received.
	GetReceivedAt() int64
	// SetReceivedAt sets the timestamp when the pong was received.
	SetReceivedAt(receivedAt int64)

	// GetSentAt returns the timestamp when the pong was sent.
	GetSentAt() int64
	// SetSentAt sets the timestamp when the pong was sent.
	SetSentAt(sentAt int64)
}

// SocketEncoder defines an interface for encoding and decoding socket messages.
type SocketEncoder interface {
	// Marshal serializes an object into a byte slice.
	Marshal(data interface{}) ([]byte, error)

	// Unmarshal deserializes a byte slice into an object.
	Unmarshal(data []byte, obj interface{}) error

	// NewHandshakeRecord creates a new instance of a HandshakeRecord.
	NewHandshakeRecord() HandshakeRecord

	// MarshalHandshake serializes a HandshakeRecord into a byte slice.
	MarshalHandshake(record HandshakeRecord) ([]byte, error)

	// UnmarshalHandshake deserializes a byte slice into a HandshakeRecord.
	UnmarshalHandshake(data []byte) (HandshakeRecord, error)

	// UnmarshalPing deserializes a byte slice into a PingRecord.
	UnmarshalPing(data []byte) (PingRecord, error)

	// UnmarshalPong deserializes a byte slice into a PongRecord.
	UnmarshalPong(data []byte) (PongRecord, error)

	// NewPongRecord creates a new instance of a PongRecord.
	NewPongRecord() PongRecord

	// NewPingRecord creates a new instance of a PingRecord.
	NewPingRecord() PingRecord

	// MarshalPong serializes a PongRecord into a byte slice.
	MarshalPong(record PongRecord) ([]byte, error)

	// MarshalPing serializes a PingRecord into a byte slice.
	MarshalPing(record PingRecord) ([]byte, error)
}
