package git_controller

import (
	"fmt"
	"log"
	"strings"

	"github.com/Minnek-Digital-Studio/cominnek/pkg/shell"
)

func CheckIfBranch(branch string) bool {
	currentBranch := GetCurrentBranch();

	if(currentBranch != branch) {
		fmt.Println("This is not " + branch + " branch")
		return false
	}

	return true
}

func CheckChangesFromOrigin() bool {
	err, out, errout := shell.Out("git status");

	if err != nil {
		fmt.Println(out)
		fmt.Println(errout)
		log.Fatal(errout)
	}

	ready := strings.Contains(out, "Your branch is up to date")
	return !ready
}