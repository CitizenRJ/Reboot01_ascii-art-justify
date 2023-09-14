package asciiArtJustify

import (
	"fmt"
	"strings"
)

// Print the full outcome in the terminal
func PrintBannersWithColors(Str, colors string, banners []string) {
	// fmt.Println(banners[9])
	colors = strings.ToLower(colors)
	num := 0
	RGB := "rgb"
	HEX := "#"
	HSL := "hsl"
	// ANSI := "\\033["
	ANSICheck := false
	// ResetColor := "\\033[0m"
	var color []string
	//this one for RGP to ansi nums.
	if strings.HasPrefix(colors, RGB) {
		color = RGBtoNum(colors)
		//this one for HEX to ansi nums.
	} else if strings.HasPrefix(colors, HEX) {
		color = HextoRGB(colors)
		//this one for HSL to ansi nums.
	} else if strings.HasPrefix(colors, HSL) {
		color = HSLtoRGB(colors)
		//for ANSI
	} else {
		//for the left (word: red, green, blue, etc...).
		color = WordColors(colors)
	}
	if !(ANSICheck) {
		colors = "\033[38;2;" + color[0] + ";" + color[1] + ";" + color[2] + "m"
	}
	for _, ch := range banners {
		num = num + 1
		if ch == "" {
			if num < len(banners) {
				fmt.Println()
				continue
			} else {
				continue
			}
		}
		if Str == "" {
			for _, j := range ch {
				fmt.Print(colors, j)
			}
		} else {
			h := 0
			count := 0
			match := false
			for _, j := range ch {
				if !match || count >= len(Str) {
					h = h + 1
				}
				check := true
				for q := 0; q < len(Str); q++ {
					if rune(Str[q]) == j {
						word := ch
						if count < len(Str) {
							if Str == word[h-1:h+len(Str)-1] || (match && count < len(Str)) {
								match = true
								count = count + 1
								// if {
								fmt.Print(colors, ch[h-1+q])
								check = false
							}
							break
						}
					}
				}
				if check == true {
					fmt.Print("\033[0m", ch)
				}
			}
			count = 0
		}
		fmt.Println("\033[0m")
	}
}
