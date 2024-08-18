package main

import "fmt"

func main() {
	PrintHourglass(9, "X")
}

func PrintHourglass(width int, outline string) {
	for i := range width {
		for j := range width {
			switch {
			case i == 0 || i == width-1:
				fmt.Print(outline)
			case i == j || j == width-1-i:
				fmt.Print(outline)
			default:
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}
