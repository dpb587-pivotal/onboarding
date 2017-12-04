package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	yaml "gopkg.in/yaml.v2"
)

var slugger = regexp.MustCompilePOSIX(`[^a-z0-9\-]`)

type Story struct {
	Priority    int
	Type        string
	Title       string
	Labels      []string
	Description string
	Tasks       []string
}

type StoryFrontmatter struct {
	Title  string   `yaml:"title"`
	Labels []string `yaml:"labels,omitempty"`
	Tasks  []string `yaml:"tasks,omitempty"`
	Type   string   `yaml:"type"`
	Weight int      `yaml:"weight"`
}

func (s Story) DumpMarkdown() []byte {
	frontmatter, err := yaml.Marshal(StoryFrontmatter{
		Title:  s.Title,
		Labels: s.Labels,
		Tasks:  s.Tasks,
		Type:   fmt.Sprintf("story-%s", s.Type),
		Weight: s.Priority,
	})
	if err != nil {
		panic(err)
	}

	return []byte(fmt.Sprintf("---\n%s\n---\n\n%s", strings.TrimSpace(string(frontmatter)), s.Description))
}

func main() {
	var dir = filepath.Join("/", "Users", "dpb587", "Projects", "pivotal-cf", "onboarding")

	files, err := filepath.Glob(filepath.Join(dir, "*.prolific"))
	if err != nil {
		panic(err)
	}

	for _, file := range files {
		base := strings.TrimSuffix(filepath.Base(file), ".prolific")
		dashbash := strings.Replace(base, "_", "-", -1)
		filedir := filepath.Join(dir, "content", dashbash, "stories")

		err = os.MkdirAll(filedir, 0755)
		if err != nil {
			panic(err)
		}

		data, err := ioutil.ReadFile(file)
		if err != nil {
			panic(err)
		}

		datastr := fmt.Sprintf("%s", data)

		stories := strings.Split(strings.TrimSpace(datastr), "\n---\n")

		for storyIdx, storyStr := range stories {
			storyStr = strings.TrimSpace(storyStr)
			if storyStr == "" {
				continue
			}

			story := Story{
				Priority: storyIdx + 1,
			}

			if strings.HasPrefix(storyStr, "[RELEASE] ") {
				story.Type = "release"
				storyStr = storyStr[10:]
			} else {
				story.Type = "feature"
			}

			storySplit := strings.SplitN(storyStr, "\n", 2)
			story.Title = storySplit[0]

			storyLines := strings.Split(storyStr, "\n")
			for _, storyLine := range storyLines {
				if strings.HasPrefix(storyLine, "L: ") {
					story.Labels = strings.Split(storyLine[3:], ", ")
				} else if strings.HasPrefix(storyLine, "- [ ] ") {
					story.Tasks = append(story.Tasks, storyLine[6:])
				} else {
					story.Description = strings.TrimPrefix(strings.Join([]string{story.Description, storyLine}, "\n"), "\n")
				}
			}

			storyslug := strings.Trim(string(slugger.ReplaceAll([]byte(strings.Replace(strings.ToLower(story.Title), " ", "-", -1)), nil)), "-")
			storymd := filepath.Join(filedir, fmt.Sprintf("%s.md", storyslug))

			err := ioutil.WriteFile(storymd, story.DumpMarkdown(), 0644)
			if err != nil {
				panic(err)
			}
		}
	}
}
