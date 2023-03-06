package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	// parse the given serial number
	result := parseSerialNumber("SerialNumber=018277314057")
	fmt.Println("the serial number is", result)
	return
	// Gives functionality to calculator
	reader := bufio.NewReader(os.Stdin)
	//Gives value 1
	fmt.Print("Value 1: ")
	input1, _ := reader.ReadString('\n')
	float1, err := strconv.ParseFloat(strings.TrimSpace(input1), 64)
	if err != nil {
		fmt.Println(err)
		panic("Value 1 must be a number")
	}
	//Grabs value 2
	fmt.Print("Value 2: ")
	input2, _ := reader.ReadString('\n')
	float2, err := strconv.ParseFloat(strings.TrimSpace(input2), 64)
	if err != nil {
		fmt.Println(err)
		panic("Value 2 must be a number")
	}
	// Gives Sum
	sum := float1 + float2
	sum = math.Round(sum*100) / 100
	fmt.Printf("The sum of %v and %v is %v\n\n", float1, float2, sum)
}
//Returns serial number
func parseSerialNumber(input string) string {
	result := strings.Replace(input, "SerialNumber=", "", 1)
	return result
}
