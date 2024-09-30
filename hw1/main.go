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

var cMap = map[int]Color{
    315: cRed,
    529: cGreen,
    668: cYellow,
    424: cBlue,
    664: cPurple,
    427: cCyan,
    545: cWhite,
}

type Color string

func (c *Color) reset() {
    *c = reset
}

type Arg func(map[string]int)

func size(i int) Arg {
    return func(m map[string]int) {
        m["size"] = i
    }
}

func char(i int) Arg {
    return func(m map[string]int) {
        m["char"] = i
    }
}

func color(str string) Arg {
    var hash int
    for _, s := range str {
        hash += int(s)
    }
    return func(m map[string]int) {
        m["color"] = hash
    }
}

func stamp(size int, outline rune, c Color) {
    fmt.Print(c)
    c.reset()
    defer fmt.Print(c)

    for i := range size {
        for j := range size {
            switch {
            case i == 0 || i == size-1:
                fmt.Print(string(outline))
            case i == j || j == size-1-i:
                fmt.Print(string(outline))
            default:
                fmt.Print(" ")
            }
        }
        fmt.Println()
    }
}

func sandglass(args ...Arg) {
    var m = map[string]int{
        "size": 15,
        "char": 'X',
    }
    for _, arg := range args {
        arg(m)
    }
    stamp(m["size"], rune(m["char"]), cMap[m["color"]])
}

func main() {
    sandglass(size(7), char('*'), color("green"))
    sandglass()
}
