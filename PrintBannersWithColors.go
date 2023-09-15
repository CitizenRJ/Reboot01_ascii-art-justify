package asciiArtJustify

import (
	"fmt"
	"math"
	"os"
	"strings"
)

// Print the full outcome in the triminal
func PrintBannersWithColors(Str, colors, Pose string, banners, arr, FinalArr []string) {
	colors = strings.ToLower(colors)
	num := 0
	RGB := "rgb"
	HEX := "#"
	HSL := "hsl"

	ForAppened := ""
	Pose = strings.ToLower(Pose)
	ANSICheck := false
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
	} else {
		//for the left (word: red, green, blue, etc...).
		color = WordColors(colors)
	}

	TWight := TCmond(8, len(FinalArr[0]))
	TWight2 := 0
	if !(ANSICheck) {
		colors = "\033[38;2;" + color[0] + ";" + color[1] + ";" + color[2] + "m"
	}
	for _, ch := range banners {
		BBBBB := 0
		HAHAHAHA := true
		num = num + 1
		if ch == "" {
			if num < len(banners) {
				fmt.Println()
				continue
			} else {
				continue
			}
		}
		for i := 0; i < 8; i++ {
			first := true
			lenlen := 0
			ForAppened = ""
			for _, j := range ch {
				n := (j-32)*9 + 1
				ForAppened = ForAppened + arr[int(n)+i]
				lenlen = len(ForAppened)
			}
			if Str == "" {
				ForAppened = ""
				for v, j := range ch {
					n := (j-32)*9 + 1
					if Pose == "left" || Pose == "" {
						fmt.Print(colors, arr[int(n)+i])
					} else if Pose == "right" {
						if first == true {
							TWight2 = (TWight - lenlen) - 1
							if HAHAHAHA == true {
								BBBBB = TWight2
								HAHAHAHA = false
							}
							if v == 0 {
								for j := 0; j < BBBBB; j++ {
									ForAppened = " " + ForAppened
								}
							}
							first = false
							HAHAHAHA = false
							fmt.Print(ForAppened, colors, arr[int(n)+i])
						} else {
							fmt.Print(colors, arr[int(n)+i])
						}
					} else if Pose == "center" {
						if first == true {
							TWight2 = (TWight/2 - (lenlen-1)/2)
							if HAHAHAHA == true {
								BBBBB = TWight2
								HAHAHAHA = false
							}
							if v == 0 {
								for j := 0; j < BBBBB; j++ {
									ForAppened = " " + ForAppened
								}
							}
							first = false
							HAHAHAHA = false
							fmt.Print(ForAppened, colors, arr[int(n)+i])
						} else {
							fmt.Print(colors, arr[int(n)+i])
						}

					} else if Pose == "justify" {
						if (strings.Count(ch, " ")) != 0 {
							if arr[int(n)+i] == "      " {
								if first == true {
									if (strings.Count(ch, " ")) == 1 {
										TWight2 = (TWight - lenlen)
									} else {
										TWight2 = (int(math.Round(float64((TWight - lenlen) / (strings.Count(ch, " "))))))
									}
									if HAHAHAHA == true {
										BBBBB = TWight2
										HAHAHAHA = false
									}
									ForAppened = ForAppened + strings.Repeat(" ", BBBBB)
									first = false
									HAHAHAHA = false
									fmt.Print(colors, arr[int(n)+i], ForAppened)
								} else {
									fmt.Print(colors, arr[int(n)+i], ForAppened)
								}
							} else {
								fmt.Print(colors, arr[int(n)+i])
							}
						} else {
							if first == true {
								TWight2 = (TWight/2 - (lenlen-1)/2)
								if HAHAHAHA == true {
									BBBBB = TWight2
									HAHAHAHA = false
								}
								for j := 0; j < BBBBB; j++ {
									ForAppened = " " + ForAppened
								}
								first = false
								HAHAHAHA = false
								fmt.Print(ForAppened, colors, arr[int(n)+i])
							} else {
								fmt.Print(colors, arr[int(n)+i])
							}
						}
					} else {
						fmt.Println("Error: Wrong Pose.")
						os.Exit(0)
					}

				}

			} else {
				h := 0
				count := 0
				match := false
				ForAppened = ""
				for v, j := range ch {
					n := (j-32)*9 + 1
					if Pose == "left" || Pose == "" {
					} else if Pose == "right" {
						TWight2 = (TWight - lenlen) - 1
						if v == 0 {
							for j := 0; j < TWight2; j++ {
								ForAppened = " " + ForAppened
							}
						}
					} else if Pose == "center" {
						TWight2 = (TWight/2 - (lenlen-1)/2)
						if v == 0 {
							for j := 0; j < TWight2; j++ {
								ForAppened = " " + ForAppened
							}
						}
					} else if Pose == "justify" {
						if (strings.Count(ch, " ")) != 0 {
							if arr[int(n)+i] == "      " {
								if (strings.Count(ch, " ")) == 1 {
									TWight2 = (TWight - lenlen) - 1
								} else {
									TWight2 = (int(math.Round(float64((TWight - lenlen) / (strings.Count(ch, " ")))))) - 1
								}
								ForAppened = ForAppened + strings.Repeat(" ", TWight2)
							}
						} else {
							TWight2 = (TWight/2 - (lenlen-1)/2)
							if v == 0 {
								for j := 0; j < TWight2; j++ {
									ForAppened = " " + ForAppened
								}
							}
							ForAppened = ForAppened
						}
					} else {
						fmt.Println("Error: Wrong Pose.")
						os.Exit(0)
					}
				}

				for v, j := range ch {
					ForAppened = ""

					if !match || count >= len(Str) {
						h = h + 1
					}

					check := true
					n := (j-32)*9 + 1
					for q := 0; q < len(Str); q++ {
						if rune(Str[q]) == j {
							word := ch
							if count < len(Str) {
								if len(Str) < len(word) {
									if Str == word[h-1:h+len(Str)-1] || (match && count < len(Str)) {
										match = true
										count = count + 1
										if Pose == "left" || Pose == "" {
											fmt.Print(colors, arr[int(n)+i])
										} else if Pose == "right" {
											if first == true {
												TWight2 = (TWight - lenlen) - 1
												if HAHAHAHA == true {
													BBBBB = TWight2
													HAHAHAHA = false
												}
												if v == 0 {
													for j := 0; j < BBBBB; j++ {
														ForAppened = " " + ForAppened
													}
												}
												first = false
												HAHAHAHA = false
												fmt.Print(ForAppened, colors, arr[int(n)+i])
											} else {
												fmt.Print(colors, arr[int(n)+i])
											}
										} else if Pose == "center" {
											if first == true {
												TWight2 = (TWight/2 - (lenlen-1)/2)
												if HAHAHAHA == true {
													BBBBB = TWight2
													HAHAHAHA = false
												}
												if v == 0 {
													for j := 0; j < BBBBB; j++ {
														ForAppened = " " + ForAppened
													}
												}
												first = false
												HAHAHAHA = false
												fmt.Print(ForAppened, colors, arr[int(n)+i])
											} else {
												fmt.Print(colors, arr[int(n)+i])
											}

										} else if Pose == "justify" {
											if (strings.Count(ch, " ")) != 0 {
												if arr[int(n)+i] == "      " {
													if first == true {
														if (strings.Count(ch, " ")) == 1 {
															TWight2 = (TWight - lenlen)
														} else {
															TWight2 = (int(math.Round(float64((TWight - lenlen) / (strings.Count(ch, " "))))))
														}
														if HAHAHAHA == true {
															BBBBB = TWight2
															HAHAHAHA = false
														}
														ForAppened = ForAppened + strings.Repeat(" ", BBBBB)
														first = false
														HAHAHAHA = false
														fmt.Print(colors, arr[int(n)+i], ForAppened)
													} else {
														fmt.Print(colors, arr[int(n)+i], ForAppened)
													}
												} else {
													fmt.Print(colors, arr[int(n)+i])
												}
											} else {
												if first == true {
													TWight2 = (TWight/2 - (lenlen-1)/2)
													if HAHAHAHA == true {
														BBBBB = TWight2
														HAHAHAHA = false
													}
													for j := 0; j < BBBBB; j++ {
														ForAppened = " " + ForAppened
													}
													first = false
													HAHAHAHA = false
													fmt.Print(ForAppened, colors, arr[int(n)+i])
												} else {
													fmt.Print(colors, arr[int(n)+i])
												}
											}
										} else {
											fmt.Println("Error: Wrong Pose.")
											os.Exit(0)
										}
										check = false
									}
								}
								break
							}
						}
					}
					if check == true {

						if Pose == "left" || Pose == "" {
							fmt.Print("\033[0m", arr[int(n)+i])
						} else if Pose == "right" {
							if first == true {
								TWight2 = (TWight - lenlen) - 1
								if HAHAHAHA == true {
									BBBBB = TWight2
									HAHAHAHA = false
								}
								if v == 0 {
									for j := 0; j < BBBBB; j++ {
										ForAppened = " " + ForAppened
									}
								}
								first = false
								HAHAHAHA = false
								fmt.Print(ForAppened, "\033[0m", arr[int(n)+i])
							} else {
								fmt.Print("\033[0m", arr[int(n)+i])
							}
						} else if Pose == "center" {
							if first == true {
								// TWight2 = 0
								TWight2 = (((TWight) / 2) - ((lenlen - 1) / 2))
								if HAHAHAHA == true {
									BBBBB = TWight2
									HAHAHAHA = false
								}
								if v == 0 {
									for j := 0; j < BBBBB; j++ {
										ForAppened = " " + ForAppened
									}
								}
								first = false
								HAHAHAHA = false
								fmt.Print(ForAppened, "\033[0m", arr[int(n)+i])
							} else {
								fmt.Print("\033[0m", arr[int(n)+i])
							}

						} else if Pose == "justify" {
							if (strings.Count(ch, " ")) != 0 {
								if arr[int(n)+i] == "      " {
									if first == true {
										if (strings.Count(ch, " ")) == 1 {
											TWight2 = (TWight - lenlen)
										} else {
											TWight2 = (int(math.Round(float64((TWight - lenlen) / (strings.Count(ch, " "))))))
										}
										if HAHAHAHA == true {
											BBBBB = TWight2
											HAHAHAHA = false
										}
										ForAppened = ForAppened + strings.Repeat(" ", BBBBB)
										first = false
										HAHAHAHA = false
										fmt.Print("\033[0m", arr[int(n)+i], ForAppened)
									} else {
										fmt.Print("\033[0m", arr[int(n)+i], ForAppened)
									}
								} else {
									fmt.Print("\033[0m", arr[int(n)+i])
								}
							} else {
								if first == true {
									TWight2 = (TWight/2 - (lenlen-1)/2)
									if HAHAHAHA == true {
										BBBBB = TWight2
										HAHAHAHA = false
									}
									for j := 0; j < BBBBB; j++ {
										ForAppened = " " + ForAppened
									}
									first = false
									HAHAHAHA = false
									fmt.Print(ForAppened, "\033[0m", arr[int(n)+i])
								} else {
									fmt.Print("\033[0m", arr[int(n)+i])
								}
							}
						} else {
							fmt.Println("Error: Wrong Pose.")
							os.Exit(0)
						}
					}
				}
				count = 0
			}
			fmt.Println("\033[0m")
		}
	}
}
