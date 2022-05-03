package main

import (
	"ColumnStore/core"
)

func main() {
	var cs = new(core.ColumnStore)

	var atts = []core.AttrInfo{
		{Name: "id", Type: core.INT, Enc: core.NOCOMP},
		{Name: "Name", Type: core.STRING, Enc: core.NOCOMP},
	}

	cs.CreateRelation("students", atts)
	var rs = cs.Relations["students"].Scan(atts[0:2])
	rs.Print()
}
