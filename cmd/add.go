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
	"text/tabwriter"
	"time"

	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new task to the todo list",
	Args:  cobra.ExactArgs(1),
	Run:   addCommand,
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func addCommand(cmd *cobra.Command, args []string) {
	taskDescription := args[0]
	models := data.Models{}
	payload := data.Task{
		Name:      taskDescription,
		Created:   time.Now().Format("2006-01-02T15:04:05-07:00"),
		Completed: false,
	}
	taskId, err := models.Task.Store(payload)

	if err != nil {
		log.Println("Failed to save new task:", err)

		return
	}

	task, err := models.Task.Show(taskId)

	if err != nil {
		log.Println("Failed to get new task:", err)

		return
	}

	printNewTask(task)
}

func printNewTask(task *data.Task) {
	tabWriter := new(tabwriter.Writer)
	tabWriter.Init(os.Stdout, 0, 8, 0, '\t', 0)

	defer tabWriter.Flush()

	fmt.Fprintln(tabWriter, "ID\tTask\tCreated")

	timeDifference, err := helpers.CalculateTimeDifference(task.Created)

	if err != nil {
		log.Println("Failed to convert tasks 'created' property from string to time:", err)

		return
	}

	fmt.Fprintf(tabWriter, "%d\t%s\t%s\n", task.ID, task.Name, timeDifference)
}
