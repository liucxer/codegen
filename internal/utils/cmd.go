package utils

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path"
	"runtime"
	"strings"

	"github.com/fatih/color"
)

func StdRun(args ...string) {
	sh := "sh"
	if runtime.GOOS == "windows" {
		sh = "bash"
	}

	stdRun(exec.Command(sh, "-c", strings.Join(args, " ")))
}

func stdRun(cmd *exec.Cmd) {
	cwd, _ := os.Getwd()

	fmt.Fprintf(os.Stdout, "%s %s\n", color.CyanString(path.Join(cwd, cmd.Dir)), strings.Join(cmd.Args, " "))

	{
		stdoutPipe, err := cmd.StdoutPipe()
		if err != nil {
			panic(err)
		}
		go scanAndStdout(bufio.NewScanner(stdoutPipe))
	}
	{
		stderrPipe, err := cmd.StderrPipe()
		if err != nil {
			panic(err)
		}
		go scanAndStderr(bufio.NewScanner(stderrPipe))
	}

	if err := cmd.Run(); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}

func scanAndStdout(scanner *bufio.Scanner) {
	for scanner.Scan() {
		fmt.Fprintln(os.Stdout, scanner.Text())
	}
}

func scanAndStderr(scanner *bufio.Scanner) {
	for scanner.Scan() {
		fmt.Fprintln(os.Stderr, scanner.Text())
	}
}
