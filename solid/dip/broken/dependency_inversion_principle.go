package dip

import "fmt"

/*
  Higher level modules (HLM) should not depend on lower level modules (LLM)
  Both should depend on abstractions
*/

type Person struct {
	name string
	// other useful stuff
}

type RelationshipType int

const (
	Parent RelationshipType = iota
	Child
	Sibling
)

/*
  Info
  e.g. from(Steve) Parent(relationship) to(Ben)
*/
type Info struct {
	from             *Person
	relationshipType RelationshipType
	to               *Person
}

// Relationships Low level module - could be a DB or some cloud storage
type Relationships struct {
	relations []*Info
}

func (r *Relationships) AddParentAndChild(parent, child *Person) {
	r.relations = append(r.relations, &Info{
		from: parent, relationshipType: Parent, to: child,
	})
	r.relations = append(r.relations, &Info{
		from: child, relationshipType: Child, to: parent,
	})
}

// Research is a HLM
type Research struct {
	// breaks DIP
	relationship *Relationships
}

func (r *Research) Investigate(name string) []string {
	// breaks DIP: uses internals (e.g. what happens if storage for relationships change?)
	relations := r.relationship.relations
	res := make([]string, 0, len(relations))
	for _, r := range r.relationship.relations {
		if name == r.from.name && Parent == r.relationshipType {
			res = append(res, fmt.Sprintf("%s has a child called %s", name, r.to.name))
		}
	}
	return res
}
