package main

import (
    "fmt"

    "github.com/thoas/go-funk"
)

func main() {
    slice1 := []int{1, 3, 5, 7, 9}
    slice2 := []int{1, 2, 3, 4, 5, 6}

    // Subtract returns the subtraction between two collections.
    subtract := funk.Subtract(slice1, slice2)
    fmt.Println(subtract)

    // Intersect returns the intersection between two collections.
    intersect := funk.Intersect(slice1, slice2)
    fmt.Println(intersect)

    // Difference returns the difference between two collections.
    slice_1_Diff, slice_2_Diff := funk.Difference(slice1, slice2)
    fmt.Println(slice_1_Diff)
    fmt.Println(slice_2_Diff)

    // Zip returns a list of tuples, where the i-th tuple contains the i-th element from each of the input iterables.
    zip := funk.Zip(slice1, slice2)
    fmt.Println(zip)

    // Subset returns true if collection x is a subset of y.
    subset := funk.Subset(slice1, slice2)
    fmt.Println(subset)

    // Equal returns if the two objects are equal
    Equal := funk.Equal(slice1, slice2)
    fmt.Println(Equal)

    // Contains returns true if an element is present in a iteratee.
    convertSlice := funk.Contains(slice2, slice1)
    fmt.Println(convertSlice)
}