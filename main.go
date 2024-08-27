package main

import (
	"fmt"
	"math"
	mathskills "mathskills/pkg"
	"os"
)

// Main function
func main() {
	if len(os.Args) != 2{
		mathskills.ErrorHandler("Please enter a single file.")
	}
	 file := os.Args[1]
	d, err := os.ReadFile(file)

	if mathskills.CheckError(err) {
		mathskills.ErrorHandler("Try a valid file!")
		return
	}

	data := (mathskills.Format(d))
	average := mathskills.Average(data)
	median := mathskills.Median(data)
	variance := mathskills.Variance(data)
	std := mathskills.Std(data)
	fmt.Println("Average: ", math.Round(average))
	fmt.Println("Median: ", math.Round(median))
	fmt.Println("Variance: ", math.Round(variance))
	fmt.Println("Standard Deviation: ", math.Round(std))
}
