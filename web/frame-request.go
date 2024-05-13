package web

import (
	"fmt"

	"github.com/sxmbaka/lacuna/mode"
)

func frameRequest(ml *mode.MoreLike, sl *mode.SoundsLike, sp *mode.SpelledLike, max int) string {
	url := baseURL + "?"

	if ml.Active {
		url += "&ml=" + handleSpaces(ml.Arg)
	}
	if sl.Active {
		url += "&sl=" + handleSpaces(sl.Arg)
	}
	if sp.Active {
		url += "&sp=" + handleSpaces(sp.Arg)
	}
	url += "&max=" + fmt.Sprint(max)
	fmt.Println(url)
	return url
}
