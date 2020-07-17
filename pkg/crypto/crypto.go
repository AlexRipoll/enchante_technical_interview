package crypto

type Hashing interface {
	Hash(stringToHash string) (string, error)
	CheckHash(stringToHash string, hash string) error
}