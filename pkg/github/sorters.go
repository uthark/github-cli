package github

// StargazersSorter implements Sorter interface and sorts by stargazers count.
type StargazersSorter RepositoriesWithPR

func (r StargazersSorter) Len() int           { return len(r) }
func (r StargazersSorter) Swap(i, j int)      { r[i], r[j] = r[j], r[i] }
func (r StargazersSorter) Less(i, j int) bool { return r[i].StargazersCount > r[j].StargazersCount }

// ForksSorter implements Sorter interface and sorts by forks count.
type ForksSorter RepositoriesWithPR

func (r ForksSorter) Len() int           { return len(r) }
func (r ForksSorter) Swap(i, j int)      { r[i], r[j] = r[j], r[i] }
func (r ForksSorter) Less(i, j int) bool { return r[i].ForksCount > r[j].ForksCount }

func (r RepositoriesWithPR) Len() int           { return len(r) }
func (r RepositoriesWithPR) Swap(i, j int)      { r[i], r[j] = r[j], r[i] }
func (r RepositoriesWithPR) Less(i, j int) bool { return r[i].PRCount > r[j].PRCount }

// PRsSorter implements Sorter interface and sorts by PRs count.
type PRsSorter RepositoriesWithPR

func (r PRsSorter) Len() int           { return len(r) }
func (r PRsSorter) Swap(i, j int)      { r[i], r[j] = r[j], r[i] }
func (r PRsSorter) Less(i, j int) bool { return r[i].PRCount > r[j].PRCount }

// ContribSorter implements Sorter interface and sorts by PRs/forks count.
type ContribSorter RepositoriesWithPR

func (r ContribSorter) Len() int           { return len(r) }
func (r ContribSorter) Swap(i, j int)      { r[i], r[j] = r[j], r[i] }
func (r ContribSorter) Less(i, j int) bool { return r[i].Contrib() > r[j].Contrib() }
