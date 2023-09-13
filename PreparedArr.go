package asciiArtJustify

import (
	"fmt"
	"os"
	"strings"
)

// Check for valid of characters by runes from 32 to 126
func PreparedArr(Pose string, banners, arr []string) []string {
	var ReturnArr []string
	var argsArr []string
	bannersLen := len(banners)
	fmt.Println(Pose)
	// num := 0
	ReturnArr = ForArry(banners, arr)
	ArrLen := len(ReturnArr[0])
	TWight := 0
	// TWight = TCmond(8, len(ReturnArr[0]))
	// fmt.Println(TWight)
	for _, o := range ReturnArr {
		if o == "" {
			continue
		} else {
			Pose = strings.ToLower(Pose)
		for i := 0; i < 8; i++ {
			if Pose == "left" {
				break
			} else if Pose == "right" {
				TWight = TCmond(8, len(ReturnArr[0]))
				TWight = TWight - ArrLen
				for j := 0; j < TWight; j++ {
					ReturnArr[i] = " " + ReturnArr[i]
				}
			} else if Pose == "center" {
				TWight = TCmond(8, len(ReturnArr[0]))
				TWight = TWight/2 - ArrLen/2
				for j := 0; j < TWight; j++ {
					ReturnArr[i] = " " + ReturnArr[i]
				}
			} else if Pose == "justify" {
				for l := 0; l < bannersLen; l++ {
					argsArr = strings.Split(banners[l], " ")
					if bannersLen == 1 {
						TWight = TCmond(8, len(ReturnArr[0]))
						TWight = TWight/2 - ArrLen/2
						for j := 0; j < TWight; j++ {
							ReturnArr[i] = " " + ReturnArr[i]
						}
					}
				}

			} else {
				fmt.Println("Error: Wrong Type of Pose.")
				os.Exit(0)
			}
		}
		}
		
	}
	return ReturnArr
}

// else if bannersLen == 2 {
// 	TWight = TCmond(8, len(ReturnArr[0]))
// 	TWight = TWight - ArrLen
// 	MyArr := ForArrSpilet(banners[], arr)
// 	fmt.Println(1)
// 	for j := 0; j < TWight; j++ {
// 		MyArr[2][0] = " " + MyArr[2][0]
// 	}
// 	ReturnArr[i] = ""
// 	for f := 1; f < bannersLen-1; f++ {
// 		ReturnArr[i] = ReturnArr[i] + MyArr[i][f]
// 	}
// } else {
// 	TWight = TCmond(8, len(ReturnArr[0]))
// 	TWight = (TWight-ArrLen)/bannersLen - 2
// 	MyArr := ForArrSpilet(banners, arr)
// 	for j := 0; j < TWight; j++ {
// 		for f := 1; f < bannersLen-1; f++ {
// 			MyArr[f][0] = " " + MyArr[f][0]
// 		}

// 	}
// 	ReturnArr[i] = ""
// 	for f := 1; f < bannersLen-1; f++ {
// 		ReturnArr[i] = ReturnArr[i] + MyArr[i][f]
// 	}
// }
