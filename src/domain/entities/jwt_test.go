package entities

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestJWTConstants(t *testing.T) {
	assert.EqualValues(t, 24, expirationTime)
}

func TestGetNewJWT(t *testing.T) {
	assert := assert.New(t)

	jwt := GetNewJWT()
	assert.False(jwt.IsExpired(), "new JWT should not be expired")
	assert.EqualValues("", jwt.AccessToken, "new JWT should not have access token")
	assert.EqualValues(0, jwt.UserId, "new JWT should not associated with a user")
}

func TestJWTIsExpired(t *testing.T) {
	assert := assert.New(t)
	jwt := JWT{}

	assert.True(jwt.IsExpired(), "empty jwt should be expired by default")
	jwt.Expires = time.Now().UTC().Add(3 * time.Hour).Unix()
	assert.False(jwt.IsExpired(), "jwt expiring 3 hour from now should NOT be expired")
}
