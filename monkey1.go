// This program is a rewrite of a ruby program that I made to mimic the infinate monkeys
// with infinate type writers trying to write the works of Shakespeare.
// It is very quick but would still take eons to find the whole works of Shakespeare.
// https://en.wikipedia.org/wiki/Infinite_monkey_theorem
// RandString was adapted from 
// http://stackoverflow.com/questions/22892120/how-to-generate-a-random-string-of-a-fixed-length-in-golang 
// This version has no threading and acts as a single monkey.
// Usage monkey <word in lower case>
// No input checking is done as this is experimental code and not really a program of any use.
package main
import (
    "time"
    "fmt"
    "math/rand"
    "os"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyz "
const (
    letterIdxBits = 6                    // 6 bits to represent a letter index
    letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
    letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

func RandString(n int,src rand.Source) string {
    b := make([]byte, n)
    // A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
    for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
        if remain == 0 {
            cache, remain = src.Int63(), letterIdxMax
        }
        if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
            b[i] = letterBytes[idx]
            i--
        }
        cache >>= letterIdxBits
        remain--
    }
    return string(b)
}

func monkeyTyper(stringToMatch string, stringLength int) string {
    startTime := time.Now()
    for i := 1; i > 0; i++ {
        if stringToMatch == RandString(stringLength, rand.NewSource(time.Now().UnixNano())) {
            fmt.Printf("It took %d tries to type out '%s' and took %s of time.\n", i, stringToMatch, time.Since(startTime))
            os.Exit(0)
        }
    }
    return "error"
}

func main() { 
    monkeyTyper(os.Args[1], len(os.Args[1]))
}
