package asciiArtJustify

func ForArrSpilet(banners, arr []string) []string {
	var ReturnArr []string
	ForAppened := ""
	num := 0
	for _, ch := range banners {
		num = num + 1

		for i := 0; i < 8; i++ {
			for _, j := range ch {
				n := (j-32)*9 + 1
				if ch == "" {
					if num < len(banners) {
						ForAppened = ForAppened + arr[i]
					}
				} else {
					ForAppened = ForAppened + arr[int(n)+i]
				}
			}
			ReturnArr = append(ReturnArr, ForAppened)
			ForAppened = ""
		}
	}
	return ReturnArr
}
