package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"strings"
	"sync"
	"syscall"

	"github.com/fatih/color"
)

var colors = []color.Attribute{
	color.FgGreen,
	color.BgBlue,
	color.BgCyan,
	color.BgYellow,
	color.BgMagenta,
	color.FgHiGreen,
	color.BgHiBlue,
	color.BgHiCyan,
	color.BgHiYellow,
	color.BgHiMagenta,
}

type colorPrinter func(...interface{}) string

func handleSigterm(sig chan os.Signal, cmd *exec.Cmd) {
	<-sig
	cmd.Process.Signal(syscall.SIGTERM)
}

func spawnProcess(command []string, wg *sync.WaitGroup, printer colorPrinter) error {
	sigterm := make(chan os.Signal, 1)
	signal.Notify(sigterm, syscall.SIGTERM, syscall.SIGINT)
	cmd := exec.Command(command[0], command[1:]...)
	pname := command[1]
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		wg.Done()
		return err
	}
	go handleSigterm(sigterm, cmd)
	scanner := bufio.NewScanner(stdout)
	err = cmd.Start()
	if err != nil {
		wg.Done()
		return err
	}
	for scanner.Scan() {
		fmt.Printf("%s: %s\n", printer(pname), scanner.Text())
	}

	err = cmd.Wait()
	wg.Done()
	return err
}

func performSSH(hostsFiles []string) {
	cmds := [][]string{}
	for _, hf := range hostsFiles {
		hfa := strings.Split(hf, ",")
		cmds = append(cmds, []string{"ssh", hfa[0], "tail -f " + hfa[1]})
	}
	fmt.Println(cmds)
	var wg sync.WaitGroup
	for i, c := range cmds {
		colAttr := colors[i%10]
		printer := color.New(colAttr).SprintFunc()
		go spawnProcess(c, &wg, printer)
		wg.Add(1)
	}
	wg.Wait()
}

func main() {
	if len(os.Args) == 1 {
		color.Red("Usage: stail ssh-host-name,file-path")
		os.Exit(0)
	}
	performSSH(os.Args[1:])
}
