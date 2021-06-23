package cmd

import (
	"github.com/spf13/cobra"
)

/*
 * APPNAME COMMAND ARG --FLAG
 * example hugo server --port=1313 -- 'server' is a command, and 'port' is a flag
 * example codebase https://github.com/schadokar/my-calc
 */
var rootCmd = &cobra.Command{
	Use:   "fledgectl",
	Short: "FLEDGE CLI Tool",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

/**
 * This is the first function which gets called whenever a package initialize in the golang.
 * https://towardsdatascience.com/how-to-create-a-cli-in-golang-with-cobra-d729641c7177
 * The order of function calls is as follow
 * 	init --> main --> cobra.OnInitialize --> command
 */
func init() {}

func Execute() error {
	return rootCmd.Execute()
}
