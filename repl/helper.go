package repl

import (
	"fmt"
	"os"

	"github.com/ayushsatyam146/opus-midi/store"
	"github.com/ayushsatyam146/opus-midi/utils"
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
            utils.Must(err)
            noteGroups := utils.ParseNotes(ActiveNotes)
            utils.ParseChords(noteGroups)
            utils.PrintNoteGroups(noteGroups)
            
            fmt.Printf("closing MIDI Port %v.....\n", in)
            os.Exit(0)
        }
        // fmt.Println("You entered:", input)
    }
}

