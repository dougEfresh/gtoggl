// Copyright © 2016 Douglas Chimento
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

	"github.com/spf13/cobra"
	"os"
)

var wlistCmd = &cobra.Command{
	Use:   "workspace",
	Short: "A brief description of your command",
	Long: `sasd`,
	Run: func(cmd *cobra.Command, args []string) {
		printJson(tc.WorkspaceClient.List())
	},
}

var wgetCmd = &cobra.Command{
	Use:   "get",
	Short: "get a workspace by id",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		id,err := cmd.Flags().GetUint64("id")
		if err != nil {
			os.Exit(-1)
		}
		printJson(tc.WorkspaceClient.Get(id))
	},
}

func init() {
	RootCmd.AddCommand(wlistCmd)
	wgetCmd.Flags().Uint64P("id","i",0,"id of resource")
	wlistCmd.AddCommand(wgetCmd)
}
