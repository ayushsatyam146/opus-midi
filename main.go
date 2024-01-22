package main

import (
	"fmt"

	"github.com/ayushsatyam146/opus-midi/repl"
	"github.com/ayushsatyam146/opus-midi/store"
	"github.com/ayushsatyam146/opus-midi/utils"
	"gitlab.com/gomidi/rtmididrv"
)

func main() {
	drv, err := rtmididrv.New()
	utils.Must(err)

	defer drv.Close()

	ins, err := drv.Ins()
	utils.Must(err)

	// takes the first MIDI input channel
	in := ins[0]

	fmt.Printf("opening MIDI Port %v\n", in)
	utils.Must(in.Open())

	defer in.Close()

	ActiveNotes := make(map[int64][]store.Note)

	utils.MIDIReader(in, ActiveNotes)

	repl.REPL(in, ActiveNotes)
}
