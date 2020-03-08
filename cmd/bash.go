/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

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
	"os"

	"github.com/spf13/cobra"
)

// bashCmd represents the bash command
var bashCmd = &cobra.Command{
	Use:   "bash",
	Short: "Generate bash completion",
	Long: `This command will generate a Bash script for enable the go-web-tunnel autocompletion.
	Make sure that you have https://github.com/scop/bash-completion properly installed.
	You can verify by running type _init_completion.
	If the previus command succeeds, you're already set, otherwise click the above link and follow the installation guide.

	Now add the completion script genereted by running go-web-tunnel completion bash in your /etc/bash_completion.d
	running the following command:
		$ sudo sh -c  "go-web-tunnel completion bash > /etc/bash_completion.d/go-web-tunnel" "
	`,
	Run: func(cmd *cobra.Command, args []string) {
		rootCmd.GenBashCompletion(os.Stdout)
	},
}

func init() {
	completionCmd.AddCommand(bashCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// bashCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// bashCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
