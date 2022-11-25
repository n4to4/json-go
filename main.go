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
	var key string
	haveKey := false
	var idx int
	for idx = 0; idx < len(src); idx++ {
		c := src[idx]
		switch c {
		case '}':
			return m, nil
		case '"':
			s, i, err := parseString(src[idx+1:])
			fmt.Printf("case %q, %d\n", s, i)
			if err != nil {
				return nil, err
			}
			if haveKey {
				m[key] = s
				haveKey = false
			} else {
				key = s
				haveKey = true
			}
			idx += i + 1
		case ':':
		}
	}
	return m, nil
}

func parseString(src string) (string, int, error) {
	fmt.Printf("parseString: %q\n", src)

	var str string
	for i, c := range src {
		switch c {
		case '"':
			return str, i, nil
		default:
			str = str + string(c)
		}
	}
	return "", 0, fmt.Errorf("parseString: missing closing `\"`: %q", src)
}
