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
	"fmt"
	"github.com/aspiration-labs/git-gear/internal/jira"
	"strings"

	"github.com/spf13/cobra"
)

// projectKeyListCmd represents the projectKeyList command
var projectKeyListCmd = &cobra.Command{
	Use:   "projectKeyList",
	Short: "Print the list of valid projects on the Jira server",
	Run: projectKeyList,
}

func projectKeyList(cmd *cobra.Command, args []string) {
	jiraRepo := jira.NewJiraRepo()
	fmt.Println(strings.Join(jiraRepo.GetProjectKeys(), "\n"))
}

func init() {
	jiraCmd.AddCommand(projectKeyListCmd)
}
