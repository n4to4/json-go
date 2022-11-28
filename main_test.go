package main

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestParseObject(t *testing.T) {
	tests := []struct {
		name string
		json string
		want map[string]any
	}{
		{
			name: "an empty object",
			json: `{}`,
			want: map[string]any{},
		},
		{
			name: "key value",
			json: `{"name":"taro","pet":"lambda"}`,
			want: map[string]any{"name": "taro", "pet": "lambda"},
		},
		{
			name: "skip whitespace",
			json: `{ "name" : "taro"  ,  "pet"  :  "lambda"  }`,
			want: map[string]any{"name": "taro", "pet": "lambda"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var obj map[string]any
			if err := json.Unmarshal([]byte(tt.json), &obj); err != nil {
				t.Errorf("Got err %s", err)
			}
			if !reflect.DeepEqual(tt.want, obj) {
				t.Errorf("json.Unmarshal: want %#v, got %#v", tt.want, obj)
			}

			var obj2 map[string]any
			if err := Unmarshal(tt.json, &obj2); err != nil {
				t.Fail()
			}
			if !reflect.DeepEqual(tt.want, obj2) {
				t.Errorf("Unmarshal: want %#v, got %#v", tt.want, obj2)
			}
		})
	}
}
