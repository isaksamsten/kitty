// License: GPLv3 Copyright: 2022, Kovid Goyal, <kovid at kovidgoyal.net>

package main

import (
	"os"

	"kitty/kittens/ssh"
	"kitty/tools/cli"
	"kitty/tools/cmd/completion"
	"kitty/tools/cmd/tool"
)

func main() {
	krm := os.Getenv("KITTY_KITTEN_RUN_MODULE")
	os.Unsetenv("KITTY_KITTEN_RUN_MODULE")
	switch krm {
	case "ssh_askpass":
		ssh.RunSSHAskpass()
		return
	}
	root := cli.NewRootCommand()
	root.ShortDescription = "Fast, statically compiled implementations for various kittens (command line tools for use with kitty)"
	root.Usage = "command [command options] [command args]"
	root.Run = func(cmd *cli.Command, args []string) (int, error) {
		cmd.ShowHelp()
		return 0, nil
	}

	tool.KittyToolEntryPoints(root)
	completion.EntryPoint(root)

	root.Exec()
}
