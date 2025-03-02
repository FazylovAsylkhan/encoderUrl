package handler

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/FazylovAsylkhan/encoderUrl/internal/config"
	"github.com/stretchr/testify/suite"
)

type HandlerTestSuite struct {
	suite.Suite
	h handler
}
func (s *HandlerTestSuite) SetupSuite() {
	cfg := config.Config{
		BaseURL: "",
		Address: "",
	}
	_, handler := Init(&cfg)
	s.h = *handler
}

func (s *HandlerTestSuite) TestHandlerUrl() {
	s.Run("GET request with wrong method", func() {
		payload := strings.NewReader("profiles")
		req, err := http.NewRequest("GET", "/", payload)
		s.Require().NoError(err)

		req.Header.Set("Content-Type", "text/plain")
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(s.h.handlerUrl)
		handler.ServeHTTP(rr, req)

		s.Assert().Equal(http.StatusBadRequest, rr.Code, "POST handler returned wrong status code")
		expected := "Only POST method is allowed\n"
		s.Assert().Equal(expected, rr.Body.String(), "Unexpected error text")
	})
	s.Run("POST request without payload", func() {
		req, err := http.NewRequest("POST", "/", nil)
		s.Require().NoError(err)

		req.Header.Set("Content-Type", "text/plain")
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(s.h.handlerUrl)
		handler.ServeHTTP(rr, req)

		s.Assert().Equal(http.StatusBadRequest, rr.Code, "POST handler returned wrong status code")
		expected := "Empty body\n"
		s.Assert().Equal(expected, rr.Body.String(), "Unexpected error text")
	})
	s.Run("POST request with success", func() {
		payload := strings.NewReader("profiles")
		req, err := http.NewRequest("POST", "/", payload)
		s.Require().NoError(err)

		req.Header.Set("Content-Type", "text/plain")
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(s.h.handlerUrl)
		handler.ServeHTTP(rr, req)

		s.Assert().Equal(http.StatusCreated, rr.Code, "POST handler returned wrong status code")
		expected := "/d9707283"
		s.Assert().Equal(expected, rr.Body.String(), "POST handler returned unexpected body")
	})
}

func (s *HandlerTestSuite) TestHandlerHash() {
	s.Run("POST request with wrongMethod", func() {
		link := "profiles"
		payload := strings.NewReader(link)
		reqPost, err := http.NewRequest("POST", "/", payload)
		s.Require().NoError(err)
		reqPost.Header.Set("Content-Type", "text/plain")
		rrPost := httptest.NewRecorder()
		handler := http.HandlerFunc(s.h.handlerUrl)
		handler.ServeHTTP(rrPost, reqPost)
		url := rrPost.Body.String()

		reqGet, err := http.NewRequest("POST", url, nil)
		s.Require().NoError(err)

		rrGet := httptest.NewRecorder()
		handler = http.HandlerFunc(s.h.handlerHash)
		handler.ServeHTTP(rrGet, reqGet)
		

		s.Assert().Equal(http.StatusBadRequest, rrGet.Code, "GET handler returned wrong status code")
		result := rrGet.Body.String()
		expected := "Only GET method is allowed\n"
		s.Assert().Equal(expected, result, "Unexpected error text")
	})
	s.Run("GET request with success", func() {
		link := "profiles"
		payload := strings.NewReader(link)
		reqPost, err := http.NewRequest("POST", "/", payload)
		s.Require().NoError(err)
		reqPost.Header.Set("Content-Type", "text/plain")
		rrPost := httptest.NewRecorder()
		handler := http.HandlerFunc(s.h.handlerUrl)
		handler.ServeHTTP(rrPost, reqPost)

		url := rrPost.Body.String()
		reqGet, err := http.NewRequest("GET", url, nil)
		s.Require().NoError(err)

		rrGet := httptest.NewRecorder()
		handler = http.HandlerFunc(s.h.handlerHash)
		handler.ServeHTTP(rrGet, reqGet)
		

		s.Assert().Equal(http.StatusTemporaryRedirect, rrGet.Code, "GET handler returned wrong status code")
		expected := fmt.Sprintf("%v/%v", s.h.cfg.BaseURL, link)
		result := strings.Join(rrGet.Header()["Location"], ", ")
		s.Assert().Equal(expected, result, "GET handler returned unexpected body")
	})
}

func TestHandlerTestSuite(t *testing.T) {
	suite.Run(t, new(HandlerTestSuite))
}
