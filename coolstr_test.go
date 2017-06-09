package coolstr

import "testing"

func Test_coolStr(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			"Empty test",
			args{
				input: "\r\n",
			},
			0,
		},
		{
			"Simple word test",
			args{
				input: "This",
			},
			1,
		},
		{
			"Simple sentence test",
			args{
				input: "This is a word.",
			},
			4,
		},
		{
			"Sentence with number",
			args{
				input: "There are 100 words",
			},
			4,
		},
		{
			"Sentence with name",
			args{
				input: "That is Mr.Zhang Mr. Zhang",
			},
			4,
		},
		{
			"Sentence with wrap",
			args{
				input: "That is Mr.Zhang\r\nMr. Zhang",
			},
			4,
		},
		{
			"Sentence with word wraped",
			args{
				input: "That is sub-\r\nway subway",
			},
			4,
		},
		{
			"Sentence with phone number",
			args{
				input: "The number is 135-1353-1353 135-1353-1353",
			},
			4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := coolStr(tt.args.input); got != tt.want {
				t.Errorf("coolStr() = %v, want %v", got, tt.want)
			}
		})
	}
}
