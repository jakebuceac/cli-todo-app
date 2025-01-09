/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"cli-todo-app/data"
	"log"
	"strconv"

	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Removes a task for the todo list by it's id",
	Args:  cobra.ExactArgs(1),
	Run:   deleteCommand,
}

func init() {
	rootCmd.AddCommand(deleteCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func deleteCommand(cmd *cobra.Command, args []string) {
	taskId, err := strconv.Atoi(args[0])

	if err != nil {
		log.Println("Could not convert task ID to int:", err)

		return
	}

	models := data.Models{}
	task, err := models.Task.Show(int64(taskId))

	if err != nil {
		log.Println("Could not find task:", err)

		return
	}

	err = task.Delete()

	if err != nil {
		log.Println("Could not delete task:", err)

		return
	}
}
