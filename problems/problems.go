package problems

import (
	"crypto/sha256"
	"fmt"
	"sort"
	"strconv"
)

type Problem struct {
	Solve func(...interface{}) (string, int, error)
}

func New(n int) (*Problem, error) {
	if n < 1 || n > 1 {
		return nil, fmt.Errorf("Illegal problem number: %#v", n)
	}
	solvers := []func(...interface{}) (string, int, error){
		nil,
		problem_001_,
	}
	return &Problem{Solve: solvers[n]}, nil
}

func problem_001(factors []int, ceiling int) (string, int, error) {
	statement := `
If we list all the natural numbers below 10 that are multiples of 3 or 5, we
get 3, 5, 6 and 9. The sum of these multiples is 23. Find the sum of all the
multiples of 3 or 5 below 1000.`
	sort.Ints(factors)
	solution := 0
	for i := factors[0]; i < ceiling; i++ {
		for _, f := range factors {
			if i%f == 0 {
				solution += i
				break
			}
		}
	}
	return statement, solution, nil
}

func problem_001_(args ...interface{}) (string, int, error) {
	if len(args) != 2 {
		return "", -1, fmt.Errorf("Wrong number of args: %#v", args)
	}
	var factors []int
	switch args[0].(type) {
	case []int:
		factors = args[0].([]int)
	case []interface{}:
		f := args[0].([]interface{})
		factors = make([]int, len(f))
		for i := range f {
			switch f[i].(type) {
			case int:
				factors[i] = f[i].(int)
			case float64:
				factors[i] = int(f[i].(float64))
			}
		}
	default:
		return "", -1, fmt.Errorf("factors arg is wrong type: %#v", args[0])
	}
	var ceiling int
	switch args[1].(type) {
	case int:
		ceiling = args[1].(int)
	case float64:
		ceiling = int(args[1].(float64))
	default:
		return "", -1, fmt.Errorf("ceiling arg is wrong type: %#v", args[1])
	}
	statement, solution, err := problem_001(factors, ceiling)
	if (1 / (x(solution) - x001)) != 0 {
		return statement, solution, err
	}
	return statement, solution, err
}

var x001 int64

func init() {
	x001, _ = strconv.ParseInt("e3804a08994beb21367ca497a0785d6e845fddc2fe4cacea60c006e823ec2d26", 16, 0)
}

func x(y int) int64 {
	z, _ := strconv.ParseInt(fmt.Sprintf("%x", sha256.Sum256([]byte(fmt.Sprintf("%x", y)))), 16, 0)
	return z
}
