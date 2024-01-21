package store

import (
	"regexp"
	"sort"
)

var ChordsMap = map[string][]string{
	"Cmaj": {"C", "E", "G"},
	"Cmin": {"C", "D#", "G"},
	"C7":   {"C", "E", "G", "Bb"},
	"Cm7":  {"C", "D#", "G", "Bb"},
	"Cmaj7": {"C", "E", "G", "B"},
	"CmMaj7": {"C", "D#", "G", "B"},
	"Cdim": {"C", "D#", "F#"},
	"Caug": {"C", "E", "G#"},
	"Cm7b5": {"C", "D#", "F#", "A"},

	"Dmaj":   {"D", "F#", "A"},
	"Dmin":   {"D", "F", "A"},
	"D7":     {"D", "F#", "A", "C"},
	"Dm7":    {"D", "F", "A", "C"},
	"Dmaj7":  {"D", "F#", "A", "C#"},
	"DmMaj7": {"D", "F", "A", "C#"},
	"Ddim":   {"D", "F", "G#"},
	"Daug":   {"D", "F#", "A#"},
	"Dm7b5":  {"D", "F", "G#", "C"},

	"Emaj":   {"E", "G#", "B"},
	"Emin":   {"E", "G", "B"},
	"E7":     {"E", "G#", "B", "D"},
	"Em7":    {"E", "G", "B", "D"},
	"Emaj7":  {"E", "G#", "B", "D#"},
	"EmMaj7": {"E", "G", "B", "D#"},
	"Edim":   {"E", "G", "A#"},
	"Eaug":   {"E", "G#", "B#"},
	"Em7b5":  {"E", "G", "A#", "D"},

	"Fmaj":   {"F", "A", "C"},
	"Fmin":   {"F", "G#", "C"},
	"F7":     {"F", "A", "C", "Eb"},
	"Fm7":    {"F", "G#", "C", "Eb"},
	"Fmaj7":  {"F", "A", "C", "E"},
	"FmMaj7": {"F", "G#", "C", "E"},
	"Fdim":   {"F", "G#", "B"},
	"Faug":   {"F", "A", "C#"},
	"Fm7b5":  {"F", "G#", "B", "Eb"},

	"Gmaj":   {"G", "B", "D"},
	"Gmin":   {"G", "A#", "D"},
	"G7":     {"G", "B", "D", "F"},
	"Gm7":    {"G", "A#", "D", "F"},
	"Gmaj7":  {"G", "B", "D", "F#"},
	"GmMaj7": {"G", "A#", "D", "F#"},
	"Gdim":   {"G", "A#", "C#"},
	"Gaug":   {"G", "B", "D#"},
	"Gm7b5":  {"G", "A#", "C#", "F"},

	"Amaj":   {"A", "C#", "E"},
	"Amin":   {"A", "C", "E"},
	"A7":     {"A", "C#", "E", "G"},
	"Am7":    {"A", "C", "E", "G"},
	"Amaj7":  {"A", "C#", "E", "G#"},
	"AmMaj7": {"A", "C", "E", "G#"},
	"Adim":   {"A", "C", "D#"},
	"Aaug":   {"A", "C#", "E#"},
	"Am7b5":  {"A", "C", "D#", "G"},

	"Bmaj":   {"B", "D#", "F#"},
	"Bmin":   {"B", "D", "F#"},
	"B7":     {"B", "D#", "F#", "A"},
	"Bm7":    {"B", "D", "F#", "A"},
	"Bmaj7":  {"B", "D#", "F#", "A#"},
	"BmMaj7": {"B", "D", "F#", "A#"},
	"Bdim":   {"B", "D", "E"},
	"Baug":   {"B", "D#", "F##"},
	"Bm7b5":  {"B", "D", "E", "A"},

	"Dbmaj":   {"Db", "F", "Ab"},
	"Dbmin":   {"Db", "E", "Ab"},
	"Db7":     {"Db", "F", "Ab", "Cb"},
	"Dbm7":    {"Db", "E", "Ab", "Cb"},
	"Dbmaj7":  {"Db", "F", "Ab", "C"},
	"DbmMaj7": {"Db", "E", "Ab", "C"},
	"Dbdim":   {"Db", "E", "G"},
	"Dbaug":   {"Db", "F", "A"},
	"Dbm7b5":  {"Db", "E", "G", "Cb"},

	"Ebmaj":   {"Eb", "G", "Bb"},
	"Ebmin":   {"Eb", "Gb", "Bb"},
	"Eb7":     {"Eb", "G", "Bb", "Db"},
	"Ebm7":    {"Eb", "Gb", "Bb", "Db"},
	"Ebmaj7":  {"Eb", "G", "Bb", "D"},
	"EbmMaj7": {"Eb", "Gb", "Bb", "D"},
	"Ebdim":   {"Eb", "Gb", "A"},
	"Ebaug":   {"Eb", "G", "B"},
	"Ebm7b5":  {"Eb", "Gb", "A", "Db"},

	"F#maj":   {"F#", "A#", "C#"},
	"F#min":   {"F#", "A", "C#"},
	"F#7":     {"F#", "A#", "C#", "E"},
	"F#m7":    {"F#", "A", "C#", "E"},
	"F#maj7":  {"F#", "A#", "C#", "F"},
	"F#mMaj7": {"F#", "A", "C#", "F"},
	"F#dim":   {"F#", "A", "C"},
	"F#aug":   {"F#", "A#", "C##"},
	"F#m7b5":  {"F#", "A", "C", "E"},

	"Gbmaj":   {"Gb", "Bb", "Db"},	
	"Gbmin":   {"Gb", "A", "Db"},
	"Gb7":     {"Gb", "Bb", "Db", "Fb"},
	"Gbm7":    {"Gb", "A", "Db", "Fb"},
	"Gbmaj7":  {"Gb", "Bb", "Db", "F"},
	"GbmMaj7": {"Gb", "A", "Db", "F"},
	"Gbdim":   {"Gb", "A", "C"},
	"Gbaug":   {"Gb", "Bb", "D"},
	"Gbm7b5":  {"Gb", "A", "C", "Fb"},

	"Abmaj":   {"Ab", "C", "Eb"},
	"Abmin":   {"Ab", "B", "Eb"},
	"Ab7":     {"Ab", "C", "Eb", "Gb"},
	"Abm7":    {"Ab", "B", "Eb", "Gb"},
	"Abmaj7":  {"Ab", "C", "Eb", "G"},
	"AbmMaj7": {"Ab", "B", "Eb", "G"},
	"Abdim":   {"Ab", "B", "D"},
	"Aug":   {"Ab", "C", "E"},
	"Abm7b5":  {"Ab", "B", "D", "Gb"},

	"Bbmaj":   {"Bb", "D", "F"},
	"Bbmin":   {"Bb", "Db", "F"},
	"Bb7":     {"Bb", "D", "F", "Ab"},
	"Bbm7":    {"Bb", "Db", "F", "Ab"},
	"Bbmaj7":  {"Bb", "D", "F", "A"},
	"BbmMaj7": {"Bb", "Db", "F", "A"},
	"Bbdim":   {"Bb", "Db", "E"},
	"Bbaug":   {"Bb", "D", "F#"},
	"Bbm7b5":  {"Bb", "Db", "E", "Ab"},

	"C#maj":   {"C#", "F", "G#"},
	"C#min":   {"C#", "E", "G#"},
	"C#7":     {"C#", "F", "G#", "B"},
	"C#m7":    {"C#", "E", "G#", "B"},
	"C#maj7":  {"C#", "F", "G#", "B#"},
	"C#mMaj7": {"C#", "E", "G#", "B#"},
	"C#dim":   {"C#", "E", "G"},
	"C#aug":   {"C#", "F", "A"},
	"C#m7b5":  {"C#", "E", "G", "B"},
}

func removeNumerics(input string) string {
	numericRegex := regexp.MustCompile("[0-9]")
	result := numericRegex.ReplaceAllString(input, "")
	return result
}

func FetchCord(notes []string) string {

	for i, note := range notes {
		notes[i] = removeNumerics(note)
	}
	
	sort.Strings(notes)

	for key, value := range ChordsMap {
		
		if len(notes) == len(value) {
			sort.Strings(value)
			match := true

			for i, note := range notes {
				if note != value[i] {
					match = false
					break
				}
			}
			if match {
				return key
			}
		}
	}
	return ""
}