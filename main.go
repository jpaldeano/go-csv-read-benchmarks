package main

import (
	"encoding/csv"
	"go-file-reader/junior"
	"go-file-reader/senior"
	"os"
	"sync"
	"time"
)

func externalCall(s []string) {
	time.Sleep(250 * time.Millisecond)
	//fmt.Println(s)
}

func main() {
	f, _ := os.Open("example1k.csv")
	defer f.Close()
	csvReader := csv.NewReader(f)

	junior.JuniorLevel(csvReader, externalCall)

	var wg sync.WaitGroup
	senior.SeniorLevel(&wg, csvReader, externalCall)
}
