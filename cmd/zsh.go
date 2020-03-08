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

// zshCmd represents the zsh command
var zshCmd = &cobra.Command{
	Use:   "zsh",
	Short: "Generate zsh completion",
	Long: `This command will generate a Zash script for enable the go-web-tunnel autocompletion.
	The go-web-tunnel completion script for Zsh can be generated with this command.
	Sourcing the completion script in your shell will enable go-web-tunnel autocompletion.
	To do so in all your shell sessions, add the following to your ~/.zshrc
		$ sudo sh -c "go-web-tunnel completion zsh > /usr/local/share/zsh/site-functions/_go-web-tunnel"  \
		&& echo "autoload -U compinit && compinit" >> ~/.zshrc

		$ source ~/.zshrc

	`,
	Run: func(cmd *cobra.Command, args []string) {
		rootCmd.GenZshCompletion(os.Stdout)
	},
}

func init() {
	completionCmd.AddCommand(zshCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// zshCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// zshCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
