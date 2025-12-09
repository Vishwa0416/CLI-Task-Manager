package cmd

import (
	"cli-task-manager/internal/tasks"
	"github.com/spf13/cobra"
)

var showDone bool
var showPending bool

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all tasks",
	Run: func(cmd *cobra.Command, args []string) {
		if showDone {
			tasks.ListDoneTasks()
		} else if showPending {
			tasks.ListPendingTasks()
		} else {
			tasks.ListTasks()
		}
	},
}

func init() {
	listCmd.Flags().BoolVar(&showDone, "done", false, "Show only completed tasks")
	listCmd.Flags().BoolVar(&showPending, "pending", false, "Show only pending tasks")
	rootCmd.AddCommand(listCmd)
}
