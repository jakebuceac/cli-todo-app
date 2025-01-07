package data

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"
)

type Models struct {
	Task Task
}

type Task struct {
	ID        string
	Name      string
	Created   string
	Completed bool
}

func (t *Task) Index() ([]Task, error) {
	mydir, err := os.Getwd()

	if err != nil {
		log.Println("Failed to get current directory:", err)
	}

	records, err := readCsvFile(mydir + "/data/todo-list.csv")

	if err != nil {
		log.Println("Failed to read csv file:", err)
	}

	var tasks []Task

	for index, record := range records {
		if index == 0 {
			continue
		}

		taskFinished, err := strconv.ParseBool(record[3])

		if err != nil {
			return nil, err
		}

		tasks = append(tasks, Task{
			ID:        record[0],
			Name:      record[1],
			Created:   record[2],
			Completed: taskFinished,
		})
	}

	return tasks, nil
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
