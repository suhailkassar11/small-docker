package runtime

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

func Run(command []string) {
	args := append([]string{"child"}, command...)
	cmd := exec.Command("/proc/self/exe", args...)
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWPID | syscall.CLONE_NEWNS | syscall.CLONE_NEWUTS,
	}
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	fmt.Println("RUN")
	err := cmd.Run()
	if err != nil {

		panic(err)
	}
}

func Child(command []string) {

	fmt.Println("CHILD")

	syscall.Sethostname([]byte("tiny-docker"))

	syscall.Mount("", "/", "", syscall.MS_PRIVATE|syscall.MS_REC, "")

	syscall.Mount("proc", "/proc", "proc", 0, "")

	if err := syscall.Exec(command[0], command, os.Environ()); err != nil {
		panic(err)
	}
}
