package tui

import (
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"

	"github.com/mistakenelf/fm/statusbar"
)

// Update handles all UI interactions.
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		cmd  tea.Cmd
		cmds []tea.Cmd
	)

	switch msg := msg.(type) {
	case statusMessageTimeoutMsg:
		m.statusMessage = ""
	case tea.WindowSizeMsg:
		halfSize := msg.Width / 2
		height := msg.Height - statusbar.Height

		cmds = append(cmds, m.image.SetSizeCmd(halfSize, height))
		cmds = append(cmds, m.markdown.SetSizeCmd(halfSize, height))

		m.filetree.SetSize(halfSize, height)
		m.secondaryFiletree.SetSize(halfSize, height)
		m.help.SetSize(halfSize, height)
		m.code.SetSize(halfSize, height)
		m.pdf.SetSize(halfSize, height)
		m.statusbar.SetSize(msg.Width)

		return m, tea.Batch(cmds...)
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keyMap.Quit):
			return m, tea.Quit
		case key.Matches(msg, m.keyMap.OpenFile):
			if !m.showTextInput {
				cmds = append(cmds, m.openFileCmd())
			}
		case key.Matches(msg, m.keyMap.ResetState):
			m.state = idleState
			m.showTextInput = false
			m.disableAllViewports()
			m.resetViewports()
			m.filetree.SetDisabled(false)
			m.textinput.Blur()
			m.filetree.CreatingNewDirectory = false
			m.filetree.CreatingNewFile = false

			m.textinput, cmd = m.textinput.Update(msg)
			cmds = append(cmds, cmd)

			cmds = append(cmds, m.filetree.GetDirectoryListingCmd(m.directoryBeforeMove))

			m.textinput.Reset()
		case key.Matches(msg, m.keyMap.MoveDirectoryItem):
			m.directoryBeforeMove = m.filetree.GetSelectedItem().CurrentDirectory
			m.state = showMoveState
			m.filetree.SetDisabled(true)
		case key.Matches(msg, m.keyMap.ShowTextInput):
			if m.activePane == 0 && !m.filetree.CreatingNewDirectory && !m.filetree.CreatingNewFile {
				m.showTextInput = true
				m.textinput.Focus()
				m.disableAllViewports()

				m.textinput, cmd = m.textinput.Update(msg)
				cmds = append(cmds, cmd)

				m.textinput.Reset()
			}
		case key.Matches(msg, m.keyMap.SubmitTextInput):
			if m.filetree.CreatingNewFile {
				cmds = append(cmds, m.filetree.CreateFileCmd(m.textinput.Value()))
			}

			if m.filetree.CreatingNewDirectory {
				cmds = append(cmds, m.filetree.CreateDirectoryCmd(m.textinput.Value()))
			}

			m.resetViewports()
			m.textinput.Blur()
			m.textinput.Reset()
			m.showTextInput = false
			m.activePane = 0
		case key.Matches(msg, m.keyMap.TogglePane):
			if !m.showTextInput {
				m.activePane = (m.activePane + 1) % 2

				if m.activePane == 0 {
					m.filetree.SetDisabled(false)
					m.disableAllViewports()
				} else {
					m.filetree.SetDisabled(true)

					switch m.state {
					case idleState:
						m.disableAllViewports()
						m.help.SetViewportDisabled(false)
					case showCodeState:
						m.disableAllViewports()
						m.code.SetViewportDisabled(false)
					case showImageState:
						m.disableAllViewports()
						m.image.SetViewportDisabled(false)
					case showPdfState:
						m.disableAllViewports()
						m.pdf.SetViewportDisabled(false)
					case showMarkdownState:
						m.disableAllViewports()
						m.markdown.SetViewportDisabled(false)
					}
				}
			}
		case key.Matches(msg, m.keyMap.GotoTop):
			if m.activePane != 0 {
				m.resetViewports()
			}
		case key.Matches(msg, m.keyMap.GotoBottom):
			if m.activePane != 0 {
				m.code.GotoBottom()
				m.pdf.GotoBottom()
				m.markdown.GotoBottom()
				m.help.GotoBottom()
				m.image.GotoBottom()
			}
		}
	}

	if m.filetree.CreatingNewDirectory || m.filetree.CreatingNewFile {
		m.textinput, cmd = m.textinput.Update(msg)
		cmds = append(cmds, cmd)
	}

	m.filetree, cmd = m.filetree.Update(msg)
	cmds = append(cmds, cmd)

	m.secondaryFiletree, cmd = m.secondaryFiletree.Update(msg)
	cmds = append(cmds, cmd)

	m.code, cmd = m.code.Update(msg)
	cmds = append(cmds, cmd)

	m.markdown, cmd = m.markdown.Update(msg)
	cmds = append(cmds, cmd)

	m.image, cmd = m.image.Update(msg)
	cmds = append(cmds, cmd)

	m.pdf, cmd = m.pdf.Update(msg)
	cmds = append(cmds, cmd)

	m.help, cmd = m.help.Update(msg)
	cmds = append(cmds, cmd)

	m.updateStatusBar()

	return m, tea.Batch(cmds...)
}
