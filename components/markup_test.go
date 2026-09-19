package components

import (
	"testing"
)

func Test_inlineLinks(t *testing.T) {
	type args struct {
		markup string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Empty string",
			args: args{markup: ""},
			want: "",
		},
		{
			name: "No matches",
			args: args{markup: "This is a text without any links."},
			want: "This is a text without any links.",
		},
		{
			name: "Single match",
			args: args{markup: "Click [here](http://example.com) for more information."},
			want: "Click <a href=\"http://example.com\">here</a> for more information.",
		},
		{
			name: "Multiple matches",
			args: args{markup: "Visit [Google](http://google.com) and [GitHub](http://github.com) for more."},
			want: "Visit <a href=\"http://google.com\">Google</a> and <a href=\"http://github.com\">GitHub</a> for more.",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := inlineLinks(tt.args.markup); got != tt.want {
				t.Errorf("inlineLinks() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_stripMarkup(t *testing.T) {
	type args struct {
		markup string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Empty string",
			args: args{markup: ""},
			want: "",
		},
		{
			name: "No matches",
			args: args{markup: "This is a text without any links."},
			want: "This is a text without any links.",
		},
		{
			name: "Single match",
			args: args{markup: "Click [here](http://example.com) for more information."},
			want: "Click here http://example.com for more information.",
		},
		{
			name: "Multiple matches",
			args: args{markup: "Visit [Google](http://google.com) and [GitHub](http://github.com) for more."},
			want: "Visit Google http://google.com and GitHub http://github.com for more.",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := stripMarkup(tt.args.markup); got != tt.want {
				t.Errorf("stripMarkup() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_inlineMarkup_EscapesHTML(t *testing.T) {
	tests := []struct {
		name   string
		markup string
		want   string
	}{
		{
			name:   "Tags",
			markup: `<script>alert("hi")</script>`,
			want:   `&lt;script&gt;alert(&quot;hi&quot;)&lt;/script&gt;`,
		},
		{
			name:   "Ampersand and apostrophe",
			markup: "Tom & Jerry's",
			want:   "Tom &amp; Jerry's",
		},
		{
			name:   "Markup around escaped text",
			markup: `**<b>** *"quoted"*`,
			want:   `<strong>&lt;b&gt;</strong> <em>&quot;quoted&quot;</em>`,
		},
		{
			name:   "Quote in link URL",
			markup: `[link](https://example.com/?a=1&b="2")`,
			want:   `<a href="https://example.com/?a=1&amp;b=&quot;2&quot;">link</a>`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := inlineMarkup(tt.markup); got != tt.want {
				t.Errorf("inlineMarkup() = %v, want %v", got, tt.want)
			}
		})
	}
}
