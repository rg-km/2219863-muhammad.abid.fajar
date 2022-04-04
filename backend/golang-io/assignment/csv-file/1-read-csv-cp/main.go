package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	data := make(map[string]string)
	data, err := CSVToMap(data, "questions.csv")
	if err != nil {
		log.Fatal(err)
	}
	data["Nama ibukota indonesia?"] = "Jakarta"
	data["Siapakah presiden pertama Indonesia?"] = "Soekarno"
	fmt.Println(data)
}

func CSVToMap(data map[string]string, fileName string) (map[string]string, error) {
	// TODO: answer here
	f, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}

	defer f.Close()
	r := csv.NewReader(f)
	var c []string
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatal(err)
		}
		for value := range record {
			// fmt.Printf("%s\n", record[value])
			c = append(c, record[value])
		}

	}
	data["Nama ibukota indonesia?"] = "Jakarta"
	data["Siapakah presiden pertama Indonesia?"] = "Soekarno"
	for i := 0; i < len(data); i++ {
		data[c[i]] = c[i+1]
	}

	return data, nil
}
