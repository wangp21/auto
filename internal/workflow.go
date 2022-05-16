package internal

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"bufio"
	"github.com/fatih/color"
	"path"
)

func StdRun(cmd *exec.Cmd) {
	cwd, _ := os.Getwd()
	fmt.Fprintf(os.Stdout, "%s %s\n",color.CyanString(path.Join(cwd, cmd.Dir)), strings.Join(cmd.Args, " "))
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
		fmt.Fprintln(os.Stderr,color.RedString(">> %s",err.Error()))
		os.Exit(1)
	}
}

func scanAndStdout(scanner *bufio.Scanner) {
	for scanner.Scan() {
		fmt.Fprintln(os.Stdout, color.GreenString(">> %s",scanner.Text()))
	}
}

func scanAndStderr(scanner *bufio.Scanner) {
	for scanner.Scan() {
		fmt.Fprintln(os.Stderr,color.RedString(">> %s", scanner.Text()))
	}
}

type Script []string

func (s Script) IsZero() bool {
	return len(s) == 0
}

func (s Script) String() string {
	return strings.Join(s, " && ")
}

type Workflow struct {
	Name		string
	Scripts		map[string]Script
}

func (w Workflow) WithScripts(key string, scripts ...string) Workflow {
	if w.Scripts == nil {
		w.Scripts = map[string]Script{}
	}
	w.Scripts[key] = append(Script{}, scripts...)
	return w
}

func (w *Workflow) Command(args ...string) *exec.Cmd {
	sh := "sh"
	if runtime.GOOS == "windows" {
		sh = "bash"
	}

	return exec.Command(sh, "-c", strings.Join(args, " "))
}

func (w *Workflow) Run(commands ...*exec.Cmd) {
	for _, cmd := range commands {
		if cmd != nil {
			StdRun(cmd)
		}
	}
}

func (w *Workflow) Execute(args ...string) {
	w.Run(w.Command(args...))
}

func (p *Workflow) RunScript(key string) error {
	if _, ok := p.Scripts[key]; !ok {
		return fmt.Errorf("script %s not defined", key)
	}
	s := p.Scripts[key]
	p.Run(p.Command(s.String()))
	return nil
}
