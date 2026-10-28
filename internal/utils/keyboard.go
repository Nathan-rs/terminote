package utils

import "github.com/gdamore/tcell/v2"

type KeyActions struct {
	New    func()
	Edit   func()
	Quit   func()
	Escape func()
}

func HandleKeyboard(currentPage string, actions KeyActions) func(*tcell.EventKey) *tcell.EventKey {
	return func(event *tcell.EventKey) *tcell.EventKey {

		switch currentPage {

		case PageMain:
			switch event.Key() {
			case tcell.KeyEnter:
				if actions.Edit != nil {
					actions.Edit()
					return nil
				}

			case tcell.KeyRune:
				switch event.Rune() {
				case 'n', 'N':
					if actions.New != nil {
						actions.New()
						return nil
					}

				case 'q', 'Q':
					if actions.Quit != nil {
						actions.Quit()
						return nil
					}
				}
			}

		case PageCreateTask, PageEditTask:
			switch event.Key() {
			case tcell.KeyEscape:
				if actions.Escape != nil {
					actions.Escape()
					return nil
				}
			}
		}

		return event
	}
}
