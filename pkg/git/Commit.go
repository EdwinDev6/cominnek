package git

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/Minnek-Digital-Studio/cominnek/controllers"
	git_controller "github.com/Minnek-Digital-Studio/cominnek/controllers/git"
	"github.com/Minnek-Digital-Studio/cominnek/controllers/loading"
	"github.com/Minnek-Digital-Studio/cominnek/pkg/shell"
	"github.com/fatih/color"
)

func _commit(msg string, body string, ctype string, scope string, ticket string) string {
	cmd := git_controller.Commit(msg, body, ctype, ticket, scope)
	out, outErr, err := shell.Out(cmd)

	if err != nil {
		loading.Stop()

		if strings.Contains(out, "nothing to commit") {
			fmt.Println(out)
			fmt.Println("Aborting commit...")

			os.Exit(1)
		} else {
			fmt.Println(outErr)
			log.Fatal("Commit failed")
		}
	}

	return out
}

func _checkTicket(ticket string) string {
	if ticket == "" {
		loading.Stop()
		if !controllers.Confirm("No ticket number found. Commit anyway?", false) {
			fmt.Println("Aborting commit")
			os.Exit(0)
		}

		loading.Start("Commiting files ")
	}

	return ticket
}

func Commit(msg string, body string, ctype string, scope string) {
	loading.Start("Commiting files ")
	currentBranch := git_controller.GetCurrentBranch()

	if strings.HasPrefix(currentBranch, "bugfix/") {
		if ctype == "feat" {
			loading.Stop()
			color.HiRed("Error:")
			log.Fatal("Bugfix branch cannot have a feature commit")
			os.Exit(1)
		}
	}

	ticket := _checkTicket(git_controller.GetTicketNumber())
	out := _commit(msg, body, ctype, scope, ticket)

	loading.Stop()
	fmt.Println(out)
}

func CommitWithoutTicket(msg string, body string, ctype string, scope string) {
	loading.Start("Commiting files ")
	out := _commit(msg, body, ctype, scope, "")

	loading.Stop()
	fmt.Println(out)
}
