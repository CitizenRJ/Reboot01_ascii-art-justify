package asciiArtJustify

import (
	// "fmt"
	"fmt"
	"math"
	"os"
	"strings"
)

func PreparedArr(Pose string, banners, arr []string) []string {
	var ReturnArr []string
	ForAppened := ""
	Pose = strings.ToLower(Pose)
	num := 0
	TWight2 := 0
	for _, ch := range banners {
		num = num + 1
		if ch == "" {
			if num < len(banners) {
				ReturnArr = append(ReturnArr, "")
				continue
			} else {
				continue
			}
		}

		for i := 0; i < 8; i++ {
			lenlen := 0
			for _, j := range ch {
				n := (j-32)*9 + 1
				ForAppened = ForAppened + arr[int(n)+i]
				lenlen = len(ForAppened)
			}
			ForAppened = ""
			for v, j := range ch {
				n := (j-32)*9 + 1
				ForAppened = ForAppened + arr[int(n)+i]
				if Pose == "left" || Pose == "" {
				} else if Pose == "right" {
					TWight := TCmond(8, lenlen)
					TWight2 = (TWight - lenlen) - 1
					if v == 0 {
						for j := 0; j < TWight2; j++ {
							ForAppened = " " + ForAppened
						}
					}
				} else if Pose == "center" {
					TWight := TCmond(8, lenlen)
					TWight2 = (TWight/2 - (lenlen-1)/2)
					if v == 0 {
						for j := 0; j < TWight2; j++ {
							ForAppened = " " + ForAppened
						}
					}
				} else if Pose == "justify" {
					if (strings.Count(ch, " ")) != 0 {
						if arr[int(n)+i] == "      " {
							TWight := TCmond(8, lenlen)
							if (strings.Count(ch, " ")) == 1 {
								TWight2 = (TWight - lenlen) - 1
							} else {
								TWight2 = (int(math.Round(float64((TWight - lenlen) / (strings.Count(ch, " ")))))) - 1
							}
							if TWight2 <= 0 {

							} else {
								ForAppened = ForAppened + strings.Repeat(" ", TWight2)
							}
						}
					} else {
						
					}
				} else {
					fmt.Println("Error: Wrong Pose.")
					os.Exit(0)
				}
			}
			ReturnArr = append(ReturnArr, ForAppened)
			ForAppened = ""

		}
	}
	return ReturnArr
}
