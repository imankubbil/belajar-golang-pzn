package main

import (
	"fmt"
	"reflect"
)

type Sample struct {
	Name string `required:"true" max:"10"`
}

type Person struct {
	Name    string `required:"true" max:"10"`
	Address string `required:"true" max:"10"`
	Email   string `required:"true" max:"10"`
}

func readField(value any) {
	var valueType reflect.Type = reflect.TypeOf(value)
	fmt.Println("Type Name:", valueType)

	for i := 0; i < valueType.NumField(); i++ {
		var structField reflect.StructField = valueType.Field(i)
		fmt.Println("Field Name:", structField.Name)
		fmt.Println("Field Type:", structField.Type)
		fmt.Println("Tag Required :", structField.Tag.Get("required"))
		fmt.Println("Tag Max :", structField.Tag.Get("max"))
	}
}

func isValid(data interface{}) (result bool) {
	t := reflect.TypeOf(data)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if field.Tag.Get("required") == "true" {
			data := reflect.ValueOf(data).Field(i).Interface()
			result = data != ""
			if result == false {
				return result
			}
		}
	}

	return true
}

func main() {
	readField(Sample{"John"})
	readField(Person{"John", "", ""})

	person := Person{
		Name:    "John",
		Address: "123 Main Street",
		Email:   "",
	}

	fmt.Println(isValid(person))
}
