package main

import (
	"fmt"
	"github.com/beevik/ntp"
	"time"
)

func main() {
	Test2()
}

func Test2() {
	ntpTime, err := ntp.Time("pool.ntp.org") //time.apple.com
	if err != nil {
		fmt.Println(err)
	}

	ntpTimeFormatted := ntpTime.Format(time.UnixDate)
	fmt.Println("----------------------")
	fmt.Printf("Network time: %v\n", ntpTime)
	fmt.Printf("Unix Date Network time: %v\n", ntpTimeFormatted)
	fmt.Println("+++++++++++++++++++++++++++++++")
	timeFormatted := time.Now().Local().Format(time.UnixDate)
	fmt.Printf("System time: %v\n", time.Now())
	fmt.Printf("Unix Date System time: %v\n", timeFormatted)

	// so from here, use ntpTime instead of time.Now().Local()
	// or set your machine time to sync with a network time server
	// for example, on Mac OSX, you can tick to enable auto sync
	// with time.apple.com
	a := "cq31JHWN/HWQ54mC4TIVeF2HNq+1ZdVpVQecpt5Mka7uWO1BnTj7vppKvArXRL+N\nby3n2VWzAn93LSwoctLcEph8z8LO010NCCd1NU8aLhqZv8v9NE0kqssu+QI0a41j\n0b1cN+8yz/h5GfBW1KS+0Bdst2sZDYSFXQty4oTvjECi7klbOIHLt851SvWTSNM9\nybxl+VBMyM+YReaUgB7H6q61moqoGXbKWwJmCeD/6K6M3s+akVwTq2b7/3nWeh0s\nQssHhzHm9t1GUyp2KKKpfEOPMxQQBTTJmYqRSnYVXDtyUq8+Oeao+QCxdLPlvB4f\nKwf5kI5PZk4d3XVunHpJdA=="
	b := "IjVQsR/pg5clNwsNL+464Xm9DTP3Byg/Cbzev9+gSO26QgUVYQe3FN35f27/tc1t\nhWSHEzu2WjZO3pZpXh69ZKUlcp3Eh36ENRb4F/SaQUnGBvGPuwfr7gK804FvR0Wl\nXThd8UlUkKKQsPfkz8nbJjVHQVRLIAfHMrQXzaiMo4FTFSVgQunDYH1325QVaIRx\nlp7tKfxzOKNsA521gJyfExE2jVXULtxZho6tEeHkwflgHhw7OibyuZ9YVUlZTg43\nV8zkLWAtaSZscrtX7kbLrUR4x5/V0ZSt9Ph2LcKz33jXL9NiKUR0C/dLces/5oaw\nqC71QuMnPuykhJMDamDOVQ=="

	a := "LS/OOTiHgWnKeYlyq1LIogH3/E2g7LHTt5ooMtnMlRc2pMi+YcBTFOeeGRq5s1aI\n/ryGHgiqwQJdfqydnNWhsl32Z91JOmGhs8wBpv7jZSRWaPUMhiqGZaBDu6jLqDts\nmOFb3dqJj1U0oygNLAN4ZZUM48o2QITP4fO36Nci+FbiK+67lgFT7izW3E0dzCcE\n2jvKnhJlqQVa/q77TmsQ7BH27e7Dm1Q5Dzshj8TjE+68PU+Oh+LQhS8XBYHh7582\nRrHytlXYnj/IgEi9BNl3vbPnsHUybcZ31m/9q/M8eUfLh3y6bShh737OzflmoZIQ\nowv3tmHbNR8jOXqXTxzQNQ=="
	b := "Gm/A3GOF7ZhlcL2b1c/mO1NeLDvug70nz/eLRqS3agiwCfxE5Wo7v2Ad7zvwsqJ1\npo9CIxNKbTP9BbCzvGeVTuZ7eRtYKgm3ZFWuakGHKIiC6ZLFubOAINFmDkbaPHzW\nnDeBHxuJXRgcuZoLY79Ki1ouksy8TkvAXBmFsea+etSy5aC5O7sz6ov0SIosO7hY\n72vi++Dvh2gr1qNEGd6ko/jwTl4inlTM8yGSJ5HbJwrRrcao4o94BiRD5xR3aIdv\nqoQRLxsdLAn9D7sNz0B+zMJGADM7VGpjzlzW3lbOayHBO6bxrBdb0J+D/LfNzJOp\nrsDx1K/IXxRJRETXD+REyw=="

	a := "EAFe8no5D9m2E+XNaPe8aryXuf3TJbDn4sTkVHzPIz7+J3XPWpJmkaA0Pku96rTX\nm1tk8p3fe7bNOvAb7EEFPC5dyDD9wznmx6azPalLgNQT+cwjV44CcTOCquDI+5RS\nadOzRSDIFQ6CV/VN2/8u+xOHWYwm1M+ezDpqJCpKNujDpNRwsxywveaCZZo8DI9E\nkcDFhr5u915tADo5PWQWFwSJuKMhYfEhDfHKB7SLKTNosrPoWE8mnZLXfxS90Lg2\n8L+7WSN6UtWayo9vbD7zWHPqYTqKhBemHWk0A7aprHU1p95y1HddyFKVorZk2g4e\nVQI++QhypCCSuVpwhjq2sA=="
	b := "lGXliSEVGH4+diJ6yOjqAYkWxxxMmKPEbBdTgbJlwtjZ2W84PGYTpI5iMmpfn6/H\n4y/7TBS7QnRZBnlXeDSN8bVd6apsSKMpAJF3LCuay4O+SjO+SaBiuzqoXN9blPNR\nnPnunrPEGd82kJe2GzPlAx4sBRcweknAdiT2tmZouVeyDkmM+tpAk73FB+VDXLKT\nuvXXnl5xB7q7xpHeteMHpJ4EpZoy6arYSJjY3DxavfZ6C7Pwk2btUF0q0jCletvv\n5Mpj2RPpTZCqz/dcE/LMBWnSFO9O47N+Ofq90gJ/6a1S8Ni1TobCBDnAT560pR/c\ngdS9IbWmO7crCT6YeHPmrg=="
}
