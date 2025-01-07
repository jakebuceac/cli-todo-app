/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"cli-todo-app/data"
	"fmt"
	"log"
	"os"
	"text/tabwriter"
	"time"

	"github.com/mergestat/timediff"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists all of the tasks in your todo list",
	Run:   listCommand,
}

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.Flags().BoolP("all", "a", false, "Show all tasks")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func listCommand(cmd *cobra.Command, args []string) {
	showAllTasks, _ := cmd.Flags().GetBool("all")
	tabWriter := new(tabwriter.Writer)
	models := data.Models{}
	tasks, err := models.Task.Index()

	if err != nil {
		log.Println("Failed to set tasks:", err)
	}

	tabWriter.Init(os.Stdout, 0, 8, 0, '\t', 0)

	defer tabWriter.Flush()

	if showAllTasks {
		printAllTasks(tabWriter, tasks)
	} else {
		printCurrentTasks(tabWriter, tasks)
	}
}

func printCurrentTasks(tabWriter *tabwriter.Writer, tasks []data.Task) {
	fmt.Fprintln(tabWriter, "ID\tTask\tCreated")

	for _, task := range tasks {
		if !task.Completed {
			timeDifference, err := calculateTimeDifference(task.Created)

			if err != nil {
				log.Println("Failed to convert tasks 'created' property from string to time:", err)
			}

			fmt.Fprintf(tabWriter, "%d\t%s\t%s\n", task.ID, task.Name, timeDifference)
		}
	}
}

func printAllTasks(tabWriter *tabwriter.Writer, tasks []data.Task) {
	fmt.Fprintln(tabWriter, "ID\tTask\tCreated\tDone")

	for _, task := range tasks {
		timeDifference, err := calculateTimeDifference(task.Created)

		if err != nil {
			log.Println("Failed to convert tasks 'created' property from string to time:", err)
		}

		fmt.Fprintf(tabWriter, "%d\t%s\t%s\t%t\n", task.ID, task.Name, timeDifference, task.Completed)
	}
}

func calculateTimeDifference(createdAt string) (string, error) {
	time, err := time.Parse("2006-01-02T15:04:05-07:00", createdAt)

	if err != nil {
		return "", err
	}

	return timediff.TimeDiff(time), nil
}
