package radix

import (
	"errors"
	"os"
)

func CreateFile(dirname string, filename string) (*os.File, error) {
	dir, dirErr := os.Open(os.Args[1])
	if dirErr != nil {
		return nil, errors.New("")
	}

	stat, _ := dir.Stat()
	if !stat.IsDir() {
		return nil, errors.New("")
	}

	dir.Close()

	file, _ := os.Create(dir.Name() + filename)
	return file, nil
}

func OpenFile(filename string) (*os.File, error) {
	f, fErr := os.Open(filename)
	if fErr != nil {
		return nil, errors.New("")
	}

	st, sErr := f.Stat()
	if sErr != nil || st.IsDir() {
		return nil, errors.New("")
	}

	return f, nil
}
