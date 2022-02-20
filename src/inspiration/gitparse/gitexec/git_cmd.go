package gitexec

import "os/exec"

const execName = "git"

func Command(arg ...string) (string, error) {
	cmd := exec.Command(execName, arg...)
	// cmd.Output()
	execResult, err := cmd.CombinedOutput()
	return string(execResult), err
}
