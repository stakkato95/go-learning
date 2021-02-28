package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type User struct {
	ID               int    `json:"user_id"`
	Username         string `json:"name"`
	Phone            string `json:"phoneNumber"`
	NotImportantData string `json:"-"` //should not be serialized / deserialized at all
}

func main() {
	//marshalUnmarshal()
	// workWithUnstructuredJson()
	reflection()
}

func marshalUnmarshal() {
	jsonStr := `{"user_id": 42, "name": "alex", "phoneNumber": "+100500"}`

	data := []byte(jsonStr)

	u := &User{}
	json.Unmarshal(data, u)
	fmt.Printf("%#v\n", u)

	u.Phone = "+100900"
	result, err := json.Marshal(u)
	if err != nil {
		fmt.Println("Error when marshaling")
		return
	}

	fmt.Println(string(result))
}

func workWithUnstructuredJson() {
	jsonStr := `[
		{"id": 17, "username": "myname", "phone": 12345},
		{"id": "17", "address": "none", "company": "muCompany"}
	]`

	data := []byte(jsonStr)
	var usersArray interface{}
	json.Unmarshal(data, &usersArray)
	fmt.Printf("unpacked %#v\n", usersArray)

	array := usersArray.([]interface{})
	fmt.Println(array)
	user1 := array[0].(map[string]interface{})
	fmt.Println(user1)

	for key, val := range user1 {
		switch val.(type) {
		case string:
			fmt.Println("string(", key, ")->string(", val, ")")
		case float64:
			fmt.Println("string(", key, ")->float64(", val, ")")
		default:
			fmt.Println("string(", key, ")->???(", val, ")")
		}
	}
}

func reflection() {
	u := &User{
		ID:               12,
		Username:         "alex",
		Phone:            "+12345",
		NotImportantData: "some data",
	}
	printReflect(u)
}

func printReflect(iface interface{}) {
	val := reflect.ValueOf(iface).Elem()

	fields := val.NumField()
	fmt.Printf("%T has %d fields\n", iface, fields)
	for i := 0; i < fields; i++ {
		valueField := val.Field(i)
		typeField := val.Type().Field(i)
		fmt.Printf("('%s' '%s' %s), %v\n", typeField.Name, typeField.Type.Kind(), typeField.Tag, valueField)
	}
}
