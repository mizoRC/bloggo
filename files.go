package main

/*
import (
	"encoding/json"
	"io/ioutil"
	//"log"
	"os"
	"strings"
	"unicode"
)

func SpaceStringsBuilder(str string) string {
	var b strings.Builder
	b.Grow(len(str))
	for _, ch := range str {
		if !unicode.IsSpace(ch) {
			b.WriteRune(ch)
		}
	}
	return b.String()
}

func readJSONToken(fileName string, filter func(map[string]interface{}) bool) []map[string]interface{} {
	file, _ := os.Open(fileName)
	defer file.Close()

	decoder := json.NewDecoder(file)

	filteredData := []map[string]interface{}{}

	// Read the array open bracket
	decoder.Token()

	data := map[string]interface{}{}
	for decoder.More() {
		decoder.Decode(&data)

		if filter(data) {
			filteredData = append(filteredData, data)
		}
	}

	return filteredData
}

func writeJSON(post Post) {
	titleSpacesSplitted := SpaceStringsBuilder(post.Title)
	filePath := titleSpacesSplitted + ".json"
	//log.Printf("FILEPATH => %v", filePath)

	jsonString, _ := json.Marshal(post)
	ioutil.WriteFile(filePath, jsonString, os.ModePerm)
}
*/
