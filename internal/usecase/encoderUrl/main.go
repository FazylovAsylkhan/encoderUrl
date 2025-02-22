package encoderUrl

type EncoderUrlStore struct {
	mapUrls map[string]string
}

var rs EncoderUrlStore

func Init() *EncoderUrlStore {
	rs = EncoderUrlStore{
		mapUrls: make(map[string]string),
	}
	return &rs
}