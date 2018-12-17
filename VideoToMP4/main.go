/**
 * ========================================================================
 *  File: main.go
 *  Creator: Dmitri G.
 *  Date: 2018-12-17
 *  Version: 1.0.1
 *  Description: This is a main application file written in Golang.
 *  It used as program wrapper for Hadbrake-Cli
 * ========================================================================
 */

package main

import (
	"bytes"
	"flag"
	"fmt"
	"github.com/fatih/color"
	"io"
	"log"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"runtime"
	"strings"
)

// *** VARIABLES ***
var (
	core      = "HandBrakeCLI"
	encoder   = "x264"
	quality   = "20"
	a_bitrate = "160"
	crop      = "0:0:0:0"
	height    = "720"
)

// *** FUNCTIONS ***
func helpMsg(cmdName string) {
	fmt.Println()
	color.Set(color.FgYellow) // Set yellow color
	fmt.Println("Arguments ARE EMPTY!")
	fmt.Println()
	fmt.Println("USE: " + cmdName + " -hd" + " [FileName] " + ": Resize to 720p")
	fmt.Println("     " + cmdName + "     [FileName] " + ": Size same as source")
	fmt.Println()
	fmt.Println("OR PLEASE DROP THE VIDEO FILE")
	fmt.Println("=> ONTO THE PROGRAM ICON in WINDOWS!!!")
	fmt.Println()
	fmt.Println("Author: Dmitri G. (2019)")
	fmt.Println()
	color.Unset() // Unset color
}

func checkExeExists(cmdName string) {
	// Check if main program preinstalled on local system
	_, err := exec.LookPath(cmdName)
	if err != nil {
		color.Set(color.FgRed, color.Bold) // Set color
		fmt.Printf("Didn't find CORE executable\n")
		color.Unset() // Unset color
		log.Fatal("BYE\n")
	}
}

func execute(cmdName string) {
	// Encoding algorithm
	var stdoutBuf, stderrBuf bytes.Buffer
	var cmdArgs []string

	sizePtr := flag.Bool("hd", false, "Resize to 720p")

	args := os.Args[1:]

	flag.Usage = func() {
		color.Set(color.FgYellow) // Set yellow color
		fmt.Println()
		fmt.Printf("Usage of %s:\n", cmdName)
		fmt.Println()
		fmt.Printf(" -hd [FileName] : Resize to 720p\n")
		fmt.Printf("     [FileName] : Size same as source\n")
		fmt.Println()
		color.Unset() // Unset color
	}

	flag.Parse()

	// Decision Block
	switch {
	case len(args) < 1:
		helpMsg(cmdName)
		return

	case len(args) == 1:
		dir, fn := path.Split(args[0])
		name := strings.TrimSuffix(fn, filepath.Ext(fn))
		outname := dir + name + "_H264_AAC.mp4"
		cmdArgs = []string{"-i", args[0], "-o", outname, "-e", encoder,
			"-q", quality, "-B", a_bitrate, "--crop", crop}

	case len(args) > 2:
		helpMsg(cmdName)
		return
	}

	if *sizePtr {
		dir, fn := path.Split(args[1])
		name := strings.TrimSuffix(fn, filepath.Ext(fn))
		outname := dir + name + "_H264_AAC.mp4"
		cmdArgs = []string{"-i", args[1], "-o", outname, "-e", encoder,
			"-q", quality, "-l", height, "-B", a_bitrate, "--crop", crop}
	}

	color.Set(color.FgYellow)                                // Set yellow color
	fmt.Println("File encoding in PROGRESS. Wait PLEASE...") // Print message
	color.Unset()                                            // Unset color

	// Invoke the command with arguments
	cmd := exec.Command(cmdName, cmdArgs...)

	stdoutIn, _ := cmd.StdoutPipe()
	stderrIn, _ := cmd.StderrPipe()

	var errStdout, errStderr error

	stdout := io.MultiWriter(os.Stdout, &stdoutBuf)
	stderr := io.MultiWriter(os.Stderr, &stderrBuf)
	err := cmd.Start()

	if err != nil {
		log.Fatalf("cmd.Start() failed with '%s'\n", err)
	}

	go func() {
		_, errStdout = io.Copy(stdout, stdoutIn)
	}()

	go func() {
		_, errStderr = io.Copy(stderr, stderrIn)
	}()

	err = cmd.Wait()

	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}

	if errStdout != nil || errStderr != nil {
		log.Fatal("Failed to capture stdout or stderr\n")
	}

	outStr, errStr := string(stdoutBuf.Bytes()), string(stderrBuf.Bytes())
	fmt.Printf("\nout:\n%s\nerr:\n%s\n", outStr, errStr)

	color.Set(color.FgYellow) // Set yellow color
	fmt.Println("Created By: Dmitri G. (2019)")
	fmt.Println("Have a NICE DAY! :)")
	color.Unset() // Unset color
}

// *** MAIN FUNCTION ***
func main() {
	core_w := core + ".exe"

	if runtime.GOOS == "windows" {
		checkExeExists(core_w)
		execute(core_w)

		/*** PAUSE ***/
		color.Set(color.FgMagenta, color.Bold) // Set color
		fmt.Println("PRESS THE `Enter` KEY TO TERMINATE THE CONSOLE SCREEN!")
		color.Unset() // Unset color
		fmt.Scanln()
	} else {
		checkExeExists(core)
		execute(core)
	}
}
