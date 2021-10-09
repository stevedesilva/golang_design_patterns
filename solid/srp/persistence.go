package srp

import (
	"io/ioutil"
	"net/url"
	"strings"
)

/*
  Cross cutting concern - controlled from one place
*/

type Persistence struct {
	lineSeparator string
}

func (p *Persistence) Save(filename string, j *Journal) {
	// impl
	_ = ioutil.WriteFile(filename,
		[]byte(strings.Join(j.entries, p.lineSeparator)), 0644)
}

func (p *Persistence) LoadFromWeb(url *url.URL) {
	// impl
}

func (p *Persistence) Load(filename string) {
	// impl
}
