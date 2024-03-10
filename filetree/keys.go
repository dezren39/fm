package filetree

import "github.com/charmbracelet/bubbles/key"

type KeyMap struct {
	Down    key.Binding
	Up      key.Binding
	GoToTop key.Binding
}

func DefaultKeyMap() KeyMap {
	return KeyMap{
		Down:    key.NewBinding(key.WithKeys("j", "down", "ctrl+n"), key.WithHelp("j", "down")),
		Up:      key.NewBinding(key.WithKeys("k", "up", "ctrl+p"), key.WithHelp("k", "up")),
		GoToTop: key.NewBinding(key.WithKeys("g+g"), key.WithHelp("g+g", "top")),
	}
}
