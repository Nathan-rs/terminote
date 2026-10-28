package components

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func Header() *tview.TextView {
	// Headers
	header := tview.NewTextView().
		SetTextColor(tcell.NewHexColor(0x00A693)).
		SetTextAlign(tview.AlignCenter).
		SetText("| Tasks manager & Controllers [v1.0] |")
	header.SetBorder(true).SetBorderColor(tcell.NewHexColor(0xD4D4D4))

	return header
}
