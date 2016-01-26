package commands

import "github.com/abiosoft/ishell"

var shell *ishell.Shell

func Init(_shell *ishell.Shell) {
	shell = _shell

	shell.Register("show", show)
	shell.Register("list", list)
	shell.Register("help", help)
}
