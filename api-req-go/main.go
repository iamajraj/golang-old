package main

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"

	// "io/ioutil"
	// "net/http"

	"reflect"
)

type Favorite struct{
	Language string `json:"language"`
}

type User struct{
	Name string `json:"name"`
	Age int	`json:"age"`
}

func main(){

	u := User{
		Name: "Raj",
		Age: 16,
	}

	// d := reflect.TypeOf(u)

	// for i:=0; i < d.NumField(); i++{
	// 	field := d.Field(i)
	// 	ftag := field.Tag.Get("json")

	// 	fmt.Println(ftag)

	// }

	// resp, err := http.Get("https://jsonplaceholder.typicode.com/users")
	// if err != nil{
	// 	fmt.Println("Error")
	// }
	// d, err := ioutil.ReadAll(resp.Body)
	// if err != nil{
	// 	fmt.Println("Error reading the file")
	// }
	// fmt.Println(string(d))

	res, err := JSONEncode(u)

	if err != nil{
		fmt.Printf("Error: %v", err)
	}

	fmt.Println(string(res))
}

func JSONEncode(v interface{})([]byte, error){
	refObjVal := reflect.ValueOf(v)
	refObjType := reflect.TypeOf(v)

	buf := bytes.Buffer{}

	if refObjVal.Kind() != reflect.Struct{
		return buf.Bytes(), fmt.Errorf("val of kind %s is not supported", refObjVal.Kind())
	}

	buf.WriteString("{")
	pairs := []string{}

	for i := 0; i<refObjVal.NumField(); i++{
		fieldRefObj := refObjVal.Field(i)
		fieldRefObjType := refObjType.Field(i)

		tag := fieldRefObjType.Tag.Get("json")

		switch fieldRefObj.Kind() {
		case reflect.String:
			strVal := fieldRefObj.Interface().(string)
			pairs = append(pairs, `"`+tag+`":"`+strVal+`"`)
		case reflect.Int:
            intVal := fieldRefObj.Interface().(int)
            pairs = append(pairs, `"`+tag+`":`+strconv.Itoa(intVal))
		default:
			return buf.Bytes(), fmt.Errorf("struct field with name %s and kind %s is not supported", fieldRefObjType.Name, fieldRefObj.Kind())
		}
	}
	buf.WriteString(strings.Join(pairs, ","))
	buf.WriteString("}")

	return buf.Bytes(), nil

}







