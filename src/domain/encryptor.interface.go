package domain

// IEncryptor is a interface for a encryptor object
type IEncryptor interface {
	CreateHash(aString string) (string, error)
	Validate(aHash, aString string) bool
}
