package encoders

type SimpleEncoder struct {
}

func New() *SimpleEncoder {
	var sh = &SimpleEncoder{}
	return sh
}

// shortener algo
func (s *SimpleEncoder) Encode(index int32) (string, error) {
	return "", nil
}

// shortener algo
func (s *SimpleEncoder) Decode(code int32) (int32, error) {
	return 0, nil
}
