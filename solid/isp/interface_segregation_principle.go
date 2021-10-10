package isp

type Document struct {
}

// ISP
// Break up interface into smaller unit

//type Machine interface {
//	Print(d Document)
//	Fax((d Document)
//	Scan(d Document)
//}

type Printer interface {
	Print(d Document)
}

type Scanner interface {
	Scan(d Document)
}

type Faxer interface {
	Fax(d Document)
}

// Types only implement interface that they need
type MyPrinter struct {
}

func (m *MyPrinter) Print(d Document) {
	// ok
}

type Photocopier struct {

}

func (p *Photocopier) Scan(d Document) {
	// Ok
}

func (p *Photocopier) Print(d Document) {
	// Ok
}


// Combined interface
type MultiFunctionDevice interface {
	Printer
	Scanner
	// Faxer
}


// decorator
type MultiFunctionMachine struct {

}

func (m *MultiFunctionMachine) Print(d Document) {
	m.Print(d)
}

func (m *MultiFunctionMachine) Scan(d Document) {
	m.Scan(d)
}

