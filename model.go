package main
import (
	"net/http"
	"maps"
	"slices"

	"charm.land/bubbles/v2/help"
	"charm.land/bubbles/v2/textinput"
	"charm.land/bubbles/v2/viewport"
"charm.land/bubbles/v2/spinner"
	dropdown "github.com/aung-mkhant/matcha/dropdown"
	tea "charm.land/bubbletea/v2"
	lipgloss "charm.land/lipgloss/v2"
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
	spinner spinner.Model
	loading bool
}

func initialModel() model {
	client := &http.Client{}
	// url input
	ti := textinput.New()
	ti.Placeholder = "https://example.com"
	ti.SetVirtualCursor(true)
	ti.Prompt = ""
	ti.Focus()
	ti.CharLimit = 156
	ti.SetWidth(20)
	// spinner
	sp := spinner.New()
	sp.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("69"))
	sp.Spinner = spinner.Points
	// dropdown
	methods := slices.Collect(maps.Values(methodName))
	d := dropdown.New()
	d.Choices = methods
	return model{
		keys: keys,
		help: help.New(),
		client: client,
		textInput: ti,
		methodInput: d,
		spinner: sp,
	}

}
func (m model) Init() tea.Cmd {
	return nil
	}