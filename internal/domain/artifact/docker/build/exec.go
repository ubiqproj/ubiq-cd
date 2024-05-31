package build

import (
	"bytes"
	"fmt"
	"io"
	"os/exec"
	"ubiq-cd/internal/domain/artifact"
)

func execCmd(
	workingDir artifact.AbsolutePath,
	options ...CmdEntryOptionApplier,
) error {
	stdout, stderr := bytes.Buffer{}, bytes.Buffer{}
	cmd := newCmd(newCmdEntry(options...),
		withWorkingDir(workingDir),
		withStdout(&stdout),
		withStderr(&stderr),
	)

	return cmd.Run()
}

type cmdEntryOption struct {
	dockerPath     string
	tags           []string
	buildArgs      map[string]string
	dockerfileName string
}

func defaultCmdOption() cmdEntryOption {
	return cmdEntryOption{
		dockerPath:     "docker",
		tags:           []string{},
		dockerfileName: "Dockerfile",
	}
}

func WithDocker(path string) CmdEntryOptionApplier {
	return func(option *cmdEntryOption) {
		option.dockerPath = path
	}
}

func WithTag(tag string) CmdEntryOptionApplier {
	return func(option *cmdEntryOption) {
		option.tags = append(option.tags, tag)
	}
}

func WithBuildArgs(buildArgs map[string]string) CmdEntryOptionApplier {
	return func(option *cmdEntryOption) {
		option.buildArgs = buildArgs
	}
}

func withDockerfile(filename string) CmdEntryOptionApplier {
	return func(option *cmdEntryOption) {
		option.dockerfileName = filename
	}
}

type CmdEntryOptionApplier func(*cmdEntryOption)

func newCmdEntry(options ...CmdEntryOptionApplier) (entry []string) {
	option := defaultCmdOption()
	for _, apply := range options {
		apply(&option)
	}

	// ...
	length := 5                    /* docker build -f <dockerfile> . */
	length += 2 * len(option.tags) /* -t <tag> */
	length += bool2int(len(option.buildArgs) > 0) +
		len(option.buildArgs) /* --build-args <args>... */

	entry = make([]string, 0, length)
	entry = append(entry, option.dockerPath, "build")
	entry = append(entry, tagOptions(option.tags)...)
	entry = append(entry, buildArgsOption(option.buildArgs)...)
	entry = append(entry, dockerfileOption(option.dockerfileName)...)
	entry = append(entry, ".")
	return entry
}

func bool2int(b bool) int {
	if b {
		return 1
	}
	return 0
}

func tagOptions(tags []string) (cmdArgs []string) {
	if len(tags) == 0 {
		return nil
	}

	cmdArgs = make([]string, 0, 2*len(tags))
	for _, tag := range tags {
		cmdArgs = append(cmdArgs, "-t", tag)
	}
	return cmdArgs
}

func buildArgsOption(buildArgs map[string]string) (cmdArgs []string) {
	if len(buildArgs) == 0 {
		return nil
	}

	cmdArgs = make([]string, 0, 1+len(buildArgs))
	cmdArgs = append(cmdArgs, "--build-args")
	for name, value := range buildArgs {
		cmdArgs = append(cmdArgs, fmt.Sprintf("%s='%s'", name, value))
	}
	return cmdArgs
}

func dockerfileOption(dockerfileName string) (cmdArgs []string) {
	if dockerfileName == "Dockerfile" {
		return nil
	}
	return []string{"-f", dockerfileName}
}

type cmdOption struct {
	workingDir *artifact.AbsolutePath
	stdout     *io.Writer
	stderr     *io.Writer
}

type cmdOptionApplier func(*cmdOption)

func withWorkingDir(path artifact.AbsolutePath) cmdOptionApplier {
	return func(option *cmdOption) {
		option.workingDir = &path
	}
}

func withStdout(out io.Writer) cmdOptionApplier {
	return func(option *cmdOption) {
		option.stdout = &out
	}
}

func withStderr(out io.Writer) cmdOptionApplier {
	return func(option *cmdOption) {
		option.stderr = &out
	}
}

func newCmd(entry []string, options ...cmdOptionApplier) *exec.Cmd {
	option := cmdOption{}
	for _, apply := range options {
		apply(&option)
	}

	cmd := exec.Command(entry[0], entry[1:]...)
	if option.workingDir != nil {
		cmd.Dir = string(*option.workingDir)
	}
	if option.stdout != nil {
		cmd.Stdout = *option.stdout
	}
	if option.stderr != nil {
		cmd.Stderr = *option.stderr
	}
	return cmd
}
