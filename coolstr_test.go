package coolstr

import "testing"

func Test_coolStr(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"empty test",
			args{
				input: "\r\n",
			},
			0,
		},
		{
			"simple word test",
			args{
				input: "This     ",
			},
			1,
		},
		{
			"simple sentence test",
			args{
				input: "This is      a word.",
			},
			4,
		},
		{
			"sentence with number",
			args{
				input: "There are 100 words",
			},
			4,
		},
		{
			"sentence with name",
			args{
				input: "That is Mr.Zhang Mr. Zhang",
			},
			5,
		},
		{
			"sentence with newline",
			args{
				input: "That is Mr.Zhang\nMr. Zhang",
			},
			5,
		},
		{
			"sentence with newline and return",
			args{
				input: "That is Mr.Zhang\r\nMr. Zhang",
			},
			5,
		},
		{
			"sentence with word wraped",
			args{
				input: "That is sub-\r\nway subway",
			},
			3,
		},
		{
			"sentence with word wraped(2)",
			args{
				input: "That is sub-\nway subway",
			},
			3,
		},
		{
			"sentence with word wraped(3)",
			args{
				input: "That is No. 123-\n456 subway",
			},
			5,
		},
		{
			"sentence with phone number",
			args{
				input: "The number is 135-1353-1353 135-1353-1353",
			},
			4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CoolStr(tt.args.input); got != tt.want {
				t.Errorf("coolStr() = %v, want %v", got, tt.want)
			}
		})
	}
}
