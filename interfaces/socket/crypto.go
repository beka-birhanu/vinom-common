package i

// Symmetric defines an interface for symmetric cryptography.
type Symmetric interface {
	// Encrypt encrypts the given plaintext using the provided key.
	// Returns the ciphertext or an error if encryption fails.
	Encrypt(p []byte, key []byte) ([]byte, error)

	// Decrypt decrypts the given ciphertext using the provided key.
	// Returns the original plaintext or an error if decryption fails.
	Decrypt(c []byte, key []byte) ([]byte, error)
}

// Asymmetric defines an interface for asymmetric cryptography.
type Asymmetric interface {
	// Encrypt encrypts the given plaintext using the provided key.
	// Returns the ciphertext or an error if encryption fails.
	Encrypt(p []byte, key []byte) ([]byte, error)

	// Decrypt decrypts the given ciphertext.
	// Returns the original plaintext or an error if decryption fails.
	Decrypt(c []byte) ([]byte, error)

	// GetPublicKey returns the public key associated with the asymmetric encryption.
	GetPublicKey() []byte
}

// HMAC defines an interface for Hash-based Message Authentication Code (HMAC).
type HMAC interface {
	// Sign generates an HMAC signature for the given data.
	// Additional optional keys can be provided for multi-key signing.
	Sign(data []byte, keys ...[]byte) []byte

	// Compare verifies whether the provided signature matches the expected HMAC.
	// Returns true if they match, otherwise false.
	Compare(expected []byte, actual []byte) bool
}
