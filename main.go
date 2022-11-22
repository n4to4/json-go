package main

import (
	"encoding/json"
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
	*dst = make(map[string]any)
	return nil
}
