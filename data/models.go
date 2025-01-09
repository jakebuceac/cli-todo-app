package data

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
)

type Models struct {
	Task Task
}

type Task struct {
	ID        int
	Name      string
	Created   string
	Completed bool
}

func (t *Task) Index() ([]*Task, error) {
	mydir, err := os.Getwd()

	if err != nil {
		log.Println("Failed to get current directory")

		return nil, err
	}

	records, err := readCsvFile(mydir + "/data/todo-list.csv")

	if err != nil {
		log.Println("Failed to read csv file")

		return nil, err
	}

	var tasks []*Task

	for index, record := range records {
		if index == 0 {
			continue
		}

		taskFinished, err := strconv.ParseBool(record[3])

		if err != nil {
			return nil, err
		}

		taskId, err := strconv.Atoi(record[0])

		if err != nil {
			return nil, err
		}

		tasks = append(tasks, &Task{
			ID:        taskId,
			Name:      record[1],
			Created:   record[2],
			Completed: taskFinished,
		})
	}

	return tasks, nil
}

func (t *Task) Store(task Task) (int64, error) {
	tasks, err := t.Index()

	if err != nil {
		log.Println("Failed to get all tasks:")

		return 0, err
	}

	// Set a unique ID for new Task
	task, err = setTaskId(task, tasks)

	if err != nil {
		log.Println("Failed to set task ID")

		return 0, err
	}

	tasks = append(tasks, &task)

	mydir, err := os.Getwd()

	if err != nil {
		log.Println("Failed to get current directory")

		return 0, err
	}

	err = writeToCsv(mydir+"/data/todo-list.csv", tasks)

	if err != nil {
		log.Println("Failed to add new task to csv")

		return 0, err
	}

	return int64(task.ID), nil
}

func (t *Task) Show(taskId int64) (*Task, error) {
	tasks, err := t.Index()

	if err != nil {
		return &Task{}, err
	}

	for _, task := range tasks {
		if taskId == int64(task.ID) {
			return task, nil
		}
	}

	return &Task{}, fmt.Errorf("could not find task with ID %d", taskId)
}

func (t *Task) Update() error {
	tasks, err := t.Index()

	if err != nil {
		log.Println("Failed to get all tasks")

		return err
	}

	for index := range tasks {
		if tasks[index].ID == t.ID {
			tasks[index].Completed = t.Completed
			break
		}
	}

	mydir, err := os.Getwd()

	if err != nil {
		log.Println("Failed to get current directory")

		return err
	}

	err = writeToCsv(mydir+"/data/todo-list.csv", tasks)

	if err != nil {
		log.Println("Failed to update task in csv")

		return err
	}

	return nil
}

func (t *Task) Delete() error {
	tasks, err := t.Index()

	if err != nil {
		log.Println("Failed to get all tasks")

		return err
	}

	for index := range tasks {
		if tasks[index].ID == t.ID {
			tasks = append(tasks[:index], tasks[index+1:]...)
			break
		}
	}

	mydir, err := os.Getwd()

	if err != nil {
		log.Println("Failed to get current directory")

		return err
	}

	err = writeToCsv(mydir+"/data/todo-list.csv", tasks)

	if err != nil {
		log.Println("Failed to delete task in csv")

		return err
	}

	return nil
}

func setTaskId(task Task, tasks []*Task) (Task, error) {
	id := 1

	if len(tasks) > 0 {
		lastId := tasks[len(tasks)-1].ID

		id = lastId + 1
	}

	task.ID = id

	return task, nil
}

func readCsvFile(filename string) ([][]string, error) {
	file, err := os.Open(filename)

	if err != nil {
		return nil, err
	}

	defer file.Close()

	csvReader := csv.NewReader(file)
	records, err := csvReader.ReadAll()

	if err != nil {
		return nil, err
	}

	return records, nil
}

func writeToCsv(filename string, tasks []*Task) error {
	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, os.ModePerm)

	if err != nil {
		return err
	}

	defer file.Close()

	csvWriter := csv.NewWriter(file)

	defer csvWriter.Flush()

	// Set headers of CSV
	_ = csvWriter.Write([]string{
		"ID",
		"Description",
		"CreatedAt",
		"IsComplete",
	})

	// Set tasks
	for _, task := range tasks {
		_ = csvWriter.Write([]string{
			strconv.Itoa(task.ID),
			task.Name,
			task.Created,
			strconv.FormatBool(task.Completed),
		})
	}

	return nil
}
