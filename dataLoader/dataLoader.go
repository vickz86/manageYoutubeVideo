package dataloader

import (
	"bufio"
	"log"
	"os"
	"strings"
)

// LoadFileLines reads a file and returns its lines as a slice of strings
func LoadFileLines(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	
	return lines, nil
}


// return a slice of slice of string, make sure same number element in each slice of slice
func SplitEachLine(data string)([][]string){
	
	// declare the variable
	var sliceOfSliceString [][]string
	
	//load data
	rawSlice,err := LoadFileLines(data)
	if err != nil {
		log.Fatal("error loading data!!")
	}

	for _,value := range rawSlice{
		elements := strings.Split(value,";")
		sliceOfSliceString = append(sliceOfSliceString,elements)
	}

	return sliceOfSliceString
}