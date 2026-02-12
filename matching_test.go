package main

import "testing"

func Test_match(t *testing.T) {
	type args struct {
		regexp []rune
		text   []rune
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "match 3-character string",
			args: args{
				regexp: []rune("abc"),
				text:   []rune("abc"),
			},
			want: true,
		},
		{
			name: "mis-match 3-character string at end",
			args: args{
				regexp: []rune("abc"),
				text:   []rune("abd"),
			},
			want: false,
		},
		{
			name: "match 3-character sub-string",
			args: args{
				regexp: []rune("abc"),
				text:   []rune("ABCDEFabcGHIJK"),
			},
			want: true,
		},
		{
			name: "match 3-character prefix",
			args: args{
				regexp: []rune("^abc"),
				text:   []rune("abcABCDEFGHIJK"),
			},
			want: true,
		},
		{
			name: "mis-match 1-character prefix",
			args: args{
				regexp: []rune("^a"),
				text:   []rune("AbcABCDEFGHIJK"),
			},
			want: false,
		},
		{
			name: "match 3-character suffix",
			args: args{
				regexp: []rune("abc$"),
				text:   []rune("ABCDEFGHIJKabc"),
			},
			want: true,
		},
		{
			name: "mis-match 1-character suffx",
			args: args{
				regexp: []rune("c$"),
				text:   []rune("AbcABCDEFGHIJKabC"),
			},
			want: false,
		},
		{
			name: "match 3-character word",
			args: args{
				regexp: []rune("^abc$"),
				text:   []rune("abc"),
			},
			want: true,
		},
		{
			name: "Dot metacharacter in regex",
			args: args{
				regexp: []rune("a.c"),
				text:   []rune("ABCDEFabcGHIJK"),
			},
			want: true,
		},
		{
			name: "Dot metacharacter starting regex 1",
			args: args{
				regexp: []rune(".bc"),
				text:   []rune("Abc"),
			},
			want: true,
		},
		{
			name: "Dot metacharacter starting regex 2",
			args: args{
				regexp: []rune(".bc"),
				text:   []rune("Pbc"),
			},
			want: true,
		},
		{
			name: "Dot metacharacter ending regex 1",
			args: args{
				regexp: []rune("abc."),
				text:   []rune("abcd"),
			},
			want: true,
		},
		{
			name: "Kleene closure starting regex 1",
			args: args{
				regexp: []rune("a*bc"),
				text:   []rune("abc"),
			},
			want: true,
		},
		{
			name: "Kleene closure starting regex 2",
			args: args{
				regexp: []rune("a*bc"),
				text:   []rune("aaaaaaaaaaaabc"),
			},
			want: true,
		},
		{
			name: "Kleene closure starting regex 3",
			args: args{
				regexp: []rune("a*bc"),
				text:   []rune("bc"),
			},
			want: true,
		},
		{
			name: "Kleene closure starting regex 4",
			args: args{
				regexp: []rune("^a*bc"),
				text:   []rune("bcABC"),
			},
			want: true,
		},
		{
			name: "Kleene closure middle of regex 1",
			args: args{
				regexp: []rune("ab*c"),
				text:   []rune("abc"),
			},
			want: true,
		},
		{
			name: "Kleene closure middle of regex 2",
			args: args{
				regexp: []rune("ab*c"),
				text:   []rune("ac"),
			},
			want: true,
		},
		{
			name: "Kleene closure middle of regex 3",
			args: args{
				regexp: []rune("ab*c"),
				text:   []rune("abbbbbbc"),
			},
			want: true,
		},
		{
			name: "Two adjacent Kleene closures 1",
			args: args{
				regexp: []rune("ab*c*d"),
				text:   []rune("ad"),
			},
			want: true,
		},
		{
			name: "Two adjacent Kleene closures 2",
			args: args{
				regexp: []rune("ab*c*d"),
				text:   []rune("abd"),
			},
			want: true,
		},
		{
			name: "Two adjacent Kleene closures 3",
			args: args{
				regexp: []rune("ab*c*d"),
				text:   []rune("acd"),
			},
			want: true,
		},
		{
			name: "Two adjacent Kleene closures 4",
			args: args{
				regexp: []rune("ab*c*d"),
				text:   []rune("abcd"),
			},
			want: true,
		},
		{
			name: "Two adjacent Kleene closures 4",
			args: args{
				regexp: []rune("ab*c*d"),
				text:   []rune("abbbcd"),
			},
			want: true,
		},
		{
			name: "Two adjacent Kleene closures 4",
			args: args{
				regexp: []rune("ab*c*d"),
				text:   []rune("abccccd"),
			},
			want: true,
		},
		{
			name: "Two adjacent Kleene closures and a dot metacharacter",
			args: args{
				regexp: []rune("ab*.c*d"),
				text:   []rune("abbbbccd"),
			},
			want: true,
		},
		{
			name: "Kleene closure of dot metacharacter starting regex",
			args: args{
				regexp: []rune(".*bc"),
				text:   []rune("aAxyzbc"),
			},
			want: true,
		},
		{
			name: "Kleene closure of dot metacharacter ending regex 1",
			args: args{
				regexp: []rune("abc.*"),
				text:   []rune("abcABC"),
			},
			want: true,
		},
		{
			name: "Kleene closure of dot metacharacter ending regex 2",
			args: args{
				regexp: []rune("abc.*"),
				text:   []rune("abcabc"),
			},
			want: true,
		},
		{
			name: "Kleene closure of dot metacharacter ending regex 3",
			args: args{
				regexp: []rune("abc.*"),
				text:   []rune("abc"),
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := match(tt.args.regexp, tt.args.text); got != tt.want {
				t.Errorf("match() = %v, want %v", got, tt.want)
			}
		})
	}
}
