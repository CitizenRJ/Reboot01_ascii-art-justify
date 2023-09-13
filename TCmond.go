package asciiArtJustify

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

// Check for valid of characters by runes from 32 to 126
func TCmond(hight, width int) int {
	cmd := exec.Command("stty", "size")
	cmd.Stdin = os.Stdin
	out, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
	}
	s := string(out)
	s = strings.TrimSpace(s)
	sArr := strings.Split(s, " ")
	heigth, err := strconv.Atoi(sArr[0])
	if err != nil {
		log.Fatal(err)
	}

	widthh, err1 := strconv.Atoi(sArr[1])
	if err1 != nil {
		log.Fatal(err)
	}
	if hight > heigth {
		fmt.Println("Error: you need to make the hight of the terminal size beggier than this.")
		os.Exit(0)
	} else if width > widthh {
		fmt.Println("Error: you need to make the width of the terminal size beggier than this.")
		os.Exit(0)
	}
	return widthh
}
