package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestHTTPRoutes(t *testing.T) {
	if os.Getenv("DEVEX_TEST_FORCE_FAIL") == "1" {
		t.Fatal("falha controlada para demonstrar que um build com testes reprovados não publica imagem")
	}

	for _, tc := range []struct {
		path string
		want map[string]string
	}{
		{path: "/", want: map[string]string{"application": "http-sample", "message": "hello from devex"}},
		{path: "/healthz", want: map[string]string{"status": "ok"}},
	} {
		t.Run(tc.path, func(t *testing.T) {
			recorder := httptest.NewRecorder()
			request := httptest.NewRequest(http.MethodGet, tc.path, nil)
			newHandler().ServeHTTP(recorder, request)

			if recorder.Code != http.StatusOK {
				t.Fatalf("status = %d; want %d", recorder.Code, http.StatusOK)
			}
			if got := recorder.Header().Get("Content-Type"); got != "application/json; charset=utf-8" {
				t.Fatalf("Content-Type = %q", got)
			}
			var body map[string]string
			if err := json.Unmarshal(recorder.Body.Bytes(), &body); err != nil {
				t.Fatal(err)
			}
			for key, value := range tc.want {
				if body[key] != value {
					t.Errorf("%s = %q; want %q", key, body[key], value)
				}
			}
		})
	}
}
