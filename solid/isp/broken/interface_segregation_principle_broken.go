package broken

type Document struct {
}

type Machine interface {
	Print(d Document)
	Fax((d Document)
	Scan(d Document)
}

type MultiFunctionPrinter struct {
}

func (m *MultiFunctionPrinter) Print(d Document) {
	// ok
}

func (m *MultiFunctionPrinter) Fax(d Document) {
	// ok
}

func (m *MultiFunctionPrinter) Scan(d Document) {
	// ok
}

/*
Forced to implement functions that it does not need because of bloated interface
Therefore we need to split the interface
 */
type OldFashionedPrinter struct {

}

func (o *OldFashionedPrinter) Print(d Document) {
	// ok
}
// Deprecated
func (o *OldFashionedPrinter) Fax() (d Document) {
	panic("operation not supported")
}
// Deprecated
func (o *OldFashionedPrinter) Scan(d Document) {
	panic("operation not supported")
}

