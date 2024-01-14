package types

// Register definition
type Register interface {
	SetUrl(url string) (string, error)
	GetUrl(code string) (string, error)
}

type Encoder interface {
	Encode(index int32) (string, error)
	Decode(code int32) (int32, error)
}
