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
	// counter := 0
	// TWight := TCmond(8, len(ReturnArr[0]))
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
			// cooo := true
			for v, j := range ch {
				n := (j-32)*9 + 1
				ForAppened = ForAppened + arr[int(n)+i]
				if Pose == "left" || Pose == "" {
					// ForAppened = ""
				} else if Pose == "right" {
					TWight := TCmond(8, lenlen)
					TWight = (TWight - lenlen) - 1
					if v == 0 {
						for j := 0; j < TWight; j++ {
							ForAppened = " " + ForAppened
						}
					}
				} else if Pose == "center" {
					TWight := TCmond(8, lenlen)
					TWight = (TWight/2 - (lenlen-1)/2)
					// fmt.Println(TWight)
					if v == 0 {
						for j := 0; j < TWight; j++ {
							ForAppened = " " + ForAppened
						}
					}
				} else if Pose == "justify" {

					// fmt.Print(cooo)
					if (strings.Count(ch, " ")) != 0 {
						if arr[int(n)+i] == "      " {
							TWight := TCmond(8, lenlen)
							// fmt.Println(TWight)
							if (strings.Count(ch, " ")) == 1 {
								TWight = (TWight - lenlen) - 1
							} else {
								TWight = (int(math.Round(float64((TWight - lenlen) / (strings.Count(ch, " ")))))) - 1
							}

							// fmt.Println(TWight,(strings.Count(ch, " ")))
							ForAppened = ForAppened + strings.Repeat(" ", TWight)
							// cooo = false

						}
					} else {
						TWight := TCmond(8, lenlen)
						TWight = (TWight/2 - (lenlen-1)/2)
						// fmt.Println(TWight)
						if v == 0 {
							for j := 0; j < TWight; j++ {
								ForAppened = " " + ForAppened
							}
						}
						ForAppened = ForAppened
					}
					// ForAppened = strings.ReplaceAll(ForAppened,"      ","      "+strings.Repeat(" ",TWight))
				} else {
					fmt.Println("Error: Wrong Pose.")
					os.Exit(0)
				}
			}
			ReturnArr = append(ReturnArr, ForAppened)
			ForAppened = ""

		}
	}
	// fmt.Println(counter, len(banners[0]))
	return ReturnArr
}
