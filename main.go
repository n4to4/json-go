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

		var bird map[string]any
		if err := json.Unmarshal([]byte(j), &bird); err != nil {
			fmt.Printf("error: %s", err)
			continue
		}

		fmt.Printf("%#v\n", bird)
		fmt.Printf("%#v\n", bird["birdType"])
		fmt.Printf("%#v\n", bird["birdtype"])
	}
}
