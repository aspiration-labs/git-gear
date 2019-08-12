package jira

import (
	"fmt"
	"github.com/andygrunwald/go-jira"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"net/url"
)

type jiraRepo struct {
	client *jira.Client
	projects *[]jira.Project
}

func NewJiraRepo() jiraRepo {
	jiraUrl := viper.GetString("jiraserver")
	if jiraUrl == "" {
		log.Fatal("--jiraserver is required, e.g., https://username:token@jira.atlassian.net:")
		log.Exit(1)
	}
	parsedUrl, err := url.Parse(jiraUrl)
	if err != nil {
		log.Fatal("Jira url bad format, should be like https://username:token@jira.atlassian.net:", err)
		log.Exit(1)
	}
	password, _ := parsedUrl.User.Password()
	auth := jira.BasicAuthTransport{
		Username: parsedUrl.User.Username(),
		Password: password,
	}
	server := fmt.Sprintf("%s://%s", parsedUrl.Scheme, parsedUrl.Host)
	client, err := jira.NewClient(auth.Client(), server)

	if err != nil {
		log.Fatal("Can't open Jira server", err)
		log.Exit(1)
	}

	return jiraRepo{
		client: client,
		projects: getProjects(client),
	}
}

func getProjects(client *jira.Client) *[]jira.Project {
	req, _ := client.NewRequest("GET", "rest/api/2/project", nil)

	projects := new([]jira.Project)
	_, err := client.Do(req, projects)
	if err != nil {
		log.Fatal("Can't access projects on Jira server", err)
		log.Exit(1)
	}
	return projects
}

func (repo jiraRepo) GetProjectKeys() []string {
	projectKeys := make([]string, len(*repo.projects))
	for i, project := range *repo.projects {
		projectKeys[i] = project.Key
	}
	return projectKeys
}
