package src

import (
	"fmt"
	"github.com/gookit/color"
	"strings"
)

// InlineResponse: Print a status inline message
func InlineResponse(designation, msg string, highlight ...string) {
	dColor := color.FgLightGreen.Render
	hColor := color.FgDarkGray.Render

	switch designation {
	case "error":
		dColor = color.FgLightRed.Render
	case "warning":
		dColor = color.FgLightYellow.Render
	}

	var defaultHighlight string
	if len(highlight) > 0 {
		defaultHighlight = hColor(highlight[0])
	}

	fmt.Printf("%s: %v %v \n", dColor(strings.ToUpper(designation)), msg, defaultHighlight)
}

// InlineItems: Do a loop thought items and print
func InlineItems(designation string, items []string) {
	dColor := color.FgDarkGray.Render
	designation = "List of " + designation + ":"
	fmt.Printf("%s \n\n", dColor(designation))
	for _, item := range items {
		fmt.Printf("- %v\n", item)
	}
}
