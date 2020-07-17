package uuidv4

type Service interface {
	Generate() (string, error)
	Validate(string) error
}
