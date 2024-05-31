package command

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

const (
	respEndline = "\r\n"
)

type Serializer interface {
	Serialize(obj any) (string, error)
	Deserialize(data string) (any, error)
}

type RESPSerializer struct {
	commandParser CommandParser
}

func NewRESPSerializer(commandParser CommandParser) *RESPSerializer {
	return &RESPSerializer{
		commandParser: commandParser,
	}
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
	res, _, err := rs.deserialize(data, 0)
	return res, err
}

func (rs *RESPSerializer) deserialize(data string, index int) (any, int, error) {
	if len(data) == 0 {
		return nil, 0, errors.New("data is empty")
	}

	var charType = data[index]
	switch charType {
	case '$':
		var extractedNumber, index, err = rs.extractNumberFromLine(data, index+1)
		if err != nil {
			return nil, 0, err
		}

		var strLen, _ = strconv.Atoi(extractedNumber)

		var strDeserialized strings.Builder
		for i := index; i < index+strLen; i++ {
			strDeserialized.WriteByte(data[i])
		}
		indexForEndline := strings.Index(data[index:], respEndline)
		if indexForEndline == -1 {
			return nil, 0, err
		}
		index = index + indexForEndline + len(respEndline)

		var returnValue any = strDeserialized.String()

		cmd, err := rs.commandParser.Parse(strDeserialized.String())
		if err == nil {
			returnValue = cmd
		}

		return returnValue, index, nil
	case '*':
		var extractedNumber, index, err = rs.extractNumberFromLine(data, index+1)
		if err != nil {
			return nil, 0, err
		}

		var elemInArray, _ = strconv.Atoi(extractedNumber)
		var elements = make([]any, 0, elemInArray)

		for range elemInArray {
			result, newIndex, err := rs.deserialize(data, index)
			if err != nil {
				return nil, 0, err
			}

			index = newIndex
			elements = append(elements, result)
		}

		return elements, 0, nil

	default:
		return nil, 0, errors.New("invalid char")
	}
}

func (rs *RESPSerializer) extractNumberFromLine(input string, start int) (string, int, error) {
	var indexPos = strings.Index(input[start:], respEndline)
	if indexPos == -1 {
		return "", start, errors.New("invalid response")
	}

	var end = start + indexPos

	return input[start:end], end + len(respEndline), nil
}
