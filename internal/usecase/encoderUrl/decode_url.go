package encoderUrl

import (
	"errors"
)

func (rs EncoderUrlStore) DecodeUrl(hash string) (string, error) {
	url, exists := rs.mapUrls[hash]
	if !exists {
		return "", errors.New("URL not found")
	}

	return url, nil
}