package dropdown

import (
	tea "charm.land/bubbletea/v2"
	lipgloss "charm.land/lipgloss/v2"
)

// Custom Dropdown Bubble
type Model struct {
	Selected int
	Choices []string
	cursor int
	open bool
}
func New() Model {
	return Model{
	}
}
func (m Model) Init() tea.Cmd {
	return nil
}
 func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
switch msg := msg.(type) {
    case tea.KeyPressMsg:
        switch msg.String() {
        case "up", "k":
            if m.cursor > 0 {
                m.cursor--
            }
        case "down", "j":
            if m.cursor < len(m.Choices)-1 {
                m.cursor++
            }
        case "enter":
			m.Selected = m.cursor
        }
    }
	return m,nil
 }
func (m Model) View() string {
    var s string 
 if m.open {
  for _, choice := range m.Choices {
        s += lipgloss.NewStyle().Border(lipgloss.NormalBorder(),true,false).Render(choice)
    }
 } else {
	s += m.Choices[m.Selected]
 }
    return s
}