# Rob Pike's Regular Expression Matching, From C to Go

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

I had read _The Practice of Programming_ book,
and seen the article back in the good old days.
Discovering Kernighan's article compelled me to try the code,
because I firmly believe you have to do something to understand that thing.

In the case of source code and computer stuff,
I will re-type what appears on-line, so that I can understand it on a line-by-line basis.
I often do minor modifications to ensure I'm really understanding and not fooling myself,
maybe using SQLite instead of MySQL, or adding an operation.
Time time, I decided to use Go instead of C.
I also ended up cut-n-pasting the C code into it's own [file](matching.go),
so I could see if some constructs (`.*`, `a*b*`) worked in my Go port.

## Building and running

```
$ git clone https://github.com/bediger4000/simpleregex.git
$ cd simpleregex
$ make test
...
    --- PASS: Test_match/Kleene_closure_of_dot_metacharacter_ending_regex_3 (0.00s)
PASS
ok      simpleregex     (cached)
$ make all
```

## Porting to Go

Sometimes C code can easily transliterate to Go and vice versa,
but this is not one of them.
Pike's code uses lots of C idioms, like `== '\0'` for end-of-string,
and `*text++`, which even the best of us have to think hard about.
It also uses 2 do-while loops, which are harder in Go.

## Go Unit Testing

I wrote Go unit testing that works with `go test`.
This uncovered a bug in my first cut at a Go version, `.*` did not work,
and two Kleene closures in a row failed (`ab*c*d` would not match `abbbccccd`).
Part of the problem involved the transliteration of a do-while loop
condition involving `*text++`.
I had a traditional for-loop with an index variable, and I "advanced"
the text to be matched via `text = text[1:]`.
The interaction was incorrect.

Unit testing of each function does not make sense, they're mutually recursive.
