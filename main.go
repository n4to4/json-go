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
	um := Unmarshaler{src: src, cur: 0, dst: dst}
	if err := um.UnmarshalObject(); err != nil {
		return err
	}
	return nil
}

type Unmarshaler struct {
	src string
	cur int
	dst *map[string]any
}

func (u *Unmarshaler) currentChar() (byte, bool) {
	if u.cur < len(u.src) {
		return u.src[u.cur], true
	} else {
		return ' ', false
	}
}

func (u *Unmarshaler) next() {
	u.cur += 1
}

func (u *Unmarshaler) nextN(c int) {
	u.cur += c
}

func (u *Unmarshaler) expect(w string) bool {
	len := len(w)
	if u.src[u.cur:u.cur+len] == w {
		u.nextN(len)
		return true
	} else {
		return false
	}
}

func (u *Unmarshaler) UnmarshalObject() error {
	c, ok := u.currentChar()
	if !ok {
		return errors.New("unexpected EOF")
	}
	if c != '{' {
		return errors.New("not supported")
	}
	u.next()

	obj, err := u.parseObject()
	if err != nil {
		return err
	}
	*u.dst = obj
	return nil
}

func (u *Unmarshaler) parseObject() (map[string]any, error) {
	m := make(map[string]any)
	var key string
	seenKey := false
	for {
		c, ok := u.currentChar()
		if !ok {
			return nil, errors.New("unexpected EOF")
		}
		switch c {
		case '}':
			return m, nil
		case '"':
			u.next()
			s, err := u.parseString()
			// fmt.Printf("string %q, %d\n", s, u.cur)
			if err != nil {
				return nil, err
			}
			if seenKey {
				m[key] = s
				seenKey = false
			} else {
				key = s
				seenKey = true
			}
		case ':':
			if !seenKey {
				return nil, errors.New("found `:` without key")
			}
			u.next()
		case ',':
			if seenKey {
				return nil, errors.New("found `,` while expecting value")
			}
			u.next()
		case ' ', '\t':
			u.next()
		default:
			if !seenKey {
				return nil, errors.New("invalid syntax")
			}
			// fmt.Printf("next: %q\n", u.src[u.cur:])
			if u.expect("true") {
				m[key] = true
				seenKey = false
				continue
			} else if u.expect("false") {
				m[key] = false
				seenKey = false
				continue
			} else if u.expect("null") {
				m[key] = nil
				seenKey = false
				continue
			} else if '0' <= c && c <= '9' {
				n, err := u.parseNumber()
				if err != nil {
					return nil, fmt.Errorf("failed to parse number")
				}
				m[key] = n
				seenKey = false
				continue
			}

			return nil, errors.New("invalid syntax")
		}
	}
}

func (u *Unmarshaler) parseString() (string, error) {
	// fmt.Printf("parseString: %q\n", u.src)

	var str string
	for {
		c, ok := u.currentChar()
		if !ok {
			return "", errors.New("unexpected EOF")
		}
		switch c {
		case '"':
			u.next()
			return str, nil
		default:
			u.next()
			str = str + string(c)
		}
	}
}

func (u *Unmarshaler) parseNumber() (float64, error) {
	return 42.0, nil
}
