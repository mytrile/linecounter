package linecounter

import "io/ioutil"

func GetStats(path string) (*Stats, error) {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	} else {
		language := DetectLanguage(path)
		stats := &Stats{Language: language}
		line_comment := []byte(language.LineComment)
		start_comment := []byte(language.StartComment)
		end_comment := []byte(language.EndComment)
		inComment := 0
		inLComment := false
		blank := true

		lp, sp, ep := 0, 0, 0

		for _, b := range file {
			if inComment == 0 && b == line_comment[lp] {
				lp++
				if lp == len(line_comment) {
					inLComment = true
					lp = 0
				}
			} else {
				lp = 0
			}
			if !inLComment && b == start_comment[sp] {
				sp++
				if sp == len(start_comment) {
					inComment++
					if inComment > 1 && !language.Nesting {
						inComment = 1
					}
					sp = 0
				}
			} else {
				sp = 0
			}
			if !inLComment && inComment > 0 && b == end_comment[ep] {
				ep++
				if ep == len(end_comment) {
					if inComment > 0 {
						inComment--
					}
					ep = 0
				}
			} else {
				ep = 0
			}

			if b != byte(' ') && b != byte('\t') && b != byte('\n') {
				blank = false
			}
			if b == byte('\n') {
				stats.TotalLines++
				if inComment > 0 || inLComment {
					inLComment = false
					stats.CommentLines++
				} else if blank {
					stats.BlankLines++
				} else {
					stats.CodeLines++
				}
				blank = true
				continue
			}
		}
		return stats, nil
	}
}

func DetectLanguage(path string) Language {
	for _, lang := range languages {
		if lang.Match(path) {
			return lang
		}
	}
	return Language{"Generic Text", mExt(".txt"), noComments}
}
