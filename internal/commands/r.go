package commands

import "github.com/untillpro/qs/gitcmds"

func R(wd string) error {
	return gitcmds.Release(wd)
}
