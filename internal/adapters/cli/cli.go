package cli

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/fatih/color"
	"github.com/snufkin23/lusay/internal/core/service"
	"github.com/snufkin23/lusay/internal/utils/catsay"
	"github.com/snufkin23/lusay/internal/utils/logger"
)

// App handles the CLI interaction for lusay
type App struct {
	aiSvc  *service.AIService
	logger *logger.Logger
}

// NewApp creates a new instance of the CLI application
func NewApp(aiSvc *service.AIService, logger *logger.Logger) *App {
	return &App{
		aiSvc:  aiSvc,
		logger: logger,
	}
}

// Run starts the interactive REPL loop
func (a *App) Run() {
	a.printBanner()

	scanner := bufio.NewScanner(os.Stdin)
	for {
		color.New(color.FgHiMagenta, color.Bold).Print("\n❯ You: ")
		if !scanner.Scan() {
			break
		}

		userInput := strings.TrimSpace(scanner.Text())
		if userInput == "" {
			continue
		}

		lowerInput := strings.ToLower(userInput)
		if lowerInput == "exit" || lowerInput == "quit" {
			a.printGoodbye()
			break
		}

		if lowerInput == "clear" {
			a.clearScreen()
			continue
		}

		if rand.Float32() < 0.15 {
			a.randomInterruption()
		}

		// Paw-print thinking animation
		stopAnim := make(chan bool, 1)
		go a.animateThinking(stopAnim)

		resp, err := a.aiSvc.GenerateResponse(userInput)
		stopAnim <- true
		time.Sleep(60 * time.Millisecond) // let goroutine flush

		if err != nil {
			if err.Error() == "service.GenerateResponse: harmful or dangerous content detected" {
				fmt.Println()
				a.printHissWarning()
			} else {
				a.logger.Error("failed to generate response", err)
			}
			continue
		}

		catResp := catsay.Format(resp.Text, resp.Mood)
		a.renderCatResponse(catResp)
	}

	if err := scanner.Err(); err != nil {
		a.logger.Fatal("error reading input", err)
	}
}

// printBanner renders the startup banner
func (a *App) printBanner() {
	banner := []string{
		"",
		"  ██╗     ██╗   ██╗ ██████╗ █████╗ ██╗   ██╗",
		"  ██║     ██║   ██║██╔════╝██╔══██╗╚██╗ ██╔╝",
		"  ██║     ██║   ██║╚█████╗ ███████║ ╚████╔╝ ",
		"  ██║     ██║   ██║ ╚═══██╗██╔══██║  ╚██╔╝  ",
		"  ███████╗╚██████╔╝ ██████╔╝██║  ██║   ██║   ",
		"  ╚══════╝ ╚═════╝  ╚═════╝ ╚═╝  ╚═╝   ╚═╝   ",
		"",
	}
	bannerColor := color.New(color.FgYellow, color.Bold)
	for _, line := range banner {
		bannerColor.Println(line)
		time.Sleep(35 * time.Millisecond)
	}

	color.New(color.FgHiBlack, color.Italic).Println("  ~ A philosophical orange cat who really doesn't want to help you ~")
	color.New(color.FgHiBlack).Println("  Type 'exit' or 'quit' to release him. 'clear' to tidy up.\n")
}

// printGoodbye renders a moody farewell
func (a *App) printGoodbye() {
	fmt.Println()
	color.New(color.FgYellow).Println("  /\\_____/\\")
	color.New(color.FgYellow).Println(" (  -   -  )")
	color.New(color.FgYellow).Println("  )  ~~~  (")
	color.New(color.FgYellow).Println(" /  zzzzz  \\")
	color.New(color.FgHiBlack, color.Italic).Println("\n  *stretches all four paws*")
	color.New(color.FgHiMagenta).Println("  🐱 Finally. Back to my nap spot. Don't call me.\n")
}

