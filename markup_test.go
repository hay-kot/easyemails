package easyemails

import "testing"

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
