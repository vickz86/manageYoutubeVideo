package dataloader

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	mainFile "github.com/vickz86/manageYoutubeVideo"
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


// return a slice of slice of string, make sure same number element in each slice of slice, data is name of file in same folder
func SplitEachLine(data string)([][]string){
	
	// declare the variable
	var sliceOfSliceString [][]string
	
	//load data
	rawSlice,err := LoadFileLines(data)
	if err != nil {
		log.Fatal("error loading data!!")
	}

	for index,value := range rawSlice{
		// split each line based on ";"
		elements := strings.Split(value,";")
		
		// create a slice of string based on slice "elements" and append it
		sliceOfSliceString = append(sliceOfSliceString,elements)


		//check if same number of element in each slice
		if index==0{
			continue
		} else {
			if len(elements) != len(sliceOfSliceString[index-1]){
				fmt.Printf("ERROR , not the same number of element at index %v",index)
			}
		}

	}

	return sliceOfSliceString
}


func CreateSliceYoutube(data string)([]mainFile.YoutubeStruct){

	//create [][]string from the data file
	sliceOfSlice :=SplitEachLine(data)


	//create the slice of struct 
	var youtubeSlice []mainFile.YoutubeStruct

	// loop through the first slice
	for _,value := range sliceOfSlice{
		
		//
		fmt.Printf("the value 3 is of type %T",value[3])	



		// define a new struct and add value
		// newYoutube := mainFile.YoutubeStruct{value[0],value[1],value[2],value[3]}
		
	}

	return youtubeSlice






	
}
