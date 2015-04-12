package pathresolver

import "strings"

type PathResolver struct {
	path []string
}

func (p *PathResolver) Add(subPath string) {
	p.path = append(p.path, subPath)
}

func (p *PathResolver) Clear() {
	p.path = []string{}
}

func (p *PathResolver) RemoveLast() {
	if len(p.path) > 0 {
		p.path = p.path[:len(p.path)-1]
	}
}

func (p *PathResolver) CurrentPath() string {
	if len(p.path) == 0 {
		return "/"
	}
	return "/" + strings.Join(p.path, "/") + "/"
}
