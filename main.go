package main

import (
	"encoding/json"
	"errors"
	"fmt"
)

func main() {
	birdJson := []string{
		`{}`,
		`{"birdType": "pigeon","description": "likes to perch on rocks"}`,
		`{"name":"string","age":42,"array":[1,true,null],"empty":null}`,
	}

	for _, j := range birdJson {
		fmt.Printf("\ninput: %q\n", j)

		var obj map[string]any
		if err := json.Unmarshal([]byte(j), &obj); err != nil {
			fmt.Printf("error: %s", err)
			continue
		}

		fmt.Printf("%#v\n", obj)
		fmt.Printf("%#v\n", obj["birdType"])
		fmt.Printf("%#v\n", obj["birdtype"])
	}
}

func Unmarshal(src string, dst *map[string]any) error {
	if src[0] != '{' {
		return errors.New("not supported")
	}

	obj, err := parseObject(src[1:])
	if err != nil {
		return err
	}
	*dst = obj

	return nil
}

func parseObject(src string) (map[string]any, error) {
	m := make(map[string]any)
	if src[0] == '}' {
		return m, nil
	}

	m = map[string]any{"name": "taro"}
	return m, nil
}
