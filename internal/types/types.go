package types

// Register definition
type Register interface {
	SetUrl(url string, code string) error
	GetShortened(url string) (string, error)
}
