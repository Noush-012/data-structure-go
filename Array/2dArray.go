package main

import "fmt"

func main() {
    // Declare a 2D array with 3 rows and 4 columns
    var myArray [3][4]int

    // Get user input for each element of the 2D array
    for i := 0; i < len(myArray); i++ {
        for j := 0; j < len(myArray[i]); j++ {
            fmt.Printf("Enter element at row %d, column %d: ", i, j)
            fmt.Scan(&myArray[i][j])
        }
    }

    // Print the value of the 2D array
    for i := 0; i < len(myArray); i++ {
        for j := 0; j < len(myArray[i]); j++ {
            fmt.Printf("%d ", myArray[i][j])
			fmt.Printf("%d ", myArray[i][j])
        }
        fmt.Println()
    }
}
