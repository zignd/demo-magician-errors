package magician

import (
	"fmt"

	"github.com/zignd/demo-magician-errors/magic"
)

// Show will get the magician to the stage to do the show
func Show() error {
	if err := magic.DoMagic(magic.BentSpoonTrick); err != nil {
		return fmt.Errorf("Failed to perform the first trick: %w", err)
	}

	if err := magic.DoMagic(magic.BentSpoonTrick); err != nil {
		return fmt.Errorf("Failed to perform the second trick: %w", err)
	}

	return nil
}
