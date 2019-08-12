package jira

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"regexp"
	"strings"
)

func (repo jiraRepo) CheckMessage(message string) bool {

	reStr := fmt.Sprintf("((%s)-\\d+)", strings.Join(repo.GetProjectKeys(), "|"))
	re, err := regexp.Compile(reStr)
	if err != nil {
		log.Fatal("Error in message regex", err)
		log.Exit(1)
	}
	matches := re.FindAllString(message, -1)

	var topLevelTickets, subTickets int
	for _, match := range matches {
		issue, _, err := repo.client.Issue.Get(match, nil)
		if err != nil {
			log.Fatal("Not a valid issue on the Jira server: ", match)
			log.Exit(1)
		}
		if issue.Fields.Type.Subtask {
			subTickets++
		} else {
			topLevelTickets++
		}
	}

	if subTickets == 0 && topLevelTickets == 0 {
		log.Fatal("no valid Jira ticket references found; add one or more PROJ-1234 to your commit message")
		log.Exit(1)
	}
	if topLevelTickets == 0 {
		log.Fatal("at least 1 parent ticket is required; add a parent ticket id to your commit message")
		log.Exit(1)
	}

	return true
}
