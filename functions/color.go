package functions

import "os"

const (
	blue  = "\033[34m"
	green = "\033[32m"
	reset = "\033[0m"
)

// color rules
// dir is blue
// executable regular file is green
func ColorizePath(path string, name string, useColor bool) (string, error) {
	info, err := os.Lstat(path)
	if err != nil {
		return "", err
	}

	if !useColor {
		return name, nil
	}

	mode := info.Mode()

	// directory
	if info.IsDir() {
		return blue + name + reset, nil
	}

	// executable regular file
	if mode.IsRegular() && (mode&0111) != 0 {
		return green + name + reset, nil
	}

	return name, nil
}
