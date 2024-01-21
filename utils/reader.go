package utils

import (
	"fmt"

	"github.com/ayushsatyam146/opus-midi/store"
	"github.com/ayushsatyam146/opus-midi/timer"
	"gitlab.com/gomidi/midi"
	"gitlab.com/gomidi/midi/reader"
	. "gitlab.com/gomidi/midi/midimessage/channel" // (Channel Messages)
)


func PrintNotes(ActiveNotes map[int64]store.Note) {
    output := "["
    for _, note := range ActiveNotes {
				// fmt.Println(key)
        temp := note.Name + "-" + fmt.Sprint(note.TimeStamp) + " "
        output += temp
    }
    output += "]"
    fmt.Println(output)
}

func Must(err error) {
		if err != nil {
			fmt.Println(err)
		}
}

func MIDIReader(in midi.In, ActiveNotes map[int64]store.Note) {
	rd := reader.New(
		reader.NoLogger(),
		// print every message
		reader.Each(func(pos *reader.Position, msg midi.Message) {

			// inspect
			// fmt.Println(msg)

			switch v := msg.(type) {
			case NoteOn:
				currentTime := timer.GetCurrentTimestamp()
				newNote := store.Note{Key: int(v.Key()), TimeStamp: currentTime, Name: store.MidiNoteMap[int(v.Key())], Velocity: int(v.Velocity())}
				ActiveNotes[currentTime] = newNote
				PrintNotes(ActiveNotes)
			case NoteOff:
				// note := ActiveNotes[int(v.Key())]
				// delete(ActiveNotes, int(v.Key()))
				// PrintNotes(ActiveNotes)
			}
		}),
	)
	err := rd.ListenTo(in)
	Must(err)
}