// Copyright © 2017 Bibai Jin <bibaijin@gmail.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"errors"
	// "fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

var outputFile string

// md2pdfCmd represents the md2pdf command
var md2pdfCmd = &cobra.Command{
	Use:   "md2pdf INPUT_FILE",
	Short: "Convert from Markdown to PDF",
	Args:  cobra.ExactArgs(1),
	RunE:  md2pdf,
}

func init() {
	RootCmd.AddCommand(md2pdfCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// md2pdfCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	md2pdfCmd.Flags().StringVarP(&outputFile, "output", "o", "", "output file")
}

func md2pdf(cmd *cobra.Command, args []string) error {
	if outputFile == "" {
		return errors.New("-o/--output is required")
	}

	execCmd := exec.Command("pandoc", args[0], "-o", outputFile, "-t", "latex", "--latex-engine=xelatex", "--data-dir=/home/vagrant/.config/pandoc", "--template=default-zh.latex", "-V", "colorlinks=true", "-V", "fontsize=12pt", "-V", "lang=zh", "-V", "papersize=a4", "-V", "documentclass=article", "-V", "linestretch=1.4", "-V", "CJKmainfont=Source Han Serif CN", "-V", "CJKoptions=BoldFont=Source Han Sans CN", "-V", "geometry=top=1.2in,bottom=1.2in,left=1.2in,right=1in", "-F", "pandoc-crossref", "-F", "pandoc-citeproc", "-M", "crossrefYaml=/home/vagrant/.config/pandoc/crossref-zh.yaml")
	execCmd.Stdout = os.Stdout
	execCmd.Stderr = os.Stderr
	if err := execCmd.Run(); err != nil {
		return err
	}
	return nil
}
