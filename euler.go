package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"jalopymusic.com/euler/problems"
)

func main() {
	solveCmd := flag.NewFlagSet("solve", flag.ExitOnError)
	solveProblem := solveCmd.Int("problem", 1, "problem")
	solveArgs := solveCmd.String("args", "[]", "args")

	if len(os.Args) < 2 {
		fmt.Println("expected 'solve' subcommand")
		os.Exit(1)
	}

	switch os.Args[1] {

	case "solve":
		solveCmd.Parse(os.Args[2:])
		p, err := problems.New(*solveProblem)
		if err != nil {
			fmt.Printf("Failed to construct problem: %#v\n", err)
			os.Exit(1)
		}
		var args []interface{}
		err = json.Unmarshal([]byte(*solveArgs), &args)
		if err != nil {
			fmt.Printf("Failed to unmarshal args: %#v\n", err)
			os.Exit(1)
		}
		statement, solution, err := p.Solve(args...)
		if err != nil {
			fmt.Printf("Failed to solve: %#v\n", err)
			os.Exit(1)
		}
		fmt.Printf("%v\n\n", statement)
		fmt.Printf("%#v -> %#v\n", *solveArgs, solution)

	default:
		fmt.Println("expected 'solve' subcommand")
		os.Exit(1)
	}

}
