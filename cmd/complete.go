/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"cli-todo-app/data"
	"cli-todo-app/helpers"
	"fmt"
	"log"
	"os"
	"strconv"
	"text/tabwriter"

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
	task, err := models.Task.Show(taskId)

	if err != nil {
		log.Println("Could not find task:", err)

		return
	}

	task.Completed = true
	task, err = models.Task.Update(task)

	if err != nil {
		log.Println("Failed to update task:", err)

		return
	}

	printUpdatedTask(task)
}

func printUpdatedTask(task data.Task) {
	tabWriter := new(tabwriter.Writer)
	tabWriter.Init(os.Stdout, 0, 8, 0, '\t', 0)

	defer tabWriter.Flush()

	fmt.Fprintln(tabWriter, "ID\tTask\tCreated\tDone")

	timeDifference, err := helpers.CalculateTimeDifference(task.Created)

	if err != nil {
		log.Println("Failed to convert tasks 'created' property from string to time:", err)

		return
	}

	fmt.Fprintf(tabWriter, "%d\t%s\t%s\t%t\n", task.ID, task.Name, timeDifference, task.Completed)
}
