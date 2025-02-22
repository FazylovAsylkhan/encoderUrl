package encoderUrl

import (
	"crypto/md5"
	"encoding/hex"
)

func getHashFrom(str string) string {
	hash := md5.Sum([]byte(str))
	return hex.EncodeToString(hash[:])[:8]
}