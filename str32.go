package main

import "encoding/json"

type str32 int32

var str32List = []string{""}
var str32Lookup = map[string]str32{"": 0}

func getStr32(s string) str32 {
	if i, ok := str32Lookup[s]; ok {
		return i
	}
	i := str32(len(str32List))
	str32List = append(str32List, s)
	str32Lookup[s] = i
	return i
}

func (s str32) String() string {
	return str32List[s]
}

func (s str32) MarshalJSON() ([]byte, error) {
	return json.Marshal(str32List[s])
}
