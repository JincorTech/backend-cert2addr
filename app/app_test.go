package app

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

const defaultCert = `
-----BEGIN CERTIFICATE-----
MIICBzCCAa6gAwIBAgIUR0wk/DLjm2PCGskw7CRue0uhLaQwCgYIKoZIzj0EAwIw
dzELMAkGA1UEBhMCVVMxEzARBgNVBAgTCkNhbGlmb3JuaWExFjAUBgNVBAcTDVNh
biBGcmFuY2lzY28xGzAZBgNVBAoTEm5ldHdvcmsuamluY29yLmNvbTEeMBwGA1UE
AxMVY2EubmV0d29yay5qaW5jb3IuY29tMB4XDTE3MTEwMjEwNDEwMFoXDTE4MTEw
MjEwNDEwMFowIzEhMB8GA1UEAwwYVXNlcjFAbmV0d29yay5qaW5jb3IuY29tMFkw
EwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAECbIXQcbZ5U7ru0XeIKDNcJcgRPmY3VrM
bGOXW+Yk0s8oIovbMWtgEZ/pdOZbynIGm8GT7OkQHLL7sxDY1mSqgqNsMGowDgYD
VR0PAQH/BAQDAgeAMAwGA1UdEwEB/wQCMAAwHQYDVR0OBBYEFPgn4/MewRsikCTt
U1cLNQHnz7TqMCsGA1UdIwQkMCKAIG1tY5pkJ0JpckTk7JtUm5pxZkKnS0IQdhIh
0PGVvRD7MAoGCCqGSM49BAMCA0cAMEQCIGc/2GTVOxaBgqQGVw3JZslyh10Ul4eo
poqbrgGwyzyeAiAEy9uFn48L2pe7YIOG8Byg98VVyJZQ0l9pzJOdApT/Rw==
-----END CERTIFICATE-----
`

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	app := Application{}
	app.Initialize("", "")

	rr := httptest.NewRecorder()
	app.Router.ServeHTTP(rr, req)

	return rr
}

func TestInvalidGetAddressCall(t *testing.T) {
	req, _ := http.NewRequest("POST", "/api/certificates/actions/getaddress", bytes.NewBuffer(nil))
	response := executeRequest(req)

	require.Equal(t, http.StatusBadRequest, response.Code)
	require.Equal(t, []byte(`{"error":"Invalid request payload","status":400}`), response.Body.Bytes())
}

func TestSuccessGetAddressCall(t *testing.T) {
	reqStrData := CertificateAddressRequest{
		Pem: defaultCert,
	}
	reqBytes, _ := json.Marshal(&reqStrData)

	req, _ := http.NewRequest("POST", "/api/certificates/actions/getaddress", bytes.NewBuffer(reqBytes))
	response := executeRequest(req)

	require.Equal(t, http.StatusOK, response.Code)
	serverData := ServerResponse{
		Data: &CertificateAddressResponse{},
	}

	require.NoError(t, json.Unmarshal(response.Body.Bytes(), &serverData))
	require.Equal(t, `0f5a50a087abd0820840af418cbb01f229b5287e`, serverData.Data.(*CertificateAddressResponse).Address)
}
