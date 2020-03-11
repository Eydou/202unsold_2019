//
// EPITECH PROJECT, 2020
// 202unsold_2019
// File description:
// main
//

package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"

	functions "./functions"
)

var strFlag = flag.String("h", "-h", "Help")

func init() {
	flag.StringVar(strFlag, "", "",
		"USAGE\n\t  ./202unsold a b\n\nDESCRIPTION\n\t   a\tconstant computed from the past results\n\t   b\tconstant computed from the past results")
}
func main() {
	args := os.Args
	flag.Parse()
	string1 := flag.Arg(0)
	string2 := flag.Arg(1)
	number1, err := strconv.Atoi(string1)
	if err != nil {
		fmt.Println(err, number1)
		os.Exit(84)
	}
	number2, err := strconv.Atoi(string2)
	if err != nil {
		fmt.Println(err, number2)
		os.Exit(84)
	}
	if number1 < 50 || number2 < 50 {
		os.Exit(84)
	}
	if ok, err := functions.ErrorArgs(args); err != nil {
		fmt.Println("error:", err)
		os.Exit(84)
	} else {
		functions.Math(float64(number1), float64(number2))
		os.Exit(ok)
	}
}
