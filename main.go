package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func IsSep(r rune) bool { return unicode.IsSpace(r) || unicode.IsPunct(r) || unicode.IsSymbol(r) } //from classwork https://go.dev/play/

func words(s string) []string { //from classwork https://go.dev/play/
	var w []string
	var tmp []rune
	wasSep := true
	for _, r := range s {
		if IsSep(r) {
			if !wasSep {
				w = append(w, string(tmp))
				// tmp = tmp[:0]
				tmp = nil
			}
			wasSep = true
		} else {
			wasSep = false
			tmp = append(tmp, r)
		}
	}
	if len(tmp) > 0 {
		w = append(w, string(tmp))
	}
	return w
}

func find(word string, words []string) int {
	for s, w := range words {
		if w == word {
			return s
		}
	}
	return -1 // hint by terminal
}

func wordOccurrences(s string) (ww []string, occurrences []uint) {
	w := words(s)
	for _, word := range w {
		if find(word, ww) == -1 {
			ww = append(ww, word)
			occurrences = append(occurrences, 1)
		} else {
			occurrences[find(word, ww)] += 1
		}
	}
	return ww, occurrences
}

// code done with help

func main() {
	fmt.Print("Enter a string: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	s := scanner.Text()
	ww, occurrences := wordOccurrences(s)
	for i := range ww {
		fmt.Print(ww[i], ": ", occurrences[i], "\n")
	}
}
