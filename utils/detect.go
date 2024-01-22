package utils

import (
	"sort"

	"github.com/ayushsatyam146/opus-midi/store"
)

func ParseNotes(ActiveNotes map[int64][]store.Note) [][]store.Note {
	totalNotes := 0
	for _, notes := range ActiveNotes {
		totalNotes += len(notes)
	}
	song := make([]store.Note, totalNotes)
	for _, notes := range ActiveNotes {
		for _, note := range notes {
			song = append(song, note)
		}
	}

	sort.Slice(song, func(i, j int) bool {
		return song[i].TimeStamp < song[j].TimeStamp
	})

	NoteGroups := [][]store.Note{}
	LocalGroup := []store.Note{}
	for i, note := range song {
		if i == 0 {
			LocalGroup = append(LocalGroup, note)
		} else {
			if note.TimeStamp-song[i-1].TimeStamp < 100 {
				LocalGroup = append(LocalGroup, note)
			} else {
				NoteGroups = append(NoteGroups, LocalGroup)
				LocalGroup = []store.Note{}
				LocalGroup = append(LocalGroup, note)
			}
		}
	}
	NoteGroups = append(NoteGroups, LocalGroup)
	return NoteGroups
}
