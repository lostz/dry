package termui

import (
	"github.com/gizak/termui"
	"github.com/moncho/dry/ui"
)

//Table defines common behaviour for table widgets
type Table interface {
	ColumnWidths() []int
}

//TableHeader is a table header widget
type TableHeader struct {
	X, Y              int
	Height, Width     int
	Columns           []*termui.Par
	ColumnSpacing     int
	fixedWidthColumns []*termui.Par
	varWidthColumns   []*termui.Par
	Theme             *ui.ColorTheme
	columnWidths      []int
}

//NewHeader creates a header of height 1 that uses the given Theme
func NewHeader(Theme *ui.ColorTheme) *TableHeader {
	return &TableHeader{Height: 1, Theme: Theme}
}

//GetHeight return this header's height
func (th *TableHeader) GetHeight() int {
	return th.Height
}

//SetWidth sets the width of this header
func (th *TableHeader) SetWidth(w int) {
	x := th.X
	th.Width = w
	//Set width on each non-fixed width column
	iw := th.calcColumnWidth()
	for _, col := range th.Columns {
		col.SetX(x)
		if col.Width == -1 {
			col.SetWidth(iw)
		}
		x += col.Width + th.ColumnSpacing
		th.columnWidths = append(th.columnWidths, col.Width)
	}
}

//SetX sets the X position of this header
func (th *TableHeader) SetX(x int) {
	th.X = x
}

//SetY sets the Y position of this header
func (th *TableHeader) SetY(y int) {
	for _, p := range th.Columns {
		p.SetY(y)
	}
	th.Y = y
}

//Buffer returns the content of this header as a buffer
func (th *TableHeader) Buffer() termui.Buffer {
	buf := termui.NewBuffer()
	for _, p := range th.Columns {
		buf.Merge(p.Buffer())
	}
	return buf
}

//AddColumn adds a column to this header
func (th *TableHeader) AddColumn(s string) {
	p := newHeaderColumn(s, th)
	th.varWidthColumns = append(th.varWidthColumns, p)
	th.Columns = append(th.Columns, p)
}

//AddFixedWidthColumn adds a column to this header with a fixed width
func (th *TableHeader) AddFixedWidthColumn(s string, width int) {
	p := newHeaderColumn(s, th)
	p.Width = width
	th.fixedWidthColumns = append(th.fixedWidthColumns, p)
	th.Columns = append(th.Columns, p)

}

//CalcColumnWidth calculates the column width for non-fixed width
//columns on this header
func (th *TableHeader) calcColumnWidth() int {
	fixedWidthColumnsSpacing := 0
	for _, column := range th.fixedWidthColumns {
		fixedWidthColumnsSpacing += column.Width
	}
	colCount := len(th.varWidthColumns)
	spacing := th.ColumnSpacing*colCount + fixedWidthColumnsSpacing
	return (th.Width - spacing) / colCount
}

//ColumnCount returns the number of columns on this header
func (th *TableHeader) ColumnCount() int {
	return len(th.Columns)
}

//ColumnWidths returns the width of each column of the table
func (th *TableHeader) ColumnWidths() []int {
	return th.columnWidths
}

func newHeaderColumn(columnTitle string, th *TableHeader) *termui.Par {
	p := termui.NewPar(columnTitle)
	p.Height = th.Height
	p.Border = false
	p.Bg = termui.Attribute(th.Theme.Bg)
	p.TextBgColor = termui.Attribute(th.Theme.Bg)
	p.TextFgColor = termui.ColorWhite
	p.Width = -1
	return p
}
