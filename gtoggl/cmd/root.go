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
	"fmt"
	"os"

	"encoding/json"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gopkg.in/dougEfresh/gtoggl.v8"
	"gopkg.in/dougEfresh/toggl-http-client.v8"
)

var cfgFile string

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "gtoggl",
	Short: "Toggl API cli",
	Long:  `Toggl CLI`,
	PreRun: func(cmd *cobra.Command, args []string) {
		fmt.Println("asdasdsad")
		//return errors.New("some random error")
	},
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//	Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports Persistent Flags, which, if defined here,
	// will be global for your application.
	RootCmd.PersistentFlags().BoolP("debug", "d", false, "Debuging")
	RootCmd.PersistentFlags().StringP("token", "t", "", "api token")
	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.gtoggl.yaml)")
	viper.BindPFlag("token", RootCmd.PersistentFlags().Lookup("token"))
	// Cobra also supports local flags, which will only run
	// when this action is called directly.

}

type debugger struct {
	debug bool
}

func (l *debugger) Printf(format string, v ...interface{}) {
	if l.debug {
		fmt.Printf(format, v)
	}
}

var tc *gtoggl.TogglClient

func getClient(d bool) *gtoggl.TogglClient {
	tc, err := gtoggl.NewClient(viper.GetString("token"), ghttp.SetTraceLogger(&debugger{debug: d}))
	if err != nil {

	}
	return tc
}

func printJson(a interface{}, err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error %s", err)
		os.Exit(-1)
	}
	j, _ := json.Marshal(a)
	fmt.Fprintf(os.Stdout, "%+s\n", j)
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" { // enable ability to specify config file via flag
		viper.SetConfigFile(cfgFile)
	}

	viper.SetConfigName(".gtoggl") // name of config file (without extension)
	viper.AddConfigPath("$HOME")   // adding home directory as first search path
	viper.AutomaticEnv()           // read in environment variables that match
	d, _ := RootCmd.Flags().GetBool("debug")
	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		if d {
			fmt.Fprintf(os.Stderr, "Using config file:%s\n", viper.ConfigFileUsed())
		}
	}

	if h, _ := RootCmd.Flags().GetBool("help"); !h {
		if viper.GetString("token") == "" {
			fmt.Fprintf(os.Stderr, "Token Required\n")
			RootCmd.Help()
			os.Exit(-1)
		}
		tc = getClient(d)
	}
}
