// +build !windows

package commands

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnknowShell(t *testing.T) {
	originalShell := os.Getenv("SHELL")
	os.Setenv("SHELL", "")
	defer os.Setenv("SHELL", originalShell)
	shell, err := detectShell()
	fmt.Println(shell)
	assert.Equal(t, err, ErrUnknownShell)
	assert.Equal(t, "", shell)
}
