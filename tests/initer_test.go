package tests

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMain(t *testing.T) {
	require.NotNil(t, TestApp, "TestApp should not be nil")
}
