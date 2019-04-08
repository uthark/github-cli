package github

import (
	"fmt"
)

type RepoPrinter struct {
}

func (printer RepoPrinter) Print(sortType string, prs RepositoriesWithPR) {

	fmt.Println("Sorted by: ", sortType)
	fmt.Println("Forks\tStars\tPRs\tContrib\tRepo Name")
	fmt.Println(prs)

}
