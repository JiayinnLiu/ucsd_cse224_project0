package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"os"
	"sort"
)

func readFrom(filePath string) [][]byte {

	data, err := ioutil.ReadFile(filePath)
	checkError(err)

	var size int = len(data) / 100
	records := make([][]byte, size)

	for i := range records {
		records[i] = make([]byte, 100)
	}

	for i := 0; i < size; i++ {
		records[i] = data[i*100 : (i+1)*100]
	}

	return records
}

func writeInto(filePath string, slice [][]byte) {
	outputFile, err := os.Create(filePath)
	checkError(err)
	defer outputFile.Close()

	for _, d := range slice {
		_, err = outputFile.Write(d)
		checkError(err)
	}

	return
}

func checkError(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	if len(os.Args) != 3 {
		log.Fatalf("Usage: %v inputfile outputfile\n", os.Args[0])
	}

	inputPath := os.Args[1]
	records := readFrom(inputPath)

	// sort by key
	sort.Slice(records, func(i, j int) bool {
		return bytes.Compare(records[i][:10], records[j][:10]) < 0
	})

	outputPath := os.Args[2]
	writeInto(outputPath, records)
	log.Printf("Sorting %s to %s\n", os.Args[1], os.Args[2])
}
