package github

import (
	"regexp"
	"strconv"
	"strings"
)

// Links contains information parsed from Link header.
type Links struct {
	FirstPage int32
	PrevPage  int32
	NextPage  int32
	LastPage  int32
}

// ParseLinks parses given string and extracts information about first/prev/next/last pages.
func ParseLinks(s string) (Links, error) {
	if s == "" {
		return Links{}, nil
	}
	links := strings.Split(s, ",")
	result := Links{}
	regex := regexp.MustCompile(`<.+?[?&]page=(\d+)(?:.+)?>;\s+rel="(\w+)"`)
	for _, l := range links {
		submatch := regex.FindStringSubmatch(l)
		if len(submatch) == 3 {
			value := submatch[1]
			intVal, err := strconv.Atoi(value)
			if err != nil {
				return Links{}, err
			}

			rel := submatch[2]
			switch rel {
			case "next":
				result.NextPage = int32(intVal)
			case "last":
				result.LastPage = int32(intVal)
			case "first":
				result.FirstPage = int32(intVal)
			case "prev":
				result.PrevPage = int32(intVal)

			}
		}
	}

	return result, nil
}
