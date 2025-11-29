package cmd

import (
    "github.com/spf13/cobra"
    "cli-task-manager/internal/tasks"
)

var addCmd = &cobra.Command{
    Use:   "add [task]",
    Short: "Add a new task",
    Args:  cobra.MinimumNArgs(1),
    Run: func(cmd *cobra.Command, args []string) {
        tasks.AddTask(args[0])
    },
}

func init() {
    rootCmd.AddCommand(addCmd)
}
