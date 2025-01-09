/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"cli-todo-app/data"
	"fmt"
	"log"
	"os"
	"strconv"
	"text/tabwriter"
	"time"

	"github.com/mergestat/timediff"
	"github.com/spf13/cobra"
)

// completeCmd represents the complete command
var completeCmd = &cobra.Command{
	Use:   "complete",
	Short: "Set a task as being completed",
	Args:  cobra.ExactArgs(1),
	Run:   completeCommand,
}

func init() {
	rootCmd.AddCommand(completeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// completeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// completeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func completeCommand(cmd *cobra.Command, args []string) {
	taskId, err := strconv.Atoi(args[0])

	if err != nil {
		log.Println("Could not convert task ID to int:", err)

		return
	}

	models := data.Models{}
	task, err := models.Task.Show(int64(taskId))

	if err != nil {
		log.Println("Failed to get task:", err)

		return
	}

	task.Completed = true
	err = task.Update()

	if err != nil {
		log.Println("Failed to update task:", err)

		return
	}

	printUpdatedTask(task)
}

func printUpdatedTask(task *data.Task) {
	tabWriter := new(tabwriter.Writer)
	tabWriter.Init(os.Stdout, 0, 8, 0, '\t', 0)

	defer tabWriter.Flush()

	fmt.Fprintln(tabWriter, "ID\tTask\tCreated\tDone")

	time, err := time.Parse("2006-01-02 15:04:05", task.Created)

	if err != nil {
		log.Println("Failed to format datetime string:", err)
	}

	timeDifference := timediff.TimeDiff(time)

	fmt.Fprintf(tabWriter, "%d\t%s\t%s\t%t\n", task.ID, task.Name, timeDifference, task.Completed)
}
