package discord

import "github.com/hugolgst/rich-go/client"

type RPC struct {
	loggedIn bool
	Current  *client.Activity
}

func (r *RPC) Login() {
	if !r.loggedIn {
		err := client.Login("")
		if err == nil {
			r.loggedIn = true
		} else {
			r.loggedIn = false
		}
	}
}

func (r *RPC) Logout() {
	if r.loggedIn {
		client.Logout()
		r.loggedIn = false
	}
}

func (r *RPC) SetActivity(activity client.Activity) {
	if r.loggedIn {
		r.Current = &activity
		client.SetActivity(activity)
	}
}
