package data

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

func ImportCSV(path string) ([]byte, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func ParseCSV(data []byte) {
	reader := csv.NewReader(bytes.NewReader(data))
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("Error with the processing of the records.")
			break
		}
		fmt.Println(record)
	}
}

func ExportCSV(records [][]string) {
	file, err := os.Create("")
	if err != nil {
		return
	}
	defer file.Close()

	writer := csv.NewWriter(file)

	for i := 0; i < len(records); i++ {
		err = writer.Write(records[i])
		if err != nil {
			fmt.Println("Error writing record to CSV:", err)
		}
	}
	writer.Flush()
}
