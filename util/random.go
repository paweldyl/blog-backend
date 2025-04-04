package util

import (
	"math/rand"
	"strings"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

func RandomString(n int) string {
	var letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	var sb strings.Builder
	b := len(letters)

	for i := 0; i < n; i++ {
		c := letters[rand.Intn(b)]
		sb.WriteByte(c)
	}

	return sb.String()
}

func RandomOwner() string {
	return RandomString(6)
}

func RandomID(t *testing.T) uuid.UUID {
	id, err := uuid.NewRandom()
	require.NoError(t, err)
	return id
}
