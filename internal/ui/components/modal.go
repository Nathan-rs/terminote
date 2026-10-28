package components

import "github.com/rivo/tview"

// Wrapper para centralizar qualquer componente na tela com tamanho fixo
func ModalOverlay(primitive tview.Primitive, width, heigth int) tview.Primitive {
	return tview.NewGrid().
		SetColumns(0, width, 0).
		SetRows(0, heigth, 0).
		AddItem(primitive, 1, 1, 1, 1, 0, 0, true)
}

// Cria um modal retornando um formulario
func NewModal(title string, onSave func(), onCancel func()) tview.Primitive {
	form := tview.NewForm()

	return ModalOverlay(form, 55, 11)
}
