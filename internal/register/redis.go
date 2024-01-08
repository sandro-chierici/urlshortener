package register

import (
	"errors"
)

type RedisRegister struct {
	// in memory register only for test purpose
	urls map[string]string
}

func NewRedisRegister() *RedisRegister {
	var rr = new(RedisRegister)
	rr.urls = make(map[string]string)
	return rr
}

func (r *RedisRegister) SetUrl(url string, code string) error {

	r.urls[code] = url
	return nil
}

func (r *RedisRegister) GetShortened(code string) (string, error) {

	// check if url exists
	var url, exists = r.urls[code]
	if !exists {
		return "", errors.New("%s not exists")
	} else {
		return url, nil
	}
}
