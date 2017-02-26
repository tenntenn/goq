package option

import "regexp"

type Parttern struct {
	Value  *string
	Regexp *regexp.Regexp
}

func (p *Parttern) Match(s string) bool {
	switch {
	case p.Value != nil:
		return *p.Value == s
	case p.Regexp != nil:
		return p.Regexp.MatchString(s)
	}
	return false
}
