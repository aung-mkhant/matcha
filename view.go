package main
import (
	"fmt"
	"strings"

	tea "charm.land/bubbletea/v2"
	lipgloss "charm.land/lipgloss/v2"
)

func (m model) View() tea.View {
	 var v tea.View
	v.MouseMode = tea.MouseModeCellMotion
	if !m.ready {
		v.SetContent("\n  Initializing...")
	} else {
		v.SetContent(fmt.Sprintf(
			"%s\n%s\t%s\t%s\n%s\n%s",
			 m.headerView(),
		m.methodInput.View(), 
		m.textInput.View(),
			m.spinner.View(),
		 m.respPane.View(),
		 m.footerView()))
	}
	return v
}

func (m model) headerView() string {
	title := titleStyle.Render("Matcha API Client")
	line := strings.Repeat("─", max(0, m.respPane.Width()-lipgloss.Width(title)))
return lipgloss.JoinHorizontal(lipgloss.Center, title, line)
}

func (m model) footerView() string {
	helpView := m.help.View(m.keys)
	return   strings.Repeat("\n",2) + helpView
}