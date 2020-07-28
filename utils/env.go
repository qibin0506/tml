package utils

import (
	"strings"
)

func CheckEnv() {
	if checkPython() {
		checkInstallTensorFlow()	
	}
}

func checkPython() bool {
	var notCommand = "executable file not found"
	rst := ExecCommandAndReturn("python")
	
	if strings.Contains(rst, notCommand) {
		println("PYTHON3 WAS *NOT* INSTALLED.")
		println("please visit https://www.python.org/downloads/ to download and install python.")
		return false
	}

	println("1. PYTHON3 WAS INSTALLED.\n")
	return true
}

func checkInstallTensorFlow() {
	var notFound = "no such file or directory"
	var notCommand = "executable file not found"
	var foundTf = "tensorflow"

	var pips = []string{"pip", "pip3", "pip3.5"}
	var checkTf = []string{"show", "tensorflow"}
	var installTf = []string{"install", "tensorflow==2.1.0"}

	var tfFound bool = false
	var pipCmd string

	for _, pip := range pips {
		rst := ExecCommandAndReturn(pip, checkTf...)
		if strings.Contains(rst, notFound) || strings.Contains(rst, notCommand) {
			continue
		}

		if strings.Contains(rst, foundTf) {
			tfFound = true
			println("2. TENSORFLOW WAS INSTALLED. TENSORFLOW INFO:\n")
			println(rst)
			break
		}

		pipCmd = pip
	}

	if !tfFound && pipCmd != "" {
		println("START INSTALl TENSORFLOW...\n")
		ExecCommand(pipCmd, installTf...)
	}

	println("\nCHECK FINISHED. SEE LOG ABOVE TO ENSURE SUCCESS INSTALLED TENSORFLOW AND VERSION >= 2.0.0\n")
}