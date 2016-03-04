package commands

import "github.com/abiosoft/ishell"

var shell *ishell.Shell

func Init(_shell *ishell.Shell) {
	shell = _shell

	shell.Register("rdel", record_del)
	shell.Register("radd", record_add)
	shell.Register("rupdate", record_upd)
	shell.Register("ztouch", zone_touch)
	shell.Register("zshow", zone_show)
	shell.Register("zlist", zone_list)
	shell.Register("help", help)
}
