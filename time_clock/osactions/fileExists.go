package osactions

import (
	"fmt"
	"os"
)

func FileExists(Filename string) bool {
	_, Err := os.Stat(Filename)
	fmt.Println(Filename)
	return os.IsNotExist(Err)
}
