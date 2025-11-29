package cmd

import (
    "github.com/spf13/cobra"
    "cli-task-manager/internal/tasks"
)

var listCmd = &cobra.Command{
    Use:   "list",
    Short: "List all tasks",
    Run: func(cmd *cobra.Command, args []string) {
        tasks.ListTasks()
    },
}

func init() {
    rootCmd.AddCommand(listCmd)
}
