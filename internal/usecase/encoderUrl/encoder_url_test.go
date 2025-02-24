//go:build testmode

package encoderUrl

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type DecoderUrlSuite struct {
	suite.Suite
	encoderUrl *EncoderUrlStore
}

func (s *DecoderUrlSuite) SetupTest() {
	s.encoderUrl = Init()
}

func TestDecoderUrlSuite(t *testing.T) {
	suite.Run(t, new(DecoderUrlSuite))
}

func (s DecoderUrlSuite) TestGetHashFrom() {
	expected := "d9707283"
	str := "profiles"	
	
	result := getHashFrom(str)
	
	s.Assert().Equal(expected, result)
}

func (s DecoderUrlSuite) TestEncodeUrl() {
	url := "http://www.google.com"
	expected := getHashFrom(url)

	result := s.encoderUrl.EncodeURL(url)
	s.Assert().Equal(url, s.encoderUrl.mapUrls[expected])
	s.Assert().Equal(expected, result)
}

func (s DecoderUrlSuite) TestDecodeUrl() {
	expected := "http://www.google.com"
	hash := getHashFrom(expected)

	s.encoderUrl.mapUrls[hash] = expected

	result, err := s.encoderUrl.DecodeUrl(hash)
	
	s.Assert().Nil(err)
	s.Assert().Equal(expected, result)

	expectedErr := "URL not found"
	hash = "random"

	result, err = s.encoderUrl.DecodeUrl(hash)
	
	s.Assert().NotNil(err)
	s.Assert().Equal(expectedErr, err.Error())
	s.Assert().Equal("", result)
}
