// Copyright © 2017 Bibai Jin <bibaijin@gmail.com>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

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
