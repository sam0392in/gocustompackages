package jsonreader

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func ReadJSON(dataFile string) interface{} {

	jsonFile, err := os.Open(dataFile)
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened ", dataFile)
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	// use standard interfaces{} in order to read in any JSON data:
	var result map[string]interface{}

	json.Unmarshal([]byte(byteValue), &result)

	return result
}
