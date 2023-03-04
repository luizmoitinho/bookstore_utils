package crypto_utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetSha256(t *testing.T) {
	//assert
	password := "abc123"

	//act
	result := GetSha256(password)

	assert.NotEmpty(t, result)
	assert.EqualValues(t, "6ca13d52ca70c883e0f0bb101e425a89e8624de51db2d2392593af6a84118090", result)
}
