# linecounter

linecounter is simple library for detecting code/commnet/blank lines in a file.

**NOTE**: This is basically simplified copy of [sloc](https://github.com/bytbox/sloc). All credits go to [bytbox](https://github.com/bytbox)

# Usage

~~~ go
package main

import (
    "github.com/mytrile/linecounter"
)

func main() {
    stats, err := linecounter.GetStats("/path/to/file.go")
    if err != nil {
        // Handle the error
    } else {
        // stats.Language.Name() => Go
        // stats.TotalLines      => 10
        // stats.CodeLines       => 8
        // stats.BlankLines      => 1
        // stats.CommentLines    => 1
    }

}

~~~

## Meta

* Author  : Dimitar Kostov
* Email   : mitko.kostov@gmail.com
* Website : [http://mytrile.github.com](http://mytrile.github.com)
* Twitter : [http://twitter.com/mytrile](http://twitter.com/mytrile)
