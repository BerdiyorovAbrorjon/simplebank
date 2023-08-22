package token

import (
	"testing"
	"time"

	"github.com/BerdiyorovAbrorjon/simplebank/util"
	"github.com/dgrijalva/jwt-go"
	"github.com/stretchr/testify/require"
)

func TestJWTToken(t *testing.T) {
	jwtMaker, err := NewJWTMaker(util.RandomString(32))
	require.NoError(t, err)

	username := util.RandomUsername()
	duration := time.Minute
	issuedAt := time.Now()
	expireAt := issuedAt.Add(duration)

	token, err := jwtMaker.CreateToken(username, duration)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	payload, err := jwtMaker.VerifyToken(token)

	require.NoError(t, err)
	require.NotEmpty(t, payload)
	require.Equal(t, username, payload.Username)
	require.WithinDuration(t, issuedAt, payload.IssuedAt, time.Second)
	require.WithinDuration(t, expireAt, payload.ExpireAt, time.Second)
	require.NotZero(t, payload.ID)
}

func TestExpireJWTToken(t *testing.T) {
	jwtMaker, err := NewJWTMaker(util.RandomString(32))
	require.NoError(t, err)

	username := util.RandomUsername()
	duration := -time.Minute

	token, err := jwtMaker.CreateToken(username, duration)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	payload, err := jwtMaker.VerifyToken(token)

	require.Error(t, err)
	require.EqualError(t, err, ErrExpireToken.Error())
	require.Nil(t, payload)
}

func TestInvalidJWTTokenAlgNone(t *testing.T) {
	payload, err := NewPayload(util.RandomUsername(), time.Minute)
	require.NoError(t, err)

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodNone, payload)
	token, err := jwtToken.SignedString(jwt.UnsafeAllowNoneSignatureType)
	require.NoError(t, err)

	jwtMaker, err := NewJWTMaker(util.RandomString(32))
	require.NoError(t, err)

	payload, err = jwtMaker.VerifyToken(token)
	require.Error(t, err)
	require.EqualError(t, err, ErrInvalidToken.Error())
	require.Nil(t, payload)
}
