package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sync"
)

type Data struct {
	A int `json:"a"`
	B int `json:"b"`
}

func main() {
	if len(os.Args) != 2 {
		log.Fatal("\nUsage: go run . <json file name> | Example: go run . test.txt")
	}

	data, err := readJSONFile(os.Args[1])
	if err != nil {
		log.Fatalf("Error reading JSON file: %v", err)
	}

	sum := calculateSum(data)
	fmt.Printf("Total sum result: %d\n", sum)
}

func readJSONFile(filePath string) ([]Data, error) {
	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var data []Data
	err = json.Unmarshal(file, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func calculateSum(data []Data) int {
	goroutinesCnt := 8

	var wg sync.WaitGroup
	chunkSize := len(data) / goroutinesCnt
	ch := make(chan int)

	for i := 0; i < goroutinesCnt; i++ {
		wg.Add(1)
		go func(start int) {
			defer wg.Done()
			localSum := 0
			end := start + chunkSize
			if i == goroutinesCnt-1 {
				end = len(data)
			}
			for j := start; j < end; j++ {
				localSum += data[j].A + data[j].B
			}
			ch <- localSum
		}(i * chunkSize)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	totalSum := 0
	for sum := range ch {
		totalSum += sum
	}

	return totalSum
}
