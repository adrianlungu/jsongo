package jsongo

import (
	"strings"
	"regexp"
	"strconv"
)

const separator = "."

// If the item is not found, this will be returned
var NotFound string = ""

func Get(query string, json interface{}) string {

	var element string

	// Get the first separator if it exists (i.e. object.firstItem.subItem)
	index := strings.Index(query, separator)

	// If there is no separator, get the value of the element
	if index == -1 {
		return getValue(query, json)
	}

	element = query[:index]

	json = getObject(element, json)
	if json == nil {
		return NotFound
	}

	return Get(query[index + 1:], json)

}

func getObject(element string, json interface{}) interface{} {

	// Check if the element refers to an array item (i.e. object.firstItem[1])
	if strings.Contains(element, "[") {
		return getArrayObject(element, json)
	}

	j, ok := json.(map[string]interface{})
	if !ok {
		return nil
	}

	return j[element] //json

}

func getValue(element string, json interface{}) string {

	// Check if the element is an array
	if strings.Contains(element, "[") {
		return getArrayObject(element, json).(string)
	}

	j, ok := json.(map[string]interface{})
	if !ok {
		return NotFound
	}

	value := j[element]

	// Check what kind of element we have (string, number) and return it
	switch v := value.(type) {
	case string:
		return v

	case float64:
		return strconv.FormatFloat(v, 'f', -1, 64)

	default:
		return NotFound
	}

	return NotFound

}

func getArrayObject(element string, json interface{}) interface{} {

	// Get the name of the array
	arrayName := element[:strings.Index(element, "[")]

	regex, err := regexp.Compile(`\[(.*?)\]`)
	if err != nil {
		panic(err)
	}

	// Get the index as string
	arrayIndexString := regex.FindStringSubmatch(element)

	// Convert the index from string to integer
	arrayIndex, err := strconv.Atoi(arrayIndexString[1])
	if err != nil {
		panic(err)
	}

	// Check if the JSON interface is a map
	j, ok := json.(map[string]interface{})
	if !ok {
		return NotFound
	}

	// Get the array from the JSON if it exists
	jsonArray, ok := j[arrayName].([]interface{})
	if !ok {
		return NotFound
	}

	// Check if the specified index is within range
	if len(jsonArray) <= arrayIndex {
		return NotFound
	}

	return jsonArray[arrayIndex]

}