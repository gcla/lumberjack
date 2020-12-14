// Log the panic under unix to the log file
// +build !windows,!solaris,!plan9,!js

package lumberjack

import (
	"fmt"
	"os"
	"syscall"
)

// redirectStderr to the file passed in
func redirectStderr(f *os.File) {
	if err := syscall.Dup2(int(f.Fd()), syscall.Stderr); err != nil {
		fmt.Errorf("can't dup2 logfile and stderr: %s", err)
	}
}
