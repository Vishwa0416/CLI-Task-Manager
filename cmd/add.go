package cmd

import (
    "fmt"
    "github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
    Use:   "add [task]",
    Short: "Add a new task",
    Args:  cobra.MinimumNArgs(1),
    Run: func(cmd *cobra.Command, args []string) {
        task := args[0]
        fmt.Println("Task added:", task)
        // TODO: Save task to file/db
    },
}

func init() {
    rootCmd.AddCommand(addCmd)
}
