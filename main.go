package main

import (
	"flag"
	"fmt"
	"os"
	"slices"
)

const (
	Help = iota
	UpperCamel
	LowerCamel
	SturdyCaps
)

var flags_short = map[int]string{
	Help:       "h",
	UpperCamel: "u",
	LowerCamel: "l",
	SturdyCaps: "s",
}

var flags_message = map[int]string{
	Help:       "Help message",
	UpperCamel: "Upper Camel",
	LowerCamel: "Lower Camel",
	SturdyCaps: "SturdyCaps",
}

func tellOriginalSchele(value string) int {
	rune_string := []rune(value)
	caps := 0
	underscores := 0
	for _, char := range rune_string {
		if char >= 'A' && char <= 'Z' {
			caps++
			continue
		} else if char == '_' {
			underscores++
			continue
		}
	}

	// fmt.Printf("caps %v, underscores: %v\n", caps, underscores)
	if caps > 0 && underscores > 0 {
		return UpperCamel
	} else if caps == 0 && underscores >= 0 {
		return LowerCamel
	} else {
		return SturdyCaps
	}
}

func lowerToUpperCamel(lower_string string) string {
	rune_string := []rune(lower_string)
	convert := true
	for idx := range rune_string {
		if convert {
			if rune_string[idx] >= 'a' && rune_string[idx] <= 'z' {
				rune_string[idx] += 'A' - 'a'
			}
			convert = false
		}
		if rune_string[idx] == '_' {
			convert = true
		}
	}

	// fmt.Printf("lowerToUpperCamel rune_string: %v\n", rune_string)
	return string(rune_string)
}

func lowerCamelToStrudy(lower_string string) string {
	rune_string := []rune(lower_string)
	convert := true

	for idx := 0; idx < len(rune_string); idx++ {
		for rune_string[idx] == '_' && idx+1 < len(rune_string) {
			rune_string = append(rune_string[:idx], rune_string[idx+1:]...)
			convert = true
		}

		if convert {
			if rune_string[idx] >= 'a' && rune_string[idx] <= 'z' {
				rune_string[idx] += 'A' - 'a'
			}
			convert = false
		}
	}

	// fmt.Printf("sturdyToLowerCamel rune_string: %v\n", string(rune_string))
	return string(rune_string)
}

func upperToLowerCamel(upper_string string) string {
	rune_string := []rune(upper_string)
	convert := true
	for idx := range rune_string {
		if convert {
			if rune_string[idx] >= 'A' && rune_string[idx] <= 'Z' {
				rune_string[idx] += 'a' - 'A'
			}
			convert = false
		}
		if rune_string[idx] == '_' {
			convert = true
		}
	}

	// fmt.Printf("upperToLowerCamel rune_string: %v\n", string(rune_string))
	return string(rune_string)
}

func sturdyToLowerCamel(sturdy_string string) string {
	rune_string := []rune(sturdy_string)
	inserted := false

	for idx := 0; idx < len(rune_string); idx++ {
		if rune_string[idx] >= 'A' && rune_string[idx] <= 'Z' {
			rune_string[idx] += 'a' - 'A'
		}
		if inserted == false && idx+1 < len(rune_string) && rune_string[idx+1] >= 'A' && rune_string[idx+1] <= 'Z' {
			rune_string = slices.Insert(rune_string, idx+1, '_')
			inserted = true
		} else {
			inserted = false
		}
	}

	// fmt.Printf("sturdyToLowerCamel rune_string: %v\n", string(rune_string))
	return string(rune_string)
}

func main() {
	var value string
	var is_help bool
	var mode int
	var origin_scheme int

	flag.BoolVar(&is_help, flags_short[Help], false, flags_message[Help])
	for idx, flag_short := range flags_short {
		if idx == Help {
			continue
		}
		flag.StringVar(&value, flag_short, "", flags_message[idx])
	}
	flag.Parse()

	if is_help == true || (is_help == false && len(value) == 0) {
		flag.PrintDefaults()
		os.Exit(0)
	}

	origin_scheme = tellOriginalSchele(value)
	flag.Visit(func(visit_flag *flag.Flag) {
		switch visit_flag.Name {
		case flags_short[UpperCamel]:
			{
				mode = UpperCamel
				break
			}
		case flags_short[LowerCamel]:
			{
				mode = LowerCamel
				break
			}
		case flags_short[SturdyCaps]:
			{
				mode = SturdyCaps
				break
			}
		}
	})

	if mode == origin_scheme {
		fmt.Printf("The string is already %s\n", flags_message[mode])
	} else {
		// var new_string string
		switch origin_scheme {
		case UpperCamel:
			{
				value = upperToLowerCamel(value)
				break
			}
		case SturdyCaps:
			{
				value = sturdyToLowerCamel(value)
				break
			}
		}

		switch mode {
		case UpperCamel:
			{
				value = lowerToUpperCamel(value)
				break
			}
		case SturdyCaps:
			{
				value = lowerCamelToStrudy(value)
				break
			}
		}
		fmt.Printf("new string: %v\n", value)
	}
}
