package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"charm.land/bubbles/v2/help"
	"charm.land/bubbles/v2/key"
	"charm.land/bubbles/v2/textinput"
	"charm.land/bubbles/v2/viewport"
	tea "charm.land/bubbletea/v2"
	lipgloss "charm.land/lipgloss/v2"

	dropdown "github.com/aung-mkhant/matcha/dropdown"
)


type model struct {
	keys   keyMap
	help     help.Model
	client *http.Client
	textInput textinput.Model
	lastKey string
	err error
	respPane viewport.Model // View API response
	ready bool
	methodInput dropdown.Model // Dropdown for HTTP methods
}

func initialModel() model {
	client := &http.Client{}
	ti := textinput.New()
	ti.Placeholder = "https://example.com"
	ti.SetVirtualCursor(true)
	ti.Prompt = ""
	ti.Focus()
	ti.CharLimit = 156
	ti.SetWidth(20)
	return model{
		keys: keys,
		help: help.New(),
		client: client,
		textInput: ti,
		methodInput: dropdown.New(),
	}

}
func (m model) Init() tea.Cmd {
	return textinput.Blink
	}
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
		m.respPane.SetContent("Welcome to the API client!\n\n" + "Type a URL and press Enter to send a GET request.\n\n" )	
			m.ready = true
		} else {
			m.respPane.SetWidth(msg.Width)
			m.respPane.SetHeight(msg.Height)
		}
	// Error handler
	case errMsg: 
		m.err = msg.err
	case respMsg:
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
		case msg.String() == "R":
			return m, sendQuery(m.client, m.textInput.Value())
		}
	}
	m.textInput, cmd = m.textInput.Update(msg)

	return m, cmd
}
func (m model) View() tea.View {
	 var v tea.View
	v.MouseMode = tea.MouseModeCellMotion
	if !m.ready {
		v.SetContent("\n  Initializing...")
	} else {
		v.SetContent(fmt.Sprintf(
			"%s\n%s\n%s\n%s\n%s",
			 m.headerView(),
		m.methodInput.View(), 
		m.textInput.View(),
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


















func main() {
if os.Getenv("HELP_DEBUG") != "" {
		f, err := tea.LogToFile("debug.log", "help")
		if err != nil {
			fmt.Println("Couldn't open a file for logging:", err)
			os.Exit(1)
		}
		defer f.Close() // nolint:errcheck
	}
    if _, err := tea.NewProgram(initialModel()).Run(); err != nil {
        fmt.Printf("Could not start program: %v", err)
        os.Exit(1)
    }}
// Keybindings and Help Menu
type keyMap struct {
	Up    key.Binding
	Down  key.Binding
	Left  key.Binding
	Right key.Binding
	Help  key.Binding
	Quit  key.Binding
}
func (k keyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.Help, k.Quit}
}
func (k keyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Up, k.Down, k.Left, k.Right}, // first column
		{k.Help, k.Quit},                // second column
	}
}
var keys = keyMap{
	Up: key.NewBinding(
		key.WithKeys("up", "k"),
		key.WithHelp("↑/k", "move up"),
	),
	Down: key.NewBinding(
		key.WithKeys("down", "j"),
		key.WithHelp("↓/j", "move down"),
	),
	Left: key.NewBinding(
		key.WithKeys("left", "h"),
		key.WithHelp("←/h", "move left"),
	),
	Right: key.NewBinding(
		key.WithKeys("right", "l"),
		key.WithHelp("→/l", "move right"),
	),
	Help: key.NewBinding(
		key.WithKeys("?"),
		key.WithHelp("?", "toggle help"),
	),
	Quit: key.NewBinding(
		key.WithKeys("esc", "ctrl+c"),
		key.WithHelp("ctrl+c", "quit"),
	),
}