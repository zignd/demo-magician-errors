package main

import (
	"errors"
	"fmt"

	"github.com/zignd/demo-magician-errors/magic"
	"github.com/zignd/demo-magician-errors/magician"
)

func main() {
	if err := magician.Show(); err != nil {
		fmt.Println("An error occurred during the show:")

		fmt.Println("\tWas it during the BentSpoonTrick?")
		fmt.Println("\t\terrors.Is:", errors.Is(err, magic.ErrBentSpoonTrick))
		var errBentSpoonTrick *magic.BentSpoonTrickError
		fmt.Println("\t\terrors.As:", errors.As(err, &errBentSpoonTrick))

		fmt.Println("\tWas it during the WringerTrick?")
		fmt.Println("\t\terrors.Is:", errors.Is(err, magic.ErrWringerTrick))
		var errWringerTrick *magic.WringerTrickError
		fmt.Println("\t\terrors.As:", errors.As(err, &errWringerTrick))

		fmt.Println("Here's the message returned:")
		fmt.Printf("\t%v\n", err)
		return
	}
	fmt.Println("There's no way for this show to be a success, sorry for the pessimism")
}
