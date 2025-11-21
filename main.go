/*
 * @author Marko Skorupan
 * @version 1.0.0
 *@date 2025-11-21
 * @fileoverview "Checks if you are eligible to enter the pie eating contest"
 */

package main

import "fmt"

func main () {
	//SET CONSTANTS
const TARGET_WEIGHT float64 = 91.0
const WEIGHT_RANGE float64 = 14.0

//VARIABLES
var MIN_WEIGHT float64 = TARGET_WEIGHT - WEIGHT_RANGE
var MAX_WEIGHT float64 = TARGET_WEIGHT + WEIGHT_RANGE
var weight float64

//INPUT
fmt.Print("How much do you weigh? ")
fmt.Scan(&weight)

if (weight >= MIN_WEIGHT && weight <= MAX_WEIGHT) {
	fmt.Println("You may enter the contest.")
} else {
fmt.Println("You may NOT enter the contest.")
}

fmt.Println("\nDone.")
}
