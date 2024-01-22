package store

import (
	"fmt"
	"sort"
)

func SortByKeys(notes []Note) {
	sort.Slice(notes, func(i, j int) bool {
		return notes[i].Key < notes[j].Key
	})
}

func isMajor(notes []Note) string {
	SortByKeys(notes)
	chordName := ""
	diff1 := notes[1].Key - notes[0].Key
	diff2 := notes[2].Key - notes[1].Key
	if(diff1 == 4 && diff2 == 3) {
		chordName += notes[0].Name + "maj"
	} 
	return chordName
}

func isMinor(notes []Note) string {
	SortByKeys(notes)
	chordName := ""
	diff1 := notes[1].Key - notes[0].Key
	diff2 := notes[2].Key - notes[1].Key
	if(diff1 == 3 && diff2 == 4) {
		chordName += notes[0].Name + "min"
	} 
	return chordName
}

func ParseChords(NoteGroups [][]Note) {
	fmt.Println("Parsing Chords")
	for i, group := range NoteGroups {
		if(len(group) >= 3) {
			chord := ""
			chord = isMajor(group)
			chord = isMinor(group)

			if(chord != "") {
				NoteGroups[i][0].Name = chord
				NoteGroups[i] = NoteGroups[i][:1]
			}
		} 
	}
}