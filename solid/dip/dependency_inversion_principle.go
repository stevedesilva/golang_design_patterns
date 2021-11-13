package dip

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
	from *Person
	relationshipType RelationshipType
	to *Person
}

type Relationships struct {
	relations []*Info
}