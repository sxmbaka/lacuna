package web

import (
	"fmt"
	"os"

	"github.com/sxmbaka/lacuna/mode"
)

func frameRequest(ml *mode.MoreLike, sl *mode.SoundsLike, sp *mode.SpelledLike, max int, show bool) string {
	url := baseURL + "?"

	if ml.Active {
		if !sl.Active && !sp.Active && hasWildcard(ml.Arg) {
			// if only ml is active and it has a wildcard
			fmt.Println("Error: Wildcard detected in more-like query\n\nWildcards are only supported when used with sounds-like or spelled-like queries.\nIf you use wildcards in more-like while also using other queries, the means-like query will be ignored.")
			os.Exit(1)
		}
		if !hasWildcard(ml.Arg) {
			url += "&ml=" + handleSpaces(ml.Arg)
		}
	}
	if sl.Active {
		url += "&sl=" + handleSpaces(sl.Arg)
	}
	if sp.Active {
		url += "&sp=" + handleSpaces(sp.Arg)
	}
	url += "&max=" + fmt.Sprint(max)
	if show {
		fmt.Println("Query URL:", url)
	}
	return url
}
