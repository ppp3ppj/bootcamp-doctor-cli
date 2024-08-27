package tui

import (
	"fmt"
	"math/rand"
	"os/exec"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/progress"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/ppp3ppj/bootcamp-doctor-cli/internal/resources"
	"github.com/ppp3ppj/bootcamp-doctor-cli/internal/utils"
)

var (
    currentPkgNameStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("211"))
    doneStyle = lipgloss.NewStyle().Margin(1, 2)
    checkMark = lipgloss.NewStyle().Foreground(lipgloss.Color(42)).SetString("✓")
    crossMark = lipgloss.NewStyle().Foreground(lipgloss.Color(42)).SetString("⤫")
)

type InitPromptModel struct {
	packages []resources.PackageDoctor
	index    int
	width    int
	height   int
	spinner  spinner.Model
	progress progress.Model
	done     bool
    doneCount int
}

func (m *InitPromptModel) packageStatus() lipgloss.Style {
    if m.done {
        return checkMark
    } else {
        return crossMark
    }
}

func NewModel() InitPromptModel {
    p := progress.New(
        progress.WithDefaultGradient(),
        progress.WithWidth(40),
        progress.WithoutPercentage(),
    )

    s := spinner.New()
    s.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("63"))
    return InitPromptModel{
        packages: resources.GetPackages(),
        spinner: s,
        progress: p,
    }
}

func (m InitPromptModel) Init() tea.Cmd {
    return tea.Batch(checkPackage(m.packages[m.index]), m.spinner.Tick)
}

func (m InitPromptModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width, m.height = msg.Width, msg.Height
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "esc", "q":
			return m, tea.Quit
		}
	case installedPkgMsg:
		pkg := m.packages[m.index].Name
        // can commend this. this is check last pkg

		if m.index >= len(m.packages)-1 {
			// Everything's been installed. We're done!

            m.done = checkInstalledCommend(m.packages[m.index].VersionCommand)
            if m.done {
                m.doneCount ++
            }
			return m, tea.Sequence(
				//tea.Printf("%s %s", checkMark, pkg), // print the last success message
				tea.Printf("%s %s", m.packageStatus(), pkg), // print the last success message
				tea.Quit,                            // exit the program
			)
		}

        //_ = checkInstalledCommend(m.packages[m.index].VersionCommand)
        m.done = checkInstalledCommend(m.packages[m.index].VersionCommand)
        if m.done {
            m.doneCount ++
        }
		// Update progress bar
		m.index++
		progressCmd := m.progress.SetPercent(float64(m.index) / float64(len(m.packages)))

		return m, tea.Batch(
			progressCmd,
			tea.Printf("%s %s", m.packageStatus(), pkg),     // print success message above our program
			checkPackage(m.packages[m.index]), // download the next package
		)
	case spinner.TickMsg:
		var cmd tea.Cmd
		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd
	case progress.FrameMsg:
		newModel, cmd := m.progress.Update(msg)
		if newModel, ok := newModel.(progress.Model); ok {
			m.progress = newModel
		}
		return m, cmd
	}
	return m, nil
}

func (m InitPromptModel) View() string {
    n := len(m.packages)
    w := lipgloss.Width(fmt.Sprintf("%d", n))

    //if m.done {
    return doneStyle.Render(fmt.Sprintf("Current: %d packages. Installed: %d packages\n", n, m.doneCount))
    //}

    pkgCount := fmt.Sprintf(" %*d/%*d", w, m.index, w, n)

    spin := m.spinner.View() + " "
    prog := m.progress.View()
    cellsAvail := utils.Max(0, m.width - lipgloss.Width(spin+prog+pkgCount))

    pkgName := currentPkgNameStyle.Render(m.packages[m.index].Name)
    info := lipgloss.NewStyle().MaxWidth(cellsAvail).Render("Installing " + pkgName)

    cellsRemaining := utils.Max(0, m.width - lipgloss.Width(spin+info+prog+pkgName))
    gap := strings.Repeat(" ", cellsRemaining)

    return spin + info + gap + prog + pkgCount
}

type installedPkgMsg string

func checkPackage(pkg resources.PackageDoctor) tea.Cmd {
    d := time.Microsecond * time.Duration(rand.Intn(500)) //nolint:gosec
    return tea.Tick(d, func(t time.Time) tea.Msg{
        return installedPkgMsg(pkg.Name)
    })
}

func checkInstalledCommend(command string) bool {
    //cmd, b := exec.Command("go", "version"), new(strings.Builder)
    cmds := strings.Split(command ," ")
    cmd, b := exec.Command(cmds[0], cmds[1]), new(strings.Builder)
    cmd.Stdout = b
    cmd.Run()
    if len(b.String()) > 0 {
        return true
    }
    return false
}
