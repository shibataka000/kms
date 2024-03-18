package openssl

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
)

type AES256CBCOption struct {
	Encrypt bool
	Decrypt bool
	In      []byte
	Pass    []byte
	PBKDF2  bool
	Iter    uint64
}

func AES256CBC(opts AES256CBCOption) ([]byte, error) {
	args := []string{"aes-256-cbc"}

	if opts.Encrypt {
		args = append(args, "-e")
	}
	if opts.Decrypt {
		args = append(args, "-d")
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
		f, err := createTemp(opts.Pass)
		if err != nil {
			return nil, err
		}
		defer os.Remove(f.Name())
		args = append(args, "-pass", fmt.Sprintf("file:%s", f.Name()))
	}

	return exec.Command("openssl", args...).Output()
}

func createTemp(data []byte) (*os.File, error) {
	file, err := os.CreateTemp("", "*")
	if err != nil {
		return nil, err
	}
	err = os.WriteFile(file.Name(), data, 0400)
	return file, err
}
