package core

func (r *Relation) Load(csvFile string, separator rune) {
	// TODO
}

func (r *Relation) Scan(colList []AttrInfo) Relationer {
	// TODO

	var rs *Relation = new(Relation)
	return rs
}

func (r *Relation) Select(col AttrInfo, comp Comparison, compVal interface{}) Relationer {
	// TODO
	var rs *Relation = new(Relation)
	return rs
}

func (r *Relation) Print() {
	// TODO
}

func (cs *ColumnStore) CreateRelation(tabName string, sig []AttrInfo) Relationer {
	if cs.Relations == nil {
		cs.Relations = make(map[string]Relationer)
	}

	var cols []Column = make([]Column, len(sig))
	for i, s := range sig {
		cols[i].Signature = s
	}
	var rs *Relation = new(Relation)
	rs.Name = tabName
	rs.Columns = cols

	cs.Relations[tabName] = rs
	return rs
}
