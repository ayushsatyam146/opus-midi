package repl

import (
	"fmt"
	"os"

	"github.com/ayushsatyam146/opus-midi/store"
	"gitlab.com/gomidi/midi"
)

func REPL(in midi.In, ActiveNotes map[int64]store.Note) {
	input := ""
    for {
        
        fmt.Scan(&input)
        if(input == "exit") {
            err := in.StopListening()
            fmt.Println(err)
            fmt.Printf("closing MIDI Port %v.....\n", in)
            os.Exit(0)
        } else if(input == "check") {
            err := in.StopListening()
            fmt.Println(err)
            // check ActiveNotes here for chord detection and all
            fmt.Printf("closing MIDI Port %v.....\n", in)
            os.Exit(0)
        }
        fmt.Println("You entered:", input)
    }
}

