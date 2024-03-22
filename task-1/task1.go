package task1

import (
	"errors"
	"fmt"
	"log"
	"math"
)

func validateInput(a, b int) error {
	if (a < 0 || a >= 1000) || (b < 0 || b >= 1000) {
		return errors.New("INVALID INPUT")
	}
	return nil
}

func task1(a, b int) (float64, error) {
	if err := validateInput(a, b); err != nil {
		return 0.0, err
	}
	return math.Hypot(float64(a), float64(b)), nil
}

func mainTask1() {
	var (
		a, b int
		res  float64
	)
	fmt.Scan(&a, &b)
	res, err := task1(a, b)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res)
}
