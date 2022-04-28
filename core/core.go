package core

// TODO: Session 1 - Implement the Relationer and ColumnStorer interface by
// using the given structs. Implement the methods in the core_impl.go file.
// Implement Load, Scan, Select, Print, GetRawData, CreateRelation and GetRelation.

// TODO: Session 2 - Implement HashJoin and Aggregate

// TODO: Session 3 - Parallelisation and Acceleration

type Comparison string

const (
	EQ  Comparison = "=="
	NEQ Comparison = "!="
	LT  Comparison = "<"
	GT  Comparison = ">"
	LE  Comparison = "<="
	GE  Comparison = ">="
)

type Compression int

const (
	NOCOMP Compression = iota
	RLE
	DICT
	FOR
)

type DataTypes int

const (
	INT DataTypes = iota
	FLOAT
	STRING
)

type AttrInfo struct {
	Name string
	Type DataTypes
	Enc  Compression
}

type Column struct {
	Signature AttrInfo
	Data      interface{}
}

type Relation struct {
	Name    string
	Columns []Column
}

type Relationer interface {
	Load(csvFile string, separator rune)
	Scan(colList []AttrInfo) Relationer
	Select(col AttrInfo, comp Comparison, compVal interface{}) Relationer
	Print()
}

type ColumnStore struct {
	Relations map[string]Relationer
}

type ColumnStorer interface {
	CreateRelation(tabName string, sig []AttrInfo) Relationer
	GetRelation(relName string) Relationer
}
