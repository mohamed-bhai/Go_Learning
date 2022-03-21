package main

import (  
    "fmt"
    "testing"
)

func TestHello(t *testing.T) {
	got := Hello()
	want := "Hello, world"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

jiraIssuesStr = juno.getReleaseNoteSection(junoURL, junoToken, service, version, "jira_tickets").join(",")  
103^

releaseMetadata = juno.getPayload(junoURL, junoToken, requestType).release_metadata
117

environments = releaseMetadata.environment
118

