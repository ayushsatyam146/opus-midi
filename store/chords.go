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

func GetChord(notes []Note) string {
	SortByKeys(notes)
	chordName := ""
	diff1 := notes[1].Key - notes[0].Key
	diff2 := notes[2].Key - notes[1].Key
	if diff1 == 4 && diff2 == 3 {
		chordName += notes[0].Name + "maj"
	} else if diff1 == 3 && diff2 == 4 {
		chordName += notes[0].Name + "min"
	} else if diff1 == 3 && diff2 == 3 {
		chordName += notes[0].Name + "dim"
	} else if diff1 == 7 && diff2 == 5 {
		chordName += notes[0].Name + "pow"
	} else {
		chordName = ""
	}
	return chordName
}

func GetGap(notes []Note) string {
	SortByKeys(notes)
	chordName := ""
	chordName += notes[0].Name + notes[1].Name
	return chordName
}

func GetOctave(token string) (string, int) {
	octave := 0
	filteredToken := ""
	for _, char := range token {
		if char >= '0' && char <= '9' {
			octave = int(char - '0')
		} else {
			filteredToken += string(char)
		}
	}
	return filteredToken, octave
}

func GetToken(notes []Note) string {
	SortByKeys(notes)
	chord := ""
	octave := 0

	if len(notes) == 3 {
		chord, octave = GetOctave(GetChord(notes))
	} else if len(notes) == 2 {
		chord, octave = GetOctave(GetGap(notes))
	}

	if octave > 2 {
		if chord == "Amin" {
			return "ASTERISK"
		} else if chord == "Gmaj" {
			return "BANG"
		} else if chord == "Cmaj" {
			return "SEMICOLON"
		} else if chord == "Fmaj" {
			return "MINUS"
		} else if chord == "Emin" {
			return "PLUS"
		} else if chord == "Dmin" {
			return "ASSIGN"
		} else if chord == "Bdim" {
			return "SLASH"
		} else if chord == "Cpow" {
			return "LPAREN"
		} else if chord == "Fpow" {
			return "RPAREN"
		} else if chord == "Gpow" {
			return "LEFTBRACE"
		} else if chord == "Apow" {
			return "RIGHTBRACE"
		} else if chord == "Dpow" {
			return "LET"
		} else if chord == "Epow" {
			return "RETURN"
		} else if chord == "CE" {
			return "LESSTHAN"
		} else if chord == "CG" {
			return "GREATERTHAN"
		} else if chord == "FA" {
			return "EQUAL"
		} else if chord == "FC" {
			return "NOTEQUAL"
		} else if chord == "GB" {
			return "COMMA"
		} else if chord == "GC" {
			return "IF"
		} else if chord == "GD" {
			return "ELSE"
		} else {
			return ""
		}
	}
	return ""
}

func ParseChords(NoteGroups [][]Note) {
	fmt.Println("Parsing Chords")
	token := ""
	for i, group := range NoteGroups {
		token = GetToken(group)
		if token != "" {
			NoteGroups[i][0].Name = token
			NoteGroups[i] = NoteGroups[i][:1]
		}
	}
}
