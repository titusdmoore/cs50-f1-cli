package cs50f1cli

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbletea"
)

type model struct {
	active_page int
}

func main() {

}

func initialModel() model {
	return model{
		active_page: 0,
	}
}
