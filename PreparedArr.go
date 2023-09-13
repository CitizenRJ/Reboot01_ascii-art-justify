package asciiArtJustify

import (
	// "fmt"
	"fmt"
	"os"
	"strings"
)

func PreparedArr(Pose string, banners, arr []string) []string {
	var ReturnArr []string
	ForAppened := ""
	Pose = strings.ToLower(Pose)
	num := 0

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
					fmt.Println(TWight)
					if v == 0 {
						for j := 0; j < TWight; j++ {
							ForAppened = " " + ForAppened
						}
					}
				} else if Pose == "justify" {

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
