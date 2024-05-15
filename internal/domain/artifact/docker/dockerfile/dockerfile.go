package dockerfile

import (
	"fmt"
	"strings"
)

// Data of yaml manifest (`docker.service.[name].build.stages.[...]`)
type Stage struct {
	As            *string
	Image         string
	Envinronments map[string]string `yaml:"env"`
	Copies        []Copy            `yaml:"copy"`
	Runs          []string          `yaml:"run"`
	Cmds          []string          `yaml:"cmd"`
	EntryPoint    []string          `yaml:"entrypoint"`
}

type Copy struct {
	From *string
	Src  []string
	To   string
}

func New(stages []Stage) []byte {
	dockerfile := make([]string, 0, lines(stages))
	const MARGIN = ""
	for i, stage := range stages {
		if i > 0 {
			dockerfile = append(dockerfile, MARGIN)
		}
		dockerfile = append(dockerfile, from(stage.Image, stage.As))
		dockerfile = append(dockerfile, envs(stage.Envinronments)...)
		dockerfile = append(dockerfile, copies(stage.Copies)...)
		dockerfile = append(dockerfile, runs(stage.Runs)...)
		dockerfile = append(dockerfile, cmds(stage.Cmds)...)
		if len(stage.EntryPoint) != 0 {
			dockerfile = append(dockerfile, entrypoint(stage.EntryPoint))
		}
	}
	return []byte(strings.Join(dockerfile, "\n"))
}

func lines(stages []Stage) (count int) {
	count = 0
	count += len(stages)*2 - 1 /** FROM and margin lines */
	for _, stage := range stages {
		count += len(stage.Envinronments)
		count += len(stage.Runs)
		count += len(stage.Cmds)
		if len(stage.EntryPoint) != 0 {
			count += 1
		}
	}
	return count
}

func from(image string, as *string) string {
	if as == nil {
		return fmt.Sprintf("FROM %s", image)
	}
	return fmt.Sprintf("FROM %s AS %s", image, *as)
}

func envs(envs map[string]string) []string {
	res := make([]string, 0, len(envs))
	for name, value := range envs {
		res = append(res, fmt.Sprintf("ENV %s=%s", name, value))
	}
	return res
}

func copies(copies []Copy) []string {
	res := make([]string, 0, len(copies))
	for _, c := range copies {
		if c.From != nil {
			res = append(res, fmt.Sprintf("COPY --from=%s %s %s", *c.From, strings.Join(c.Src, " "), c.To))
		} else {
			res = append(res, fmt.Sprintf("COPY %s %s", strings.Join(c.Src, " "), c.To))
		}
	}
	return res
}

func runs(runs []string) []string {
	res := make([]string, 0, len(runs))
	for _, run := range runs {
		res = append(res, fmt.Sprintf("RUN %s", run))
	}
	return res
}

func cmds(cmds []string) []string {
	res := make([]string, 0, len(cmds))
	for _, cmd := range cmds {
		res = append(res, fmt.Sprintf("CMD %s", cmd))
	}
	return res
}

func entrypoint(ep []string) string {
	if len(ep) == 0 {
		return ""
	}
	res := make([]string, 0, len(ep)+2)
	res = append(res, "ENTRYPOINT [")
	for _, v := range ep {
		res = append(res, fmt.Sprintf("\"%s\"", v))
	}
	res = append(res, "]")
	return strings.Join(res, " ")
}
