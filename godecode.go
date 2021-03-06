package godecode

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

// Transliterate Unicode string to a valid 7-bit ASCII String.
type GoDecode struct {
	path  string
	cache map[rune][]string
}

// Init data path
func (gd *GoDecode) Init(path string) {
	gd.path = path
	gd.cache = make(map[rune][]string)
}
func (gd *GoDecode) getCache(section rune) []string {
	m, ok := gd.cache[section]
	if ok {
		return m
	}
	var path string
	if os.IsPathSeparator('\\') {
		path = "\\"
	} else {
		path = "/"
	}
	f, err := os.OpenFile(fmt.Sprintf("%s%sX%03x", gd.path, path, section), os.O_RDONLY, 0660)
	if err != nil {
		panic(err)
	}
	gd.cache[section] = read(bufio.NewScanner(f))
	return gd.cache[section]
}
func read(scanner *bufio.Scanner) []string {
	var ret []string
	for scanner.Scan() {
		ret = append(ret, scanner.Text())
	}
	if scanner.Err() != nil {
		panic(scanner.Err())
	}
	return ret
}

// Transliterate an Unicode object into an ASCII string.
func (gd *GoDecode) Decode(str string) string {
	var ret []string
	for _, codepoint := range str {
		// Basic ASCII
		if codepoint < 0x80 {
			ret = append(ret, string(codepoint))
			continue
		}
		// Characters in Private Use Area and above are ignored
		if codepoint > 0xffff {
			continue
		}
		section := codepoint >> 8   // Chop off the last two hex digits
		position := codepoint % 256 // Last two hex digits
		table := gd.getCache(section)
		if table != nil && rune(len(table)) > position {
			ret = append(ret, table[position])
		}
	}
	return strings.Trim(strings.Join(ret, ""), " ")
}

// Transliterate Unicode string to a initials.
func (gd *GoDecode) Initials(str string) string {
	var ret []string
	reg := regexp.MustCompile("^\\w|\\s+\\w")
	for _, s := range reg.FindAllString(gd.Decode(str), -1) {
		ret = append(ret, strings.Replace(s, " ", "", -1))
	}
	return strings.Join(ret, "")
}
