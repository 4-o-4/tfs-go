package main

import "fmt"

const (
    reset   = "\033[0m"
    cRed    = "\033[31m"
    cGreen  = "\033[32m"
    cYellow = "\033[33m"
    cBlue   = "\033[34m"
    cPurple = "\033[35m"
    cCyan   = "\033[36m"
    cWhite  = "\033[37m"
)

var color = map[string]Color{
    "red":    cRed,
    "green":  cGreen,
    "yellow": cYellow,
    "blue":   cBlue,
    "purple": cPurple,
    "cyan":   cCyan,
    "white":  cWhite,
}

type Color string

func (c *Color) set(clr string) {
    *c = color[clr]
}

func (c *Color) reset() {
    *c = reset
}

func sandglass(size int, args ...string) {
    var outline = "X"
    var c Color

    for i, arg := range args {
        switch {
        case i > 1:
            break
        case len(arg) == 1:
            outline = arg
        default:
            c.set(arg)
        }
    }
    stamp(size, outline, c)
}

func stamp(size int, outline string, c Color) {
    fmt.Print(c)
    c.reset()
    defer fmt.Print(c)

    for i := range size {
        for j := range size {
            switch {
            case i == 0 || i == size-1:
                fmt.Print(outline)
            case i == j || j == size-1-i:
                fmt.Print(outline)
            default:
                fmt.Print(" ")
            }
        }
        fmt.Println()
    }
}

func main() {
    sandglass(11, "*", "green")
    sandglass(7)
}
