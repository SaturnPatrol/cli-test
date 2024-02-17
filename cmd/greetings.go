/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// greetingsCmd represents the greetings command
var greetingsCmd = &cobra.Command{
	Use:   "greetings",
	Short: "greets user of app",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		greetUser, _ := cmd.Flags().GetString("name")
		if greetUser != "" {
			GreetUserWithName(greetUser)
		} else {
			GreetUser()
		}
	},
}

func init() {
	rootCmd.AddCommand(greetingsCmd)
	greetingsCmd.PersistentFlags().String("name", "", "greet the given name")
}

func GreetUser() {
	fmt.Println("hi there!!!")
}

func GreetUserWithName(name string) {
	fmt.Printf("Hi, %s, welcome", name)
}
