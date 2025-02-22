package encoderUrl

func (rs EncoderUrlStore) EncodeURL(url string) string {
	hash := getHashFrom(url)
	rs.mapUrls[hash] = url
	return hash
}
