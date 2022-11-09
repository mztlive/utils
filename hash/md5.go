package hash

import (
	"crypto/md5"
	"encoding/hex"
)

// Md5Hash generates a MD5 hash of the given string.
func Md5Hash(plaintext string) (string, error) {
	h := md5.New()
	_, err := h.Write([]byte(plaintext))
	if err != nil {
		return "", err
	}

	md5Ps := hex.EncodeToString(h.Sum(nil))
	return md5Ps, nil
}
