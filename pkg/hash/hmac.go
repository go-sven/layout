package hash

import (
	"crypto/hmac"
	"encoding/hex"
	"hash"
)

// HMACHash demo: hmacStr := hash.HMACHash([]byte("data"), []byte("secret"), sha256.New)
func HMACHash(data []byte, key []byte, hashFunc func() hash.Hash) string {
	h := hmac.New(hashFunc, key)
	h.Write(data)
	return hex.EncodeToString(h.Sum(nil))
}
