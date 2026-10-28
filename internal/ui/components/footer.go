package components

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func Footer() *tview.TextView {
	footer := tview.NewTextView().
		SetTextColor(tcell.NewHexColor(0x00A693)).
		SetTextAlign(tview.AlignLeft).
		SetText("| [N] Nova Task | [Enter] Editar] | [Q] Sair |")
	footer.SetBorder(true).SetBorderColor(tcell.NewHexColor(0xD4D4D4))

	return footer
}
