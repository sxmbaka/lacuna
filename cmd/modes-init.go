package cmd

import "github.com/sxmbaka/lacuna/mode"

var (
	ml = mode.MoreLike{
		Active: false,
		Arg:    "test",
	}
	sl = mode.SoundsLike{
		Active: true,
		Arg:    "silence",
	}
	sp = mode.SpelledLike{
		Active: true,
		Arg:    "cry",
	}
)
