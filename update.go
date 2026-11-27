package main

import (
	"charm.land/bubbles/v2/key"
	"charm.land/bubbles/v2/spinner"
	"charm.land/bubbles/v2/viewport"
	tea "charm.land/bubbletea/v2"
	lipgloss "charm.land/lipgloss/v2"
)

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {	
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.help.SetWidth(msg.Width)
		if !m.ready{
		headerHeight := lipgloss.Height(m.headerView())
		//footerHeight := lipgloss.Height(m.footerView())
			m.respPane = viewport.New(viewport.WithWidth(msg.Width), viewport.WithHeight(20))
			m.respPane.YPosition = headerHeight
			m.ready = true
		} else {
			m.respPane.SetWidth(msg.Width)
			m.respPane.SetHeight(msg.Height)
		}
	// Error handler
	case errMsg: 
		m.loading = false
		m.err = msg.err
		m.respPane.SetContent( msg.err.Error())
	case respMsg:
		m.loading = false
		m.respPane.SetContent( msg.body)
	// Keybindings
	case tea.KeyPressMsg:
		switch {
		case key.Matches(msg, m.keys.Up):
			m.lastKey = "↑"
		case key.Matches(msg, m.keys.Down):
			m.lastKey = "↓"
		case key.Matches(msg, m.keys.Left):
			m.lastKey = "←"
		case key.Matches(msg, m.keys.Right):
			m.lastKey = "→"
		case key.Matches(msg, m.keys.Help):
			m.help.ShowAll = !m.help.ShowAll
		case key.Matches(msg, m.keys.Quit):
			return m, tea.Quit
		case key.Matches(msg, m.keys.Send):
			m.loading = true
			return m, tea.Batch(sendQuery(m.client, m.textInput.Value()), m.spinner.Tick)
		default:
			m.textInput, cmd = m.textInput.Update(msg)
		}
	case spinner.TickMsg:
		if m.loading {
		m.spinner, cmd = m.spinner.Update(msg)
		}
	}
	return m, cmd
}