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
				input: "",
			},
			0,
		},
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
			"simple sentence symbol",
			args{
				input: "This is   ~!@#$%^&*()_+=",
			},
			2,
		},
		{
			"sentence with number",
			args{
				input: "There are 100 words",
			},
			4,
		},
		{
			"sentence with decimal",
			args{
				input: "The price are 10.25 * 1025 ",
			},
			5,
		},
		{
			"percent",
			args{
				input: "10.25% 10.25",
			},
			2,
		},
		{
			"money",
			args{
				input: "10.25 ¥10.25",
			},
			2,
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
				input: "The number is 135-1353-1353 13513531353",
			},
			4,
		},
		{
			"sentence with comma",
			args{
				input: "A B, B",
			},
			2,
		},
		{
			"sentence with dot",
			args{
				input: "A is No. 1. No one better than him.",
			},
			9,
		},
		{
			"sentence with dash",
			args{
				input: "A is a well-known person which is well known.",
			},
			7,
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
