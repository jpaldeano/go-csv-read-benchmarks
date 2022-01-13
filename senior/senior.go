package senior

import (
	"encoding/csv"
	"fmt"
	"io"
	"sync"
)

func worker(records chan []string, wg *sync.WaitGroup, externalCall func([]string)) {
	for record := range records {
		externalCall(record)
		wg.Done()
	}
}
func SeniorLevel(wg *sync.WaitGroup, csvReader *csv.Reader, externalCall func([]string)) {
	lines := make(chan []string, 500)
	for i := 0; i < cap(lines); i++ {
		go worker(lines, wg, externalCall)
	}

	for {
		wg.Add(1)
		rStr, err := csvReader.Read()
		if err == io.EOF {
			wg.Done()
			break
		}
		if err != nil {
			wg.Done()
			fmt.Println("ERROR: ", err.Error())
			break
		}
		lines <- rStr
	}
	wg.Wait()
	close(lines)
}