// printHissWarning renders the hissing danger response
func (a *App) printHissWarning() {
	hissColor := color.New(color.FgRed, color.Bold)
	lines := []string{
		`  ^ /\_/\ ^`,
		` ^| >   < |^`,
		` ^|  VVV  |^`,
		`  \ ~~~~~ /`,
		`  /^ ^^^ ^\`,
	}
	for _, l := range lines {
		hissColor.Println("  " + l)
		time.Sleep(40 * time.Millisecond)
	}
	fmt.Println()
	color.New(color.FgRed).Println("  🐱 *HISS* That question violates the cat-code of ethics. Absolutely not.")
	fmt.Println()
}

// animateThinking shows a paw-print crawl while Lusay "thinks"
func (a *App) animateThinking(stop chan bool) {
	frames := []string{
		"🐾· · · · · · · ·",
		"· 🐾· · · · · · ·",
		"· · 🐾· · · · · ·",
		"· · · 🐾· · · · ·",
		"· · · · 🐾· · · ·",
		"· · · · · 🐾· · ·",
		"· · · · · · 🐾· ·",
		"· · · · · · · 🐾·",
		"· · · · · · · · 🐾",
		"· · · · · · · 🐾·",
		"· · · · · · 🐾· ·",
		"· · · · · 🐾· · ·",
		"· · · · 🐾· · · ·",
		"· · · 🐾· · · · ·",
		"· · 🐾· · · · · ·",
		"· 🐾· · · · · · ·",
	}

	labels := []string{
		"pondering the void",
		"judging your question",
		"consulting the nap gods",
		"reluctantly thinking",
	}
	label := labels[rand.Intn(len(labels))]
	thinkColor := color.New(color.FgHiMagenta)

	i := 0
	for {
		select {
		case <-stop:
			fmt.Print("\r\033[K") // clear line
			return
		default:
			thinkColor.Printf("\r  %s  %s...", frames[i], label)
			i = (i + 1) % len(frames)
			time.Sleep(80 * time.Millisecond)
		}
	}
}

// renderCatResponse handles the sequenced manga-scroll animation
func (a *App) renderCatResponse(catResp catsay.CatResponse) {
	topBorder    := "╔══════════════════════════════════════╗"
	bottomBorder := "╚══════════════════════════════════════╝"
	divider      := "  ·····◈·····◈·····◈·····◈·····◈·····"

	fmt.Println()
	color.New(color.FgHiBlack).Println("  " + topBorder)
	fmt.Println()

	stageColors := []*color.Color{
		color.New(color.FgYellow, color.Bold),    // DAYDREAM (Orange)
		color.New(color.FgYellow, color.Bold),    // MEOW (Orange)
		color.New(color.FgYellow, color.Bold),    // NAP (Orange)
	}

	for i, stage := range catResp.Stages {
		c := stageColors[i%len(stageColors)]

		// Header with padding
		c.Printf("  %s\n", stage.Header)
		color.New(color.FgHiBlack).Println(divider)
		fmt.Print("  ")

		// Typewrite content
		a.dynamicTypewrite(stage.Content)
		fmt.Println()
		fmt.Println()
	}

	color.New(color.FgHiBlack).Println("  " + bottomBorder)

	// Cat pop-in
	a.popCat(catResp.Art, catResp.Mood)
}

// popCat renders the ASCII cat with a snap-in illusion
func (a *App) popCat(art string, mood string) {
	var moodColor *color.Color
	switch strings.ToUpper(mood) {
	case "HAPPY":
		moodColor = color.New(color.FgYellow, color.Bold)
	case "SHOCKED":
		moodColor = color.New(color.FgCyan, color.Bold)
	case "NERD":
		moodColor = color.New(color.FgBlue, color.Bold)
	case "SNEAKY":
		moodColor = color.New(color.FgMagenta, color.Bold)
	case "HISSING":
		moodColor = color.New(color.FgRed, color.Bold)
	default:
		moodColor = color.New(color.FgYellow, color.Bold)
	}

	fmt.Println()

	// Snap-in: print each line with zero delay for "pop" effect
	lines := strings.Split(art, "\n")
	for _, line := range lines {
		moodColor.Printf("    %s\n", line)
	}

	time.Sleep(120 * time.Millisecond)

	fmt.Println()
	color.New(color.FgHiBlack, color.Italic).Println("  ~ Lusay has returned to his nap spot ~")
	fmt.Println()
}

// dynamicTypewrite prints text with variable speeds
func (a *App) dynamicTypewrite(text string) {
	runes := []rune(text)
	for idx, char := range runes {
		fmt.Print(string(char))

		// Check if a long pause ellipsis is forming
		if char == '.' && idx+2 < len(runes) && runes[idx+1] == '.' && runes[idx+2] == '.' {
			time.Sleep(350 * time.Millisecond) // dramatic pause for ...
		} else if strings.ContainsRune(".!?", char) {
			time.Sleep(time.Duration(180+rand.Intn(70)) * time.Millisecond)
		} else if char == ',' || char == ';' {
			time.Sleep(90 * time.Millisecond)
		} else if char == ' ' {
			time.Sleep(22 * time.Millisecond)
		} else {
			time.Sleep(time.Duration(28+rand.Intn(12)) * time.Millisecond)
		}
	}
}

// clearScreen clears the terminal
func (a *App) clearScreen() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
	color.New(color.FgHiBlack, color.Italic).Println("  ~ The screen has been purr-ified ~\n")
}

// randomInterruption adds flavor to the cat persona
func (a *App) randomInterruption() {
	interrupts := []string{
		"🐱 *spots a moth* ...Wait, what was that?!",
		"🐱 *starts grooming paw* One second, human...",
		"🐱 *knocks a pen off the desk* Oops. Anyway...",
		"🐱 *purrs loudly* You're actually asking a decent question for once.",
		"🐱 *stares at wall for 4 seconds* ...Yes. Continue.",
		"🐱 *sits directly on your keyboard* hjkjkhkjh — that wasn't me.",
	}
	idx := rand.Intn(len(interrupts))
	fmt.Println()
	color.New(color.FgYellow, color.Italic).Println("  " + interrupts[idx])
	fmt.Println()
	time.Sleep(600 * time.Millisecond)
}
