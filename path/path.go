package path

import (
	"fmt"
	"os"
)

type Path struct {
	path string
}

const (
	FILE = iota
	DIRECTORY
)

func (p *Path) IsFileOrDir() (fileOrDir int, ok bool) {
	fi, err := os.Stat(p.path)
	if err != nil {
		fmt.Println(err)
		return -1, false
	}
	switch mode := fi.Mode(); {
	case mode.IsDir():
		// do directory stuff
		return DIRECTORY, true
	case mode.IsRegular():
		// do file stuff
		return FILE, true
	}

	return -1, false
}
