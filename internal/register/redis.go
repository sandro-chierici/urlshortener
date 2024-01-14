package register

import "urlshortener/v2/internal/types"

type RedisRegister struct {
	encoder types.Encoder
}

/*
Build register
*/
func NewRedisRegister(enc types.Encoder) *RedisRegister {
	var rr = new(RedisRegister)
	rr.encoder = enc
	return rr
}

/*
Translate from  url to code
*/
func (r *RedisRegister) SetUrl(url string) (string, error) {

	return "", nil
}

/*
Translate from Code to url
*/
func (r *RedisRegister) GetUrl(code string) (string, error) {

	// decode from code to index
	return "", nil
}
