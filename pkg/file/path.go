package file

import (
	"fmt"
	"os"

	herrors "github.com/hohice/higo/pkg/errors"
)

var (
	Separator = os.PathSeparator
)

func MustCreateDir(path string) {
	if err := os.Mkdir(path, 0755); err != nil {
		emsg := fmt.Sprintf("MustCreateDir error,create dir:%s failed", path)
		herrors.Panic(herrors.NewDefWarpError(emsg, err))
		//herrors.Panic(herrors.NewDefineError(emsg))
	}
}
