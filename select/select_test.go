package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestWebRacer(t *testing.T) {

	t.Run("url response within 10 second limit", func(t *testing.T) {
		slowServer := createDelayedServer(40 * time.Millisecond)
		fastServer := createDelayedServer(0 * time.Millisecond)

		defer slowServer.Close()
		defer fastServer.Close()

		expected := fastServer.URL
		got, err := Racer(slowServer.URL, fastServer.URL)

		if err != nil {
			t.Errorf("got an unexpected error: %v", err)
		}

		if got != expected {
			t.Errorf("expected %q but got %q", expected, got)
		}
	})

	t.Run("should error if response time exceeds timeout limit", func(t *testing.T) {
		server := createDelayedServer(30 * time.Millisecond)
		defer server.Close()

		_, err := ConfigurableRacer(server.URL, server.URL, 20*time.Millisecond)
		if err == nil {
			t.Errorf("expected error, but got %v", err)
		}
	})

}

func createDelayedServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}
