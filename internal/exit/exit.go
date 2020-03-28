// Package exit provides various exit codes for stack, used by the likes of
// os.Exit() when things don't go to plan.
package exit

const (
	RootExecute = iota
	HomeNotFound
	ConfigNotFound
	UnpushedCommits

	CmdGitError = 128
)
