package netio

import (
	"fmt"
	"os"
)

var isterminal bool

func init() {
	isterminal = IsTerminal(os.Stdout)
}

func IsTerminal(file *os.File) bool {
	st, _ := file.Stat()
	if st == nil {
		return false
	}
	return (st.Mode() & os.ModeCharDevice) == os.ModeCharDevice
}

func Log(v ...interface{}) {
	if isterminal {
		fmt.Println(v...)
	}
}
