package functions

import (
	"io"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

func SimpleLS(w io.Writer, args []string, useColor bool) bool {
	hadErr := false

	// no args means list current directory
	if len(args) == 0 {
		return listDir(w, ".", useColor, false)
	}

	files := make([]string, 0)
	dirs := make([]string, 0)

	// split args into files and dirs
	for _, t := range args {
		info, err := os.Lstat(t)
		if err != nil {
			printAccessErr(t, err)
			hadErr = true
			continue
		}

		if info.IsDir() {
			dirs = append(dirs, t)
		} else {
			files = append(files, t)
		}
	}

	// sort both groups
	sort.Strings(files)
	sort.Strings(dirs)

	// print files first
	for _, f := range files {
		name := filepath.Base(f)

		out, err := ColorizePath(f, name, useColor)
		if err != nil {
			printAccessErr(f, err)
			hadErr = true
			continue
		}

		writeLine(w, out)
	}

	// blank line between file list and dir listings if both exist
	if len(files) > 0 && len(dirs) > 0 {
		writeLine(w, "")
	}

	// only print dir headers when more than one dir target
	multiDirs := len(dirs) > 1
	for i, d := range dirs {
		dirHadErr := listDir(w, d, useColor, multiDirs)
		if dirHadErr {
			hadErr = true
		}

		// blank line between directory blocks
		if multiDirs && i != len(dirs)-1 {
			writeLine(w, "")
		}
	}

	return hadErr
}

func listDir(w io.Writer, dir string, useColor bool, printHeader bool) bool {
	hadErr := false

	if printHeader {
		writeLine(w, dir+":")
	}

	entries, err := os.ReadDir(dir)
	if err != nil {
		printAccessErr(dir, err)
		return true
	}

	entries = dirFilter(entries)

	names := make([]string, 0, len(entries))
	for _, e := range entries {
		names = append(names, e.Name())
	}

	sort.Strings(names)

	for _, name := range names {
		full := filepath.Join(dir, name)

		out, err := ColorizePath(full, name, useColor)
		if err != nil {
			printAccessErr(full, err)
			hadErr = true
			continue
		}

		writeLine(w, out)
	}

	return hadErr
}

func dirFilter(entries []os.DirEntry) []os.DirEntry {
	out := make([]os.DirEntry, 0, len(entries))
	for _, e := range entries {
		n := e.Name()
		if strings.HasPrefix(n, ".") {
			continue
		}
		out = append(out, e)
	}
	return out
}

func writeLine(w io.Writer, s string) {
	io.WriteString(w, s)
	io.WriteString(w, "\n")
}
