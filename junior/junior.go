package junior

import (
	"encoding/csv"
	"fmt"
	"io"
)

func JuniorLevel(csvReader *csv.Reader, externalCall func([]string)) {
	fmt.Println("running junior level")
	for {
		rStr, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("ERROR: ", err.Error())
			break
		}
		externalCall(rStr)

	}
}
