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
	"log"

	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"
)

var outputDir string

// docsCmd represents the docs command
var docsCmd = &cobra.Command{
	Use:   "docs [docs type]",
	Short: "Generate docs",
	Long:  `Generate docs for the go-web-tunnel application.`,
}

var genMarkDownCmd = &cobra.Command{
	Use:   "markdown",
	Short: "Generate markdown docs",
	Long:  `Generate markdown docs for the go-web-tunnel application.`,
	Run: func(cmd *cobra.Command, args []string) {
		err := doc.GenMarkdownTree(rootCmd, outputDir)
		if err != nil {
			log.Fatal(err)
		}
	},
}

var genManCmd = &cobra.Command{
	Use:   "man",
	Short: "Generate man pages",
	Long:  `Generate Man pages for the go-web-tunnel application.`,
	Run: func(cmd *cobra.Command, args []string) {
		header := &doc.GenManHeader{
			Title:   "go-web-tunnel",
			Section: "1",
		}
		err := doc.GenManTree(rootCmd, header, outputDir)
		if err != nil {
			log.Fatal(err)
		}
	},
}

var genReSTCmd = &cobra.Command{
	Use:   "rest",
	Short: "Generate ReST docs",
	Long:  `Generate ReStructured text docs for the go-web-tunnel application.`,
	Run: func(cmd *cobra.Command, args []string) {
		err := doc.GenReSTTree(rootCmd, outputDir)
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(docsCmd)
	docsCmd.PersistentFlags().StringVarP(&outputDir, "outputDir", "o", ".", "output directory")

	docsCmd.AddCommand(genMarkDownCmd)
	docsCmd.AddCommand(genManCmd)
	docsCmd.AddCommand(genReSTCmd)
}
