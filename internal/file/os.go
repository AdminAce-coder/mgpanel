package file

import (
	"os"
)

type op interface {
	GetworkspaceAbs() *op
	Creatfile() *os.File
}
