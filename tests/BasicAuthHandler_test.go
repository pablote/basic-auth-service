package tests

import (
	"basic-auth-service/lib"
	"net/http/httptest"
	"testing"
)

func Test_BasicAuthHandler(t *testing.T) {
	table := []struct {
		expectedUsername string
		expectedPassword string
		authorization    string
		assertCode       int
	}{
		{
			expectedUsername: "admin",
			expectedPassword: "password",
			authorization:    "Basic YWRtaW46cGFzc3dvcmQ=", // admin:password
			assertCode:       200,
		},
		{
			expectedUsername: "admin",
			expectedPassword: "admin",
			authorization:    "Basic YWRtaW46cGFzc3dvcmQ=", // admin:password
			assertCode:       401,
		},
		{
			expectedUsername: "admin",
			expectedPassword: "password",
			authorization:    "",
			assertCode:       401,
		},
	}

	for _, row := range table {
		target := lib.BasicAuthHandler(row.expectedUsername, row.expectedPassword, "", []string{"*"}, []string{"*"})

		w := httptest.NewRecorder()
		r := httptest.NewRequest("GET", "/", nil)
		if len(row.authorization) > 0 {
			r.Header.Set("Authorization", row.authorization)
		}

		target.ServeHTTP(w, r)

		if w.Code != row.assertCode {
			t.Errorf(`w.Code expected=%d actual=%d`, row.assertCode, w.Code)
		}
	}
}

var intResult int

func Benchmark_BasicAuthHandler(b *testing.B) {
	for n := 0; n < b.N; n++ {
		target := lib.BasicAuthHandler("admin", "password", "", []string{"*"}, []string{"*"})

		w := httptest.NewRecorder()
		r := httptest.NewRequest("GET", "/", nil)
		r.Header.Set("Authorization", "Basic YWRtaW46cGFzc3dvcmQ=")

		target.ServeHTTP(w, r)
		intResult = w.Code
	}
}
