/*
   Copyright (C) 2016 cacahuatl < cacahuatl at autistici dot org >
   This program is free software: you can redistribute it and/or modify
   it under the terms of the GNU General Public License as published by
   the Free Software Foundation, either version 3 of the License, or
   (at your option) any later version.
   This program is distributed in the hope that it will be useful,
   but WITHOUT ANY WARRANTY; without even the implied warranty of
   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
   GNU General Public License for more details.
   You should have received a copy of the GNU General Public License
   along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/
package search

import (
	"bufio"
	"errors"
	"log"
	"os"
	"strings"
)

type Search struct {
	w bool
	n [32]*Search
}

var o = map[rune]int{
	'a': 0, 'b': 1, 'c': 2, 'd': 3, 'e': 4, 'f': 5, 'g': 6, 'h': 7, 'i': 8,
	'j': 9, 'k': 10, 'l': 11, 'm': 12, 'n': 13, 'o': 14, 'p': 15, 'q': 16, 'r': 17,
	's': 18, 't': 19, 'u': 20, 'v': 21, 'w': 22, 'x': 23, 'y': 24, 'z': 25, '2': 26,
	'3': 27, '4': 28, '5': 29, '6': 30, '7': 31,
}

func onion_only(r rune) rune {
	if _, ok := o[r]; ok {
		return r
	} else {
		return -1
	}
}

func NewSearch() *Search {
	return &Search{w: false}
}

func (s *Search) Populate(filename string) error {
	var words int = 0
	f, e := os.Open(filename)
	if e != nil {
		return e
	}
	defer f.Close()
	scan := bufio.NewScanner(f)
	for scan.Scan() {
		w := strings.TrimSpace(scan.Text())
		if len(w) < 1 || w != strings.Map(onion_only, w) {
			continue
		}
		var t *Search = s
		for _, r := range w {
			i := o[r]
			if t.n[i] == nil {
				t.n[i] = &Search{
					w: false,
				}
			}
			t = t.n[i]
		}
		t.w = true
		words++
	}
	log.Print("added ", words, " words")
	if words < 1 {
		return errors.New("empty word list ", filename)
	}
	return nil
}

func (s *Search) Search(onion string, full bool) bool {
	var t *Search = s
	for k, r := range onion {
		i, ok := o[r]
		if !ok {
			return false
		}
		if !full && t.w {
			return true
		}
		if full && t.w && s.Search(string([]byte(onion)[k:]), full) {
			return true
		}
		if t.n[i] == nil {
			return false
		}
		t = t.n[i]
	}
	return true
}
