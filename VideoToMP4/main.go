/**
 * ========================================================================
 *  File: main.go
 *  Creator: Dmitri G.
 *  Date: 2018-12-12
 *  Description: This is a main application file written in Golang.
 *  It used as program wrapper for Hadbrake-Cli
 * ========================================================================
 */

package main

import (
	"bytes"
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
)

// *** FUNCTIONS ***
func checkExeExists(exe string) {
	// Check if main program preinstalled on local system
	_, err := exec.LookPath(exe)
	if err != nil {
		color.Set(color.FgRed, color.Bold) // Set color
		fmt.Printf("Didn't find CORE executable\n")
		color.Unset() // Unset color
		log.Fatal("BYE\n")
	}
}

func execute(exe string) {
	// Encoding algorithm
	var stdoutBuf, stderrBuf bytes.Buffer
	args := os.Args[1:]

	if len(args) < 1 {
		color.Set(color.FgYellow) // Set yellow color
		fmt.Println("Arguments ARE EMPTY!")
		fmt.Println("USE: " + exe + " [FileName]")
		fmt.Println()
		fmt.Println("OR PLEASE DROP THE VIDEO FILE")
		fmt.Println("=> ONTO THE PROGRAM ICON in WINDOWS!!!")
		fmt.Println()
		fmt.Println("Author: Dmitri G. (2019)")
		fmt.Println()
		color.Unset() // Unset color
		return
	}

	color.Set(color.FgYellow)                                // Set yellow color
	fmt.Println("File encoding in PROGRESS. Wait PLEASE...") // Print message
	color.Unset()                                            // Unset color
	dir, fn := path.Split(args[0])

	name := strings.TrimSuffix(fn, filepath.Ext(fn))
	outname := dir + name + "_H264_AAC.mp4"

	// Invoke the command with arguments
	cmd := exec.Command(exe, "-i", args[0], "-o", outname,
		"-e", encoder, "-q", quality, "-B", a_bitrate, "--crop", crop)

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
		fmt.Println("PRESS THE `ENTER` KEY TO TERMINATE THE CONSOLE SCREEN!")
		color.Unset() // Unset color
		fmt.Scanln()
	} else {
		checkExeExists(core)
		execute(core)
	}
}
