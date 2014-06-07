package linecounter

type Commenter struct {
	LineComment  string
	StartComment string
	EndComment   string
	Nesting      bool
}

var (
	noComments     = Commenter{"\000", "\000", "\000", false}
	xmlComments    = Commenter{"\000", `<!--`, `-->`, false}
	cComments      = Commenter{`//`, `/*`, `*/`, false}
	cssComments    = Commenter{"\000", `/*`, `*/`, false}
	shComments     = Commenter{`#`, "\000", "\000", false}
	semiComments   = Commenter{`;`, "\000", "\000", false}
	hsComments     = Commenter{`--`, `{-`, `-}`, true}
	mlComments     = Commenter{`\000`, `(*`, `*)`, false}
	sqlComments    = Commenter{`--`, `/*`, `*/`, false}
	luaComments    = Commenter{`--`, `--[[`, `]]`, false}
	pyComments     = Commenter{`#`, `"""`, `"""`, false}
	matlabComments = Commenter{`%`, `%{`, `%}`, false}
	erlangComments = Commenter{`%`, "\000", "\000", false}
	perlComments   = Commenter{`#`, "\000", "\000", false}
)
