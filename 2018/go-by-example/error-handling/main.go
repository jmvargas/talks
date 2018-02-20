package main

import (
	"errors"
	"fmt"
	"log"
	"math"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// START OMIT

func main() {
	//Here we're pretty sure that 25 never causes error
	sqrtOf25, _ := Sqrt(25.0)
	fmt.Printf("Sqrt of 25 is %g\n", sqrtOf25)

	//Here we aren't sure
	sqrtOfSomeFunc, err := Sqrt(rand.Float64() - .5)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Sqrt of SomeFunc is %g\n", sqrtOfSomeFunc)
}

func Sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, errors.New("math: square root of negative number")
	}
	return math.Sqrt(f), nil
}

// END OMIT
