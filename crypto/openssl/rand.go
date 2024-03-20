package openssl

import (
	"os/exec"
	"strconv"
)

// RandOption is option for `openssl rand` command.
type RandOption struct {
	Base64 bool
}

// Rand exec `openssl rand` command.
func Rand(num int, opts RandOption) ([]byte, error) {
	args := []string{"rand"}

	if opts.Base64 {
		args = append(args, "-base64")
	}

	args = append(args, strconv.Itoa(num))

	return exec.Command("openssl", args...).Output()
}
