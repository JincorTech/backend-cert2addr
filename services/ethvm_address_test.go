package services

import (
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

func TestSuccessEthereumAddress(t *testing.T) {
	val, err := GetCIDFromPem([]byte(defaultCert))

	require.NotEmpty(t, val)
	require.NoError(t, err)

	address, err := GetEtheriumLikeAddressFromCid([]byte(val))

	require.NoError(t, err)
	require.Equal(t, "0f5a50a087abd0820840af418cbb01f229b5287e", address)
}

func TestFailCidValue(t *testing.T) {
	val, err := GetCIDFromPem([]byte{0})

	require.Empty(t, val)
	require.EqualError(t, err, "Expecting a PEM-encoded X509 certificate; PEM block not found")
}
