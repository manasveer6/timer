package main

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"

	"timer/tracker"
)

var rootCmd = &cobra.Command{
	Use:   "timer",
	Short: "Time Tracker CLI to manage tasks",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Time Tracker CLI")
	},
}

var startCmd = &cobra.Command{
	Use:   "start [task name]",
	Short: "Start a new task",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		taskName := args[0]
		startTime := time.Now()

		task := tracker.Task{
			Name:      taskName,
			StartTime: startTime,
		}

		if err := tracker.SaveTask(task); err != nil {
			fmt.Println("Error saving task: ", err)
		} else {
			fmt.Printf("Start task %s at %s\n", taskName, startTime.Format(time.RFC1123))
		}
	},
}

var stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Stop the current task",
	Run: func(cmd *cobra.Command, args []string) {
		stopTime := time.Now()

		task, err := tracker.LoadLastTask()
		if err != nil {
			fmt.Println("Error loading task: ", err)
			return
		}

		task.EndTime = stopTime
		task.Duration = task.EndTime.Sub(task.StartTime).String()

		fmt.Printf("Stopped task at %s\n", stopTime.Format(time.RFC1123))
		fmt.Printf("Total Duration: %s\n", task.Duration)
	},
}

var clearCmd = &cobra.Command{
	Use:   "clear",
	Short: "Clear all tasks",
	Run: func(cmd *cobra.Command, args []string) {
		err := tracker.ClearTasks()
		if err != nil {
			fmt.Println("Error clearing tasks.")
		} else {
			fmt.Println("Cleared all tasks.")
		}
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
	rootCmd.AddCommand(stopCmd)
	rootCmd.AddCommand(clearCmd)
}

func main() {
	rootCmd.Execute()
}
