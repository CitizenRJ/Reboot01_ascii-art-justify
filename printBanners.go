package asciiArtJustify

import "fmt"

// Print the full outcome in the terminal
func PrintBanners(banners []string) {
	num := 0
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
		fmt.Print(ch)
		fmt.Println()
	}
}
