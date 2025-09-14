package main

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type model struct {
	width  int
	height int
	weatherData []string
	err error
}

func initialModel() model {
	return model{}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m model) View() string {
	if m.width == 0 || m.height == 0 {
		return "Loading..."
	}

	padding := 2

	// Define styles
	borderStyle := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		Padding(padding)

	titleStyle := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("205"))

	// Create content for each section
	greetingContent := titleStyle.Render("Greeting, quote + ASCII") + "\n\n" +
		"Hello Mateo\nWelcome to your dashboard!"

	countdownsContent := titleStyle.Render("Countdowns") + "\n\n" +
		"This is where the countdown will go"

	weatherContent := titleStyle.Render("Weather") + "\n\n"
	for i, data := range m.weatherData {
		weatherContent += fmt.Sprintf("Weather %d: %s\n", i+1, data)
	}

	fitnessContent := titleStyle.Render("Fitness") + "\n\n" +
		"Fitness tracking data"

	stocksContent := titleStyle.Render("Stocks") + "\n\n" +
		"Stock prices and portfolio"

	funFactContent := titleStyle.Render("Fun fact") + "\n\n" +
		"Did you know...?"

	// Calculate dimensions
	topRowHeight := m.height/4 - padding
	middleRowHeight := (m.height*2)/4 - padding
	bottomRowHeight := m.height/4 - padding

	greetingWidth := (m.width * 3) / 5
	countdownsWidth := (m.width * 2) / 5

	weatherWidth := m.width / 3
	fitnessWidth := m.width * 1 / 3
	stocksWidth := m.width / 3

	// Style each section
	greeting := borderStyle.
		Width(greetingWidth - padding).
		Height(topRowHeight).
		Render(greetingContent)

	countdowns := borderStyle.
		Width(countdownsWidth - padding).
		Height(topRowHeight).
		Render(countdownsContent)

	weather := borderStyle.
		Width(weatherWidth - padding).
		Height(middleRowHeight).
		Render(weatherContent)

	fitness := borderStyle.
		Width(fitnessWidth - padding).
		Height(middleRowHeight).
		Render(fitnessContent)

	stocks := borderStyle.
		Width(stocksWidth - padding).
		Height(middleRowHeight).
		Render(stocksContent)

	funFact := borderStyle.
		Width(m.width - padding).
		Height(bottomRowHeight).
		Render(funFactContent)

	// Build the layout
	topRow := lipgloss.JoinHorizontal(lipgloss.Top, greeting, countdowns)
	middleRow := lipgloss.JoinHorizontal(lipgloss.Top, weather, fitness, stocks)

	dashboard := lipgloss.JoinVertical(lipgloss.Left,
		topRow,
		middleRow,
		funFact,
	)

	return dashboard 
}

func main() {
	p := tea.NewProgram(initialModel())

	FetchWeatherData()

	if _, err := p.Run(); err != nil {
		fmt.Printf("Error running program: %v", err)
	}
}