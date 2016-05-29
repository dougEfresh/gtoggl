package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"gopkg.in/dougEfresh/gtoggl.v8"
	"gopkg.in/dougEfresh/toggl-client.v8"
	"gopkg.in/dougEfresh/toggl-http-client.v8"
	"gopkg.in/dougEfresh/toggl-user.v8"
	"gopkg.in/dougEfresh/toggl-workspace.v8"
	"os"
	"strconv"
)

type debugger struct {
	debug bool
}

func (l *debugger) Printf(format string, v ...interface{}) {
	if l.debug {
		fmt.Printf(format, v)
	}
}

func main() {
	var debug = flag.Bool("d", false, "Debug")
	var token = flag.String("t", "", "Toggl API token: https://www.toggl.com/app/profile")
	var command = flag.String("c", "workspace", "command: workspace,client,project...etc ")
	flag.Parse()
	tc, err := gtoggl.NewClient(*token, ghttp.SetTraceLogger(&debugger{debug: *debug}))
	if err != nil {
		fmt.Fprint(os.Stderr, "A token is required\n")
		flag.Usage()
		os.Exit(-1)
	}
	if *command == "workspace" {
		workspace(tc, flag.Args())
	}
	if *command == "client" {
		client(tc, flag.Args())
	}
	if *command == "user" {
		user(tc, flag.Args())
	}
}

func handleError(error error) {
	if error != nil {
		fmt.Fprintln(os.Stderr, error)
		os.Exit(-1)
	}
}

func client(tc *gtoggl.TogglClient, args []string) {
	c := tc.TClient
	var client gclient.Client
	var err error
	if len(args) == 0 || args[0] == "list" {
		clients, err := c.List()
		handleError(err)
		fmt.Printf("%+v\n", clients)
	}

	if args[0] == "create" && len(args) > 1 {
		err = json.Unmarshal([]byte(args[1]), &client)
		handleError(err)
		nClient, err := c.Create(&client)
		handleError(err)
		fmt.Printf("%+v\n", nClient)
	}

	if args[0] == "update" && len(args) > 1 {
		err = json.Unmarshal([]byte(args[1]), &client)
		handleError(err)
		_, err = c.Get(client.Id)
		handleError(err)
		nClient, err := c.Update(&client)
		handleError(err)
		fmt.Printf("%+v\n", nClient)
	}

	if args[0] == "get" && len(args) > 1 {
		i, err := strconv.ParseUint(args[1], 0, 64)
		handleError(err)
		nClient, err := c.Get(i)
		handleError(err)
		fmt.Printf("%+v\n", nClient)
	}

	if args[0] == "delete" && len(args) > 1 {
		i, err := strconv.ParseUint(args[1], 0, 64)
		handleError(err)
		err = c.Delete(i)
		handleError(err)
		fmt.Printf("%+v  deleted\n", i)
	}

}

func workspace(tc *gtoggl.TogglClient, args []string) {
	wsc := tc.WorkspaceClient
	var err error
	if len(args) == 0 || args[0] == "list" {
		w, err := wsc.List()
		handleError(err)
		fmt.Printf("%+v\n", w)
		return
	}

	if args[0] == "get" && len(args) > 1 {
		i, err := strconv.ParseUint(args[1], 0, 64)
		handleError(err)
		w, err := wsc.Get(i)
		fmt.Printf("%+v\n", w)
		return
	}

	if args[0] == "update" && len(args) > 1 {
		var uWs gworkspace.Workspace
		handleError(err)
		err = json.Unmarshal([]byte(args[1]), &uWs)
		handleError(err)
		_, err := wsc.Get(uWs.Id)
		handleError(err)
		nWs, err := wsc.Update(&uWs)
		handleError(err)
		fmt.Printf("%+v\n", nWs)
		return
	}
}

func user(tc *gtoggl.TogglClient, args []string) {
	wsc := tc.UserClient
	var err error
	if len(args) == 0 || args[0] == "get" {
		w, err := wsc.Get(false)
		handleError(err)
		js, err := json.Marshal(w)
		handleError(err)
		fmt.Printf("%+s\n", js)
		return
	}

	if args[0] == "reset" {
		w, err := wsc.ResetToken()
		handleError(err)
		fmt.Printf("%+v\n", w)
		return
	}

	if args[0] == "update" && len(args) > 1 {
		var uWs guser.User
		handleError(err)
		err = json.Unmarshal([]byte(args[1]), &uWs)
		handleError(err)
		_, err := wsc.Get(true)
		handleError(err)
		nWs, err := wsc.Update(&uWs)
		handleError(err)
		fmt.Printf("%+v\n", nWs)
		return
	}
	if args[0] == "create" && len(args) > 2 {
		var uWs *guser.User
		uWs, err = wsc.Create(args[1], args[2], args[3])
		handleError(err)
		js, err := json.Marshal(uWs)
		handleError(err)
		fmt.Printf("%+s\n", js)
		return
	}

}
