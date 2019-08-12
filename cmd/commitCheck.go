/*
Copyright © 2019 labs at aspiration dot com

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"github.com/aspiration-labs/git-gear/internal/jira"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"io/ioutil"
)

// commitCheckCmd represents the commitCheck command
var commitCheckCmd = &cobra.Command{
	Use:   "commitCheck",
	Short: "Check for valid Jira smart commit messages",
	Long: `Require Jira smart commit ticket references, checking that those tickets exist in the provided Jira server.
See https://confluence.atlassian.com/fisheye/using-smart-commits-960155400.html
`,
	Args: cobra.ExactArgs(1),
	Run: commitCheck,
}

func commitCheck(cmd *cobra.Command, args []string) {
	viper.WriteConfig()
	jiraRepo := jira.NewJiraRepo()
	message, err := ioutil.ReadFile(args[0])
	if err != nil {
		log.Fatal(err)
		log.Exit(1)
	}
	jiraRepo.CheckMessage(string(message))
}

func init() {
	jiraCmd.AddCommand(commitCheckCmd)
}
