package dropdown

import (
	tea "charm.land/bubbletea/v2"
	lipgloss "charm.land/lipgloss/v2"
)

// Custom Dropdown Bubble
type Model struct {
	selected int
	choices []string
	cursor int
	open bool
}
func New() Model {
	return Model{
choices:  []string{"Buy carrots", "Buy celery", "Buy kohlrabi"},
open: true,
	}
}
func (m Model) Init() tea.Cmd {
	return nil
}
 func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
switch msg := msg.(type) {
    case tea.KeyPressMsg:
        switch msg.String() {
 // The "up" and "k" keys move the cursor up
        case "up", "k":
            if m.cursor > 0 {
                m.cursor--
            }
        // The "down" and "j" keys move the cursor down
        case "down", "j":
            if m.cursor < len(m.choices)-1 {
                m.cursor++
            }
        // The "enter" key and the space bar toggle the selected state
        // for the item that the cursor is pointing at.
        case "enter":
			m.selected = m.cursor
        }
    }
	return m,nil
 }
func (m Model) View() string {

 s := "\nPress q to quit.\n"
 if m.open {
	
  for _, choice := range m.choices {
        s += lipgloss.NewStyle().Border(lipgloss.NormalBorder(),true,false).Render(choice)
    }
 } else {
	s += m.choices[m.selected]
 }
    return s
}