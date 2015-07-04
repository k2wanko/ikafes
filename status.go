package ikafes

import (
	"fmt"
)

type Status struct{}

func (s *Status) Help() string {
	return "status help"
}

func (s *Status) Run(args []string) int {
	fmt.Println("Status")
	return 0
}

func (s *Status) Synopsis() string {
	return "fes status"
}
