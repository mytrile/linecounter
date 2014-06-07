package linecounter

import "path"

type Language struct {
	Namer
	Matcher
	Commenter
}

var languages = []Language{
	Language{"Thrift", mExt(".thrift"), cComments},

	Language{"C", mExt(".c", ".h"), cComments},
	Language{"C++", mExt(".cc", ".cpp", ".cxx", ".hh", ".hpp", ".hxx"), cComments},
	Language{"Go", mExt(".go"), cComments},
	Language{"Rust", mExt(".rs", ".rc"), cComments},
	Language{"Scala", mExt(".scala"), cComments},
	Language{"Java", mExt(".java"), cComments},

	Language{"YACC", mExt(".y"), cComments},
	Language{"Lex", mExt(".l"), cComments},

	Language{"Lua", mExt(".lua"), luaComments},

	Language{"SQL", mExt(".sql"), sqlComments},

	Language{"Haskell", mExt(".hs", ".lhs"), hsComments},
	Language{"ML", mExt(".ml", ".mli"), mlComments},

	Language{"Perl", mExt(".pl", ".pm"), perlComments},
	Language{"PHP", mExt(".php"), cComments},

	Language{"Shell", mExt(".sh"), shComments},
	Language{"Bash", mExt(".bash"), shComments},
	Language{"R", mExt(".r", ".R"), shComments},
	Language{"Tcl", mExt(".tcl"), shComments},

	Language{"MATLAB", mExt(".m"), matlabComments},

	Language{"Ruby", mExt(".rb"), shComments},
	Language{"Python", mExt(".py"), pyComments},
	Language{"Assembly", mExt(".asm", ".s"), semiComments},
	Language{"Lisp", mExt(".lsp", ".lisp"), semiComments},
	Language{"Scheme", mExt(".scm", ".scheme"), semiComments},

	Language{"Make", mName("makefile", "Makefile", "MAKEFILE"), shComments},
	Language{"CMake", mName("CMakeLists.txt"), shComments},
	Language{"Jam", mName("Jamfile", "Jamrules"), shComments},

	Language{"Markdown", mExt(".md"), noComments},

	Language{"HAML", mExt(".haml"), noComments},
	Language{"SASS", mExt(".sass"), cssComments},
	Language{"SCSS", mExt(".scss"), cssComments},

	Language{"HTML", mExt(".htm", ".html", ".xhtml"), xmlComments},
	Language{"XML", mExt(".xml"), xmlComments},
	Language{"CSS", mExt(".css"), cssComments},
	Language{"JavaScript", mExt(".js"), cComments},

	Language{"Erlang", mExt(".erl"), erlangComments},
}

type Namer string

func (n Namer) Name() string { return string(n) }

type Matcher func(string) bool

func (m Matcher) Match(file_name string) bool { return m(file_name) }

func mExt(exts ...string) Matcher {
	return func(file_name string) bool {
		for _, ext := range exts {
			if ext == path.Ext(file_name) {
				return true
			}
		}
		return false
	}
}

func mName(names ...string) Matcher {
	return func(file_name string) bool {
		for _, name := range names {
			if name == path.Base(file_name) {
				return true
			}
		}
		return false
	}
}
