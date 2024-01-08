package main

import (
	"github.com/spf13/cobra"
	new2 "merak/cmd/merak/new"
)

var Cmd = &cobra.Command{
	Use: "merak",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func main() {
	Cmd.AddCommand(new2.NewCmd)
	Cmd.Execute()
}
