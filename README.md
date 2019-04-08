# github-cli

## Description
Sample CLI for Github

## Building

### Pre-requisites

1. Go 1.12.2+. *Tested with this version*
1. Go Modules enabled (`GO111MODULE=on`) 

### Build the program
```
    # download all dependencies.
    GO111MODULE=on go mod tidy
    # Build binary.
    GO111MODULE=on go build
```

## Using

### Prepare personal Github Account

1. [Generate new Personal Token](https://github.com/settings/tokens/new) with `read:org` scope.
1. Export the token:

```
    export GITHUB_AUTH_TOKEN="<YOUR_GITHUB_USERNAME>:<YOUR_GITHUB_TOKEN>"
```

### Using OAuth 2.0
1. [Create new OAuth 2.0 Application](https://github.com/settings/applications/new)

**Application name** is `cli-github`

**Authorization callback URL** is `http://localhost:33999/oauth/callback`

**Homepage URL** is `https://github.com/uthark/github-cli`

1. Save `client_id` and `client_secret` in `config.toml`.
1. Authenticate in github:

```
    ./github-cli login
```

After successful login you can use application.

### Run application.

Example commands:

```
  ./github-cli org-repos netflix
  ./github-cli org-repos netflix -c 50 -s stars
  ./github-cli org-repos netflix -c 20 -s forks
  ./github-cli org-repos netflix -c 50 -s prs
  ./github-cli org-repos netflix -c 50 -s contrib
``` 

To view help:

```
./github-cli org-repos --help
```

### Configure which PRs to include 

Property `global.pr_state` in `config.toml` allows to specify which PRs include for calculation.   


## Further notes

1. Implement CLI using [go-github library](https://github.com/google/go-github). I decided to 
write my own just because using existing library was too simple.
1. It may worth to cache information from Github. Default limit of 5k requests may be not enough.
1. Output may be improved for further consumption (i.e rendering to JSON)
1. App supports OAuth 2.0 authorization, if configured in `config.toml`, even though for a single 
usage this is an overkill.

## References
1. [Github API](https://developer.github.com/v3/)
