package server

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"strings"
	"testing"

	"github.com/Kvertinum01/views-counter/internal/app/store"
	"github.com/alicebob/miniredis/v2"
	"github.com/stretchr/testify/assert"
)

func TestHandleRandomImage(t *testing.T) {
	testCases := []struct {
		name             string
		args             string
		wantResponseCode int
	}{
		{
			name:             "good args",
			args:             "?name=someName",
			wantResponseCode: 200,
		},
		{
			name:             "no args",
			args:             "",
			wantResponseCode: 400,
		},
		{
			name:             "too big name",
			args:             fmt.Sprintf("?name=%s", strings.Repeat("t", 17)),
			wantResponseCode: 400,
		},
	}

	s := miniredis.RunT(t)
	rdbConfig := &store.RedisConfig{
		RedisURL: s.Addr(),
	}

	p, _ := filepath.Abs("../../../images/")

	rst := store.NewRedisStore(rdbConfig)
	rst.ConfigureStore()

	server := Server{
		rst:    rst,
		imgDir: p,
	}

	handler := http.HandlerFunc(server.handleRandomImage)

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rec := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", "/counter/"+tc.args, nil)
			handler.ServeHTTP(rec, req)
			assert.Equal(t, tc.wantResponseCode, rec.Result().StatusCode)
		})
	}
}
