package dataloader

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/vickz86/manageYoutubeVideo/model"
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


func CreateSliceYoutube(data string)([]model.YoutubeStruct){

	//create [][]string from the data file
	sliceOfSlice :=SplitEachLine(data)


	//create the slice of struct 
	var youtubeSlice []model.YoutubeStruct

	// loop through the first slice
	for nb,value := range sliceOfSlice{
		
		//convert the 3 value to an int
		intStatus,err := strconv.Atoi(value[3])
		if err!=nil{
			fmt.Println("Error converting string to int:", err)
		}	

		// define a new struct and add value
		strucYoutube := model.YoutubeStruct{
			//assigne the index field based on the nb in range
			Index: nb,
			Url: value[0],
			Channel: value[1],
			Description: value[2],
			Status: intStatus,
			
		}

		youtubeSlice = append(youtubeSlice, strucYoutube)
		
	}

	return youtubeSlice






	
}
