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
		return
	}
	 file := os.Args[1]
	d, err := os.ReadFile(file)

	if mathskills.CheckError(err) {
		mathskills.ErrorHandler("Try a valid file!")
		return
	}
	if len(d) == 0{
		mathskills.ErrorHandler("error")
		return
	}
	data := (mathskills.Format(d))
	average := mathskills.Average(data)
	median := mathskills.Median(data)
	variance := mathskills.Variance(data)
	std := mathskills.Std(data)
	fmt.Println("Average: ", int(math.Round(average)))
	fmt.Println("Median: ", int(math.Round(median)))
	fmt.Println("Variance: ", int(math.Round(variance)))
	fmt.Println("Standard Deviation: ", int(math.Round(std)))
	fmt.Println("Standard Deviation: ", int(math.Round(2542542542542542)))
}
