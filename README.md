# Rob Pike's Regular Expression Matching

I came across an [article](https://www.cs.princeton.edu/courses/archive/spr09/cos333/beautiful.html)
by [Brian Kernighan](https://www.cs.princeton.edu/~bwk/),
which is about a piece of C code written
by [Rob Pike](http://herpolhode.com/rob/),
apparently for a book they wrote together,
[The Practice of Programming](ahttps://www.cs.princeton.edu/~bwk/tpop.webpage/).

I believe Kernighan and Pike also used the code in an
[article](https://jacobfilipp.com/DrDobbs/articles/DDJ/1999/9904/9904a/9904a.htm)
in the legendary [Dr Dobbs Journal](https://en.wikipedia.org/wiki/Dr._Dobb%27s_Journal).

The code implements a subset of [regular expression](https://en.wikipedia.org/wiki/Regular_expression)
text matching:

    c    matches any literal character c
    .    matches any single character
    ^    matches the beginning of the input string
    $    matches the end of the input string
    *    matches zero or more occurrences of the previous character

All this in only about 30 lines of C language code.
It's really not missing much in terms of plain old text matching,
only needing character ranges (like `[a-z]`) and alternation (like `abc|defg`).

I had read the book, and seen the article back in the good old days.
Discovering Kernighan's article compelled me to try the code,
because I firmly believe you have to do something to understand that thing.

In the case of source code and computer stuff,
I will re-type what appears on-line, so that I can understand it on a line-by-line basis.
I often do minor modifications to ensure I'm really understanding and not fooling myself,
maybe using SQLite instead of MySQL, or adding an operation.
In this case I decided to use Go instead of C.
I also ended up cut-n-pasting the C code into it's own [file](matching.go),
so I could see if some constructs (`.*`, `a*b*`) worked in my Go port.

## Porting to Go

- Lots of C idioms
- Go testing
