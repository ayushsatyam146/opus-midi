package main

import (
	"fmt"

	"time"

	"gitlab.com/gomidi/midi"
	. "gitlab.com/gomidi/midi/midimessage/channel" // (Channel Messages)
	"gitlab.com/gomidi/midi/reader"
	"gitlab.com/gomidi/rtmididrv"
)

// This example reads from the first input port
func main() {
    drv, err := rtmididrv.New()
    must(err)

    // make sure to close the driver at the end
    defer drv.Close()

    ins, err := drv.Ins()
    must(err)

    // takes the first input
    in := ins[0]

    fmt.Printf("opening MIDI Port %v\n", in)
    must(in.Open())

    defer in.Close()

    // to disable logging, pass mid.NoLogger() as option
    rd := reader.New(
        reader.NoLogger(),
        // print every message
        reader.Each(func(pos *reader.Position, msg midi.Message) {

            // inspect
            fmt.Println(msg)

            switch v := msg.(type) {
            case NoteOn:
                fmt.Printf("NoteOn at channel %v: key: %v velocity: %v\n", v.Channel(), v.Key(), v.Velocity())
            case NoteOff:
                fmt.Printf("NoteOff at channel %v: key: %v\n", v.Channel(), v.Key())
            }
        }),
    )

    // listen for MIDI
    err = rd.ListenTo(in)
    must(err)

    time.Sleep(100 * time.Second)
    err = in.StopListening()
    must(err)
    fmt.Printf("closing MIDI Port %v\n", in)
}

func must(err error) {
    if err != nil {
        panic(err.Error())
    }
}