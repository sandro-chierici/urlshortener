package register

import (
	"bytes"
	"encoding/base64"
	"encoding/binary"
	"math"
	"strings"
	"sync"
	"urlshortener/v2/internal/types"
)

type SimpleRegister struct {
	// in memory register only for test purpose
	// and as redis does i need to handle concurrency writings
	lck     sync.Mutex
	urls    []string
	encoder types.Encoder
}

/*
Build register
*/
func NewSimpleRegister(enc types.Encoder) *SimpleRegister {
	var rr = new(SimpleRegister)

	rr.encoder = enc
	// starts with chunsks of shorts
	rr.urls = make([]string, 0, math.MaxInt16)

	return rr
}

/*
Translate from  url to code
*/
func (r *SimpleRegister) SetUrl(url string) (string, error) {

	r.lck.Lock()
	defer r.lck.Unlock()

	// ==== save to new item =====
	slice := append(r.urls, url)

	// ==== encode from index to code ===

	// translate from int32 to byte array
	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.LittleEndian, len(slice))
	if err != nil {
		return "", err
	}

	// encode in base64 the byte array translation of index
	encoded := base64.StdEncoding.EncodeToString(buf.Bytes())

	// we have here always 4 bytes, translated in Base64 we have *always* double == at the end
	// so get rid of it
	return strings.Replace(encoded, "=", "", -1), nil
}

/*
Translate from Code to url
*/
func (r *SimpleRegister) GetUrl(code string) (string, error) {

	// decode from code to index

	return "", nil
}
