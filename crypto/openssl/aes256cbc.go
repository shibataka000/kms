package openssl

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
)

// AES256CBCOption is option for `openssl aes-256-cbc` command.
type AES256CBCOption struct {
	Encrypt bool
	Decrypt bool
	In      []byte
	Pass    []byte
	Salt    bool
	PBKDF2  bool
	Iter    uint64
}

// AES256CBC exec `openssl aes-256-cbc` command.
func AES256CBC(opts AES256CBCOption) ([]byte, error) {
	args := []string{"aes-256-cbc"}
	env := []string{}

	if opts.Encrypt {
		args = append(args, "-e")
	}
	if opts.Decrypt {
		args = append(args, "-d")
	}
	if opts.Salt {
		args = append(args, "-salt")
	} else {
		args = append(args, "-nosalt")
	}
	if opts.PBKDF2 {
		args = append(args, "-pbkdf2")
	}
	if opts.Iter != 0 {
		args = append(args, "-iter", strconv.FormatUint(opts.Iter, 10))
	}
	if opts.In != nil {
		f, err := createTemp(opts.In)
		if err != nil {
			return nil, err
		}
		defer os.Remove(f.Name())
		args = append(args, "-in", f.Name())
	}
	if opts.Pass != nil {
		envName := "OPENSSL_PASS"
		env = append(env, fmt.Sprintf("%s=%s", envName, string(opts.Pass)))
		args = append(args, "-pass", fmt.Sprintf("env:%s", envName))
	}

	cmd := exec.Command("openssl", args...)
	cmd.Env = append(cmd.Env, env...)

	return cmd.Output()
}

// createTemp creates a new temporary file and write data into it.
func createTemp(data []byte) (*os.File, error) {
	file, err := os.CreateTemp("", "*")
	if err != nil {
		return nil, err
	}
	err = os.WriteFile(file.Name(), data, 0400)
	return file, err
}
