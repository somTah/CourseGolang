package main

import "os"

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	data1 := []byte("Hi\n borntoDev")
	err := os.WriteFile("data.txt", data1, 0644)
	check(err)
	f, err := os.Create("employee")
	check(err)

	defer f.Close()

	data2 := []byte("Sira\n Manee")
	os.WriteFile("employee", data2, 0644)

}
