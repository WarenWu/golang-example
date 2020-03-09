package encode

import (
	"encoding/json"
	"fmt"
)
type FamilyMember struct {
	Name    string
	Age     int
	Parents []string
}

func ExJson()  {
	m := map[string]interface{} {
		"Name": "Wednesday",
		"Age":  6,
		"Parents": []interface{} {
			"Gomez",
			"Morticia",
		},
	}
    v, err:=json.Marshal(&m)
    if err != nil {
        fmt.Println(err.Error())
	}else{
		fmt.Println(v)
	}
}