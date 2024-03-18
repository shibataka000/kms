package openssl

import (
	"os/exec"
	"strconv"
)

type RandOption struct {
	Base64 bool
}

func Rand(num int, opts RandOption) ([]byte, error) {
	args := []string{"rand"}

	if opts.Base64 {
		args = append(args, "-base64")
	}

	args = append(args, strconv.Itoa(num))

	return exec.Command("openssl", args...).Output()
}
