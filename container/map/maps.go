package main

import (
	"fmt"
)

func main() {
	m := map[string]string{
		"Name":   "Jacky",
		"Nation": "China",
		"Hobby":  "Swimming",
	}
	m2 := make(map[string]int)
	var m3 map[string]int

	fmt.Println(m, m2, m3)

	for k, v := range m {
		fmt.Println(k, v)
	}

	if name, ok := m["Name"]; ok {
		fmt.Println(name)
	} else {
		fmt.Println("key error")
	}

	delete(m, "Name")
	if name, ok := m["Name"]; ok {
		fmt.Println(name)
	} else {
		fmt.Println("key error")
	}
}
