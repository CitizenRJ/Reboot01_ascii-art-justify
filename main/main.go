package main

import (
	"asciiArtJustify"
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

const (
	fileLen = 855
)

func main() {
	var FlagEnd []string
	args := os.Args[1:]
	ArgsLen := len(args)
	font := "standard"
	StrFlagArr := []string{"--output=", "--color=", "--align="}
	outputFile := ""
	ColorColor := ""
	textPos := ""
	text := ""
	Str := ""
	FlagEnd = append(os.Args[1:],"")
	FlagEnd = append(FlagEnd[:len(FlagEnd)-1])
	if ArgsLen < 1 {
		fmt.Println(len(os.Args), "is Not a valid amount of arguments.\n")
		fmt.Println("Usage: go run . [STRING] [BANNER]\nEX: go run . something standard\n\n# Ascii Art output #\nUsage: go run . [OPTION] [STRING] [BANNER]\nEX: go run . --output=<fileName.txt> something standard\n\n# Ascii Art color #\nUsage: go run . [OPTION] [STRING]\nEX: go run . --color=<color> <letters to be colored> \"something\" . \n--align")
		return
		// ///////////////////////////////////////////////////////////////////////////////////////////////////////////////////
	} else if ArgsLen > 0 && ArgsLen < 3 {
		num := 0
		flagCount := 0
		for i := 0; i < ArgsLen; i++ {
			num = num + 1
			if strings.Contains(args[i], StrFlagArr[0]) {
				outputFile = outputFileCheck(args[i])
				if ColorColor == "" {
					if i+1 < ArgsLen {
						FlagEnd = append(FlagEnd[:num-1], FlagEnd[num:]...)
						num = num - 1
					} else {
						FlagEnd = FlagEnd[:num-1]
					}

					flagCount = flagCount + 1
				} else {
					outputFile = ""
					continue
				}
			} else if strings.Contains(args[i], StrFlagArr[1]) {
				ColorColor = ColorColorCheck(args[i])
				if outputFile == "" {
					if i+1 < ArgsLen {
						FlagEnd = append(FlagEnd[:num-1], FlagEnd[num:]...)
						num = num - 1
					} else {
						FlagEnd = FlagEnd[:num-1]
					}
					flagCount = flagCount + 1
				} else {
					continue
				}
			} else if strings.Contains(args[i], StrFlagArr[2]) {
				textPos = CheckJustify(args[i])
				if i+1 < ArgsLen {
					FlagEnd = append(FlagEnd[:num-1], FlagEnd[num:]...)
					num = num - 1
				} else {
					FlagEnd = FlagEnd[:num-1]
				}
				flagCount = flagCount + 1
			}
		}
		if flagCount == 0 {
			fmt.Println("Error: Invalid arguments.")
			fmt.Println("Usage: go run . [STRING] [BANNER]\nEX: go run . something standard\n\n# Ascii Art output #\nUsage: go run . [OPTION] [STRING] [BANNER]\nEX: go run . --output=<fileName.txt> something standard\n\n# Ascii Art color #\nUsage: go run . [OPTION] [STRING]\nEX: go run . --color=<color> <letters to be colored> \"something\"")
			return
		}
		// ///////////////////////////////////////////////////////////////////////////////////////////////////////////////////
	} else if ArgsLen > 2 && ArgsLen < 7 {
		num := 0
		flagCount := 0
		
		for i := 0; i < ArgsLen; i++ {
			// fmt.Println(args)
			// fmt.Println(i, args[i])
			num = num + 1
			if strings.Contains(args[i], StrFlagArr[0]) {
				if ColorColor == "" {
					outputFile = outputFileCheck(args[i])
					if i+1 < ArgsLen {
						FlagEnd = append(FlagEnd[:num-1], FlagEnd[num:]...)
					} else {
						FlagEnd = FlagEnd[:num-1]
					}
					num = num - 1
				} else {
					continue
				}
				flagCount = (flagCount + 1)
				// fmt.Println(args)
			} else if strings.Contains(args[i], StrFlagArr[1]) {
				ColorColor = ColorColorCheck(args[i])
				if outputFile == "" {
					FlagEnd = append(FlagEnd[:num-1], FlagEnd[num:]...)
					num = num - 1
					if i+1 < ArgsLen {
						Str = FlagEnd[num]
						FlagEnd = append(FlagEnd[:num], FlagEnd[num+1:]...)
						num = num - 1
					} else {
						FlagEnd = FlagEnd[:num]
					}
				} else if outputFile != "" {
					if i+1 < ArgsLen {
						FlagEnd = append(FlagEnd[:num], FlagEnd[num+1:]...)
					} else {
						FlagEnd = FlagEnd[:num]
					}
					num = num - 1
				} else if i+1 < ArgsLen {
					if strings.Contains(args[i+1], StrFlagArr[2]) {
						continue
					} else {
						Str = args[i+1]
					if i+2 < ArgsLen {
						FlagEnd = append(FlagEnd[:num], FlagEnd[num+2:]...)
						num = num - 2
					} else {
						FlagEnd = FlagEnd[:num+1]
					}
					}
					
				}
				flagCount = flagCount + 1
				// ///////////////////////////////////////////////////////////////////////////////////////////////////////////////////
			} else if strings.Contains(args[i], StrFlagArr[2]) {
				// fmt.Println(5)
				textPos = CheckJustify(args[i])
				if Str == "" {
					if i+1 < ArgsLen {
					FlagEnd = append(FlagEnd[:num-1], FlagEnd[num:]...)
					num = num - 1
				} else {
					FlagEnd = FlagEnd[:num-1]
				}
				} else {
					textPos = CheckJustify(args[i])
					Str = ""
				}
				
				flagCount = flagCount + 1
			}
		}
		if flagCount == 0 {
			fmt.Println("Error: Invalid arguments.")
			fmt.Println("Usage: go run . [STRING] [BANNER]\nEX: go run . something standard\n\n# Ascii Art output #\nUsage: go run . [OPTION] [STRING] [BANNER]\nEX: go run . --output=<fileName.txt> something standard\n\n# Ascii Art color #\nUsage: go run . [OPTION] [STRING]\nEX: go run . --color=<color> <letters to be colored> \"something\"")
			return
		}
		// ///////////////////////////////////////////////////////////////////////////////////////////////////////////////////
	} else if ArgsLen > 6 {
		fmt.Println("Error: Invalid arguments.")
		fmt.Println("Usage: go run . [STRING] [BANNER]\nEX: go run . something standard\n\n# Ascii Art output #\nUsage: go run . [OPTION] [STRING] [BANNER]\nEX: go run . --output=<fileName.txt> something standard\n\n# Ascii Art color #\nUsage: go run . [OPTION] [STRING]\nEX: go run . --color=<color> <letters to be colored> \"something\" . \n--align")
		return
	}
	ArgsLen = len(FlagEnd)
	// fmt.Println(FlagEnd)
	if ArgsLen == 0 && Str != "" {
		FlagEnd = append(FlagEnd, Str)
		FlagEnd = append(FlagEnd, "standard")
		ArgsLen = len(FlagEnd)
		Str = ""
	} else if ArgsLen == 0 && Str == "" {
		FlagEnd = append(FlagEnd, os.Args[1])
		FlagEnd = append(FlagEnd, "standard")
		ArgsLen = len(FlagEnd)
	} else if ArgsLen == 1 {
		if Str != "" && Fonts(FlagEnd[ArgsLen-1]) {
			lol := FlagEnd[ArgsLen-1]
			FlagEnd[ArgsLen-1] = Str
			FlagEnd = append(FlagEnd, lol)
			Str = ""
			font = lol
		} else {
			FlagEnd = append(FlagEnd, "standard")
		}

		ArgsLen = len(FlagEnd)
	} else if ArgsLen == 2 {
		font = FlagEnd[ArgsLen-1]
		if Fonts(FlagEnd[ArgsLen-1]) {
			font = FlagEnd[ArgsLen-1]

		} else {
			fmt.Println("\""+font+"\"", "is Not a valid font.")
			os.Exit(0)
		}
	} else {
		fmt.Println("Error: Invalid arguments.", FlagEnd)
		fmt.Println("Usage: go run . [STRING] [BANNER]\nEX: go run . something standard\n\n# Ascii Art output #\nUsage: go run . [OPTION] [STRING] [BANNER]\nEX: go run . --output=<fileName.txt> something standard\n\n# Ascii Art color #\nUsage: go run . [OPTION] [STRING]\nEX: go run . --color=<color> <letters to be colored> \"something\"")
		return
	}

	for i := 0; i < len(args); i++ {
		if !(asciiArtJustify.IsValid(args[i])) {
			fmt.Println(args[i], "isn't a valid character/argument.")
			return
		}
	}

	text = FlagEnd[ArgsLen-2]
	if len(Str) > len(text) {
		fmt.Println("the \"", Str, "\" should be less or equal than \"", text, "\".")
		return
	} else if !(strings.Contains(text, Str)) {
		fmt.Println("the string", "'"+text+"'", "should contains", "'"+Str+"'", "to Print.")
		os.Exit(1)
	}

	// Read the content of the file
	text = strings.ReplaceAll(text, "\\t", "   ")
	argsArr := strings.Split(strings.ReplaceAll(text, "\\n", "\n"), "\n")
	arr := []string{}
	readFile, err := os.Open("fonts/" + font + ".txt")
	defer readFile.Close()

	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		arr = append(arr, fileScanner.Text())
	}

	if len(arr) != fileLen {
		fmt.Println("File is corrupted.")
		return
	}
	larg := len(argsArr)
	if larg >= 2 {
		if argsArr[larg-1] == "" && argsArr[larg-2] != "" {
			argsArr = argsArr[:larg-1]
		}
	}
	if ColorColor == "" {
		FinalArr := asciiArtJustify.PreparedArr(textPos, argsArr, arr)
	}
	

	if outputFile != "" {
		asciiArtJustify.PrintBannersInFile(outputFile, FinalArr)
	} else if ColorColor != "" {
		asciiArtJustify.PrintBannersWithColors(Str, ColorColor, textPos, argsArr, arr)
	} else {
		asciiArtJustify.PrintBanners(FinalArr)
	}
}

func Fonts(argFont string) bool {
	// font := ""
	switch argFont {
	case "shadow":
		argFont = "shadow"
	case "thinkertoy":
		argFont = "thinkertoy"
	case "standard":
		argFont = "standard"
	default:
		return false
	}
	return true
}

// outputFile Error management.
func outputFileCheck(output string) string {
	outputFile := strings.TrimPrefix(output, "--output=")
	if !strings.HasSuffix(outputFile, ".txt") {
		fmt.Println("Output file name should end with .txt")
		os.Exit(1)
	}
	return outputFile
}

func ColorColorCheck(color string) string {
	textColor := strings.TrimPrefix(color, "--color=")
	if textColor == "" {
		fmt.Println("Error: where's the color?.")
		os.Exit(0)
	}
	return textColor
}

func CheckJustify(justify string) string {
	textPos := strings.TrimPrefix(justify, "--align=")
	if textPos == "" {
		fmt.Println("Error: where's the Pose?.")
		os.Exit(0)
	}
	return textPos
}
