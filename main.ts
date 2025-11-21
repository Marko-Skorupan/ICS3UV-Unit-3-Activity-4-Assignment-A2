/**
 * @author Marko Skorupan
 * @version 1.0.0
 * @date 2025-11-21
 * @fileoverview "Checks if you are eligable to enter the pie eating contest"
 */
//SET CONSTANTS
const TARGET_WEIGHT: number = 91.0
const WEIGHT_RANGE: number = 14.0

//PROCESS
let MIN_WEIGHT = TARGET_WEIGHT - WEIGHT_RANGE
let MAX_WEIGHT = TARGET_WEIGHT + WEIGHT_RANGE

//INPUT
let weight: number = Number(prompt("How much do you weigh?"));

if (weight >= MIN_WEIGHT && weight <= MAX_WEIGHT) {
  console.log("You may enter the contest.")
}
else {
  console.log("You may NOT enter the contest.")
}

console.log("\nDone.")
