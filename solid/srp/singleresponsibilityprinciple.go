package srp

import (
	"fmt"
	"strings"
)

/*
  SRP: Single responsibility principle
            - separation of concerns
			- A type should have one primary responsibility
            - One reason to change

		Anti pattern - God object - package contain to many concerns)
*/

type Journal struct {
	entries    []string
	entryCount int
}

func (j *Journal) String() string {
	return strings.Join(j.entries, "\n")
}

func (j *Journal) AddEntry(text string) int {
	j.entryCount++
	entry := fmt.Sprintf("%d %s", j.entryCount, text)
	j.entries = append(j.entries, entry)
	return j.entryCount
}

func (j *Journal) RemoveEntry(index int) {
	j.entries = append(j.entries[:index], j.entries[index+1:]...)
}

// Adding Persistence here will break single responsibility principle
//var lineSeparator = "\n"
//func (j *Journal) Save(filename string) {
// // impl
//	_ = ioutil.WriteFile(filename,
//		[]byte(strings.Join(j.entries, lineSeparator)), 0644)
//}
//
//func (j *Journal) LoadFromWeb(url *url.URL) {
// // impl
//}
//
//func (j *Journal) Load(filename string) {
// // impl
//}
