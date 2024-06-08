package serializer

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

const (
	respEndline = "\r\n"
)
const (
	Array = '*'
	Bulk  = '$'
)

type Serializer interface {
	Serialize(obj any) (string, error)
	Deserialize(data string) (any, error)
}

type RESPSerializer struct {
}

func NewRESPSerializer() *RESPSerializer {
	return &RESPSerializer{}
}

func (rs *RESPSerializer) Serialize(obj any) (string, error) {
	switch v := obj.(type) {
	case []any:
		fmt.Println("The value is of type []any:", v)
	default:
		fmt.Println("Unknown type")
	}

	return "", nil
}

func (rs *RESPSerializer) Deserialize(data string) (any, error) {
	res, _, err := rs.deserialize(data)
	return res, err
}

func (rs *RESPSerializer) deserialize(data string) (any, string, error) {
	if len(data) == 0 {
		return nil, data, errors.New("data is empty")
	}

	var stringLeftToParse string

	var charType = data[0]
	switch charType {
	case Bulk:
		stringLeftToParse = data[1:]
		var extractedNumber, err = rs.extractNumberFromLine(stringLeftToParse)
		if err != nil {
			return nil, data, err
		}

		stringLeftToParse, ok := rs.removeLineFromString(stringLeftToParse)
		if !ok {
			return nil, data, errors.New("endline not found in " + data)
		}

		var strLen, _ = strconv.Atoi(extractedNumber)

		var strDeserialized strings.Builder
		for i := 0; i < strLen; i++ {
			strDeserialized.WriteByte(stringLeftToParse[i])
		}

		stringLeftToParse, ok = rs.removeLineFromString(stringLeftToParse)
		if !ok {
			return nil, data, errors.New("endline not found in " + data)
		}

		return strDeserialized.String(), stringLeftToParse, nil
	case Array:
		stringLeftToParse = data[1:]
		var extractedNumber, err = rs.extractNumberFromLine(stringLeftToParse)
		if err != nil {
			return nil, data, err
		}

		stringLeftToParse, ok := rs.removeLineFromString(stringLeftToParse)
		if !ok {
			return nil, data, errors.New("endline not found in " + data)
		}

		elemInArray, err := strconv.Atoi(extractedNumber)
		if err != nil {
			return nil, data, err
		}

		var elements = make([]any, 0, elemInArray)
		for range elemInArray {
			result, stringLeftIteration, err := rs.deserialize(stringLeftToParse)
			if err != nil {
				return nil, data, err
			}
			stringLeftToParse = stringLeftIteration

			elements = append(elements, result)
		}

		return elements, stringLeftToParse, nil

	default:
		return nil, data, errors.New("invalid char")
	}
}

func (rs *RESPSerializer) removeLineFromString(data string) (string, bool) {
	if strings.Index(data, respEndline) == -1 {
		return data, false
	}
	indexForSliceArrayCount := strings.Index(data, respEndline) + len(respEndline)
	return data[indexForSliceArrayCount:], true
}

func (rs *RESPSerializer) extractNumberFromLine(input string) (string, error) {
	var end = strings.Index(input, respEndline)
	if end == -1 {
		return "", errors.New("invalid response")
	}

	return input[:end], nil
}
