var listCmd = &cobra.Command{
    Use:   "list",
    Short: "List all tasks",
    Run: func(cmd *cobra.Command, args []string) {
        fmt.Println("Listing tasks...")
        // TODO: Implement reading from file
    },
}
