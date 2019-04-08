package github

import (
	"log"
	"sort"
	"sync"
)

// ShowStats shows stats for the repos in given organization.
func (client Client) ShowStats(org string, projectCount int, sortType string) error {
	repos, err := client.ListOrganizationRepos(org, projectCount, sortType)
	if err != nil {
		log.Println(err)
	}

	var sorter sort.Interface

	// For some sort types we need to fetch pull requests info.
	if sortType == SortPRs || sortType == SortContrib {
		// fetch extra info
		repoWithPRs := RepositoriesWithPR{}

		var wg sync.WaitGroup
		for _, repo := range repos {
			wg.Add(1)
			go func(repository RepositoryWithPRs) {
				// repository := repo
				prCount, err := client.CountPullRequests(org, repository.Name)
				repository.PRCount = prCount
				repoWithPRs = append(repoWithPRs, repository)
				if err != nil {
					log.Println(err)
				}
				wg.Done()
			}(repo)
		}
		wg.Wait()

		repos = repoWithPRs
	}

	sorter = getSorter(repos, sortType)
	sort.Sort(sorter)
	idx := min(sorter.Len(), projectCount)

	printer := RepoPrinter{}
	printer.Print(sortType, repos[:idx])

	return nil
}

func getSorter(repos RepositoriesWithPR, sortType string) sort.Interface {
	switch sortType {
	case SortStars:
		return StargazersSorter(repos)
	case SortForks:
		return ForksSorter(repos)
	case SortPRs:
		return PRsSorter(repos)
	case SortContrib:
		return ContribSorter(repos)
	default:
		return repos
	}
}
