package github

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateRepoRequestAsJson(t * testing.T) {
	request := CreateRepoRequest{
		Name: 		 "golang intro",
		Description: "a golang intro repo",
		Homepage: 	 "https://github",
		Private: 	 true,
		HasIssues: 	 false,
		HasProjects: true,
		HasWiki: 	 false,
	}

	if request.Private {

	}

	bytes, err := json.Marshal(request)
	assert.Nil(t, err)
	assert.NotNil(t, bytes)

	var target CreateRepoRequest
	err = json.Unmarshal(bytes, &target)
	assert.Nil(t, err)

	assert.EqualValues(t, target.Name, request.Name)
	assert.EqualValues(t, target.HasIssues, request.HasIssues)
}
