package main

import (
	"fmt"
	"reflect"
)

func main() {

	type Version struct {
		ID    string
		Stage string
	}
	type test struct {
		A       bool
		B       bool
		C       bool
		version Version
	}

	type test2 struct {
		A bool
		B bool
		C bool
	}
	// q := new(test)

	// q := &test{
	// 	false,
	// 	true,
	// 	false,
	// 	Version{
	// 		"1",
	// 		"OK",
	// 	},
	// }

	q := test{}
	isVersioned(q)

	fmt.Println("***********")
	fmt.Println(isVersionedTwo(q))
	q2 := test2{}
	fmt.Println(isVersionedTwo(q2))

	// metaValue := reflect.ValueOf(v).Elem()

	// field := metaValue.FieldByName("version")
	// if field != (reflect.Value{}) {
	// 	log.Printf("Field %s not exist in struct", "version")
	// }
}

func isVersioned(q interface{}) bool {
	v := reflect.ValueOf(q)
	fmt.Println(v)
	fmt.Println(reflect.ValueOf(q).Kind())
	fmt.Println(reflect.TypeOf(q))
	fmt.Println("------------")
	for i := 0; i < v.NumField(); i++ {
		fmt.Printf("Field:%d type:%T value:%v\n", i, v.Field(i), v.Field(i))
		fmt.Println(v.Field(i).Kind())
		fmt.Println(v.Field(i).Type())
		fmt.Println("------------")
	}

	return false
}

func isVersionedTwo(q interface{}) bool {
	v := reflect.ValueOf(q)
	if v.Kind() == reflect.Struct {
		for i := 0; i < v.NumField(); i++ {
			if v.Field(i).Type().String() == "main.Version" {
				return true
			}
		}
	}

	return false
}
