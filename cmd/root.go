package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "mybatis-code-generator",
	Short: "A Mybatis & JPA code generator",
	Long: `It is a command line tool for Spring in Java.
It contains: Mybatis Mapper/Mybatis Mapper XML/Entity/JPA Repository.
The GitHub site at https://github.com/billcoding/mybatis-code-generator`,
	Run: func(cmd *cobra.Command, args []string) {
		if versionFlag {
			PrintVersion(false)
		}
	},
}

var versionFlag bool

// Execute executes the root command.
func Execute() error {
	rootCmd.PersistentFlags().BoolVarP(&versionFlag, "version", "v", false, "version")
	return rootCmd.Execute()
}
