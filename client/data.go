package main

import (
	"github.com/andlabs/ui"
)

type NameModel struct {
	Names []string
	m     *ui.TableModel
}

func (n *NameModel) AddName(name string) {
	n.Names = append(n.Names, name)
	ui.QueueMain(func() {
		n.m.RowInserted(len(n.Names) - 1) // Update
	})
}

func (n *NameModel) RemoveName(name string) {
	// Find
	ind := -1
	for i, n := range n.Names {
		if n == name {
			ind = i
			break
		}
	}
	if ind == -1 {
		return
	}

	// Remove
	copy(n.Names[ind:], n.Names[ind+1:])
	n.Names = n.Names[:len(n.Names)-1]

	// Remove from table
	ui.QueueMain(func() {
		n.m.RowDeleted(ind)
	})
}

func (n *NameModel) ColumnTypes(m *ui.TableModel) []ui.TableValue {
	return []ui.TableValue{
		ui.TableString(""), // column 0, names
	}
}

func (n *NameModel) NumRows(m *ui.TableModel) int {
	return len(n.Names)
}

func (n *NameModel) CellValue(m *ui.TableModel, row, column int) ui.TableValue {
	return ui.TableString(n.Names[row])
}

func (n *NameModel) SetCellValue(m *ui.TableModel, row, column int, value ui.TableValue) {} // Ignore, since no columns are editable

func (n *NameModel) Clear() {
	oldLen := len(n.Names)
	n.Names = make([]string, 0)
	ui.QueueMain(func() {
		for i := 0; i < oldLen; i++ {
			n.m.RowDeleted(0)
		}
	})
}

type BindsModel struct {
	Binds map[string]string
	Rows  []string
	m     *ui.TableModel
}

func (b *BindsModel) ColumnTypes(m *ui.TableModel) []ui.TableValue {
	return []ui.TableValue{
		ui.TableString(""), // column 0, key
		ui.TableString(""), // column 1, val
	}
}

func (b *BindsModel) NumRows(m *ui.TableModel) int {
	return len(b.Binds)
}

func (b *BindsModel) CellValue(m *ui.TableModel, row, column int) ui.TableValue {
	if column == 0 {
		return ui.TableString(b.Rows[row])
	}
	return ui.TableString(b.Binds[b.Rows[row]])
}

func (b *BindsModel) SetCellValue(m *ui.TableModel, row, column int, value ui.TableValue) {} // Ignore, since no columns are editable

func (b *BindsModel) Clear() {
	oldLen := len(b.Rows)
	b.Rows = make([]string, 0)
	b.Binds = make(map[string]string)
	ui.QueueMain(func() {
		for i := 0; i < oldLen; i++ {
			b.m.RowDeleted(0)
		}
	})
}

func (b *BindsModel) SetBinds(binds map[string]string) {
	rows := make([]string, 0, len(binds))
	for k := range binds {
		rows = append(rows, k)
	}
	b.Rows = rows
	b.Binds = binds

	ui.QueueMain(func() {
		for i := 0; i < len(b.Rows); i++ {
			b.m.RowInserted(i)
		}
	})
}
