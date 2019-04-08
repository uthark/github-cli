package github

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseLinks(t *testing.T) {
	links, err := ParseLinks(`<https://api.github.com/organizations/7634182/repos?per_page=5&page=2>; rel="next",
	 <https://api.github.com/organizations/7634182/repos?per_page=5&page=7>; rel="last"`)
	assert.NoError(t, err)

	assert.Equal(t, int32(2), links.NextPage, "NextPage is not parsed")
	assert.Equal(t, int32(7), links.LastPage, "LastPage is not parsed")
}

func TestParseLinks2(t *testing.T) {
	links, err := ParseLinks(`<https://api.github.com/organizations/7634182/repos?page=4&per_page=5>; rel="prev", 
	<https://api.github.com/organizations/7634182/repos?page=6&per_page=5>; rel="next", 
	<https://api.github.com/organizations/7634182/repos?page=7&per_page=5>; rel="last", 
	<https://api.github.com/organizations/7634182/repos?page=1&per_page=5>; rel="first"`)
	assert.NoError(t, err)

	assert.Equal(t, int32(1), links.FirstPage, "FirstPage is not parsed")
	assert.Equal(t, int32(4), links.PrevPage, "PrevPage is not parsed")
	assert.Equal(t, int32(6), links.NextPage, "NextPage is not parsed")
	assert.Equal(t, int32(7), links.LastPage, "LastPage is not parsed")
}
