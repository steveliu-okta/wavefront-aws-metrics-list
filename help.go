package main

import (
	"fmt"
	"os"
)

func contains(strArr []string, check string) bool {
	for _, s := range strArr {
		if s == check {
			return true
		}
	}

	return false
}

func writeToJSONFile(data []string, filename string) error {
	file, err := os.OpenFile(fmt.Sprintf("output_data/%s_metrics.txt", filename), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return err
	}
	defer file.Close()

	for _, str := range data {
		_, err = file.WriteString(str + "\n")
		if err != nil {
			return err
		}
	}

	return nil
}
