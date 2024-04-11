package cmd

import "os"

// write data to stdout or stderr or file.
// If name is 'stdout' or empty, write into stdout. If name is 'stderr', write into stdout. Otherwise, write into file.
func write(name string, data []byte) error {
	switch name {
	case "":
		fallthrough
	case "stdout":
		_, err := os.Stdout.Write(data)
		return err
	case "stderr":
		_, err := os.Stderr.Write(data)
		return err
	default:
		return os.WriteFile(name, data, 0644)
	}
}
