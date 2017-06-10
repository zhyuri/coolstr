# Coolstr [![Build Status](https://travis-ci.org/zhyuri/coolstr.svg?branch=master)](https://travis-ci.org/zhyuri/coolstr)

coolstr is just a tiny function which can calculate how many different words are there in a string.

it would be easy to cut a sentense by space and line wrap, but it is difficult to decide whether the word you cut out is the same as the former one. 

## User Case

there are some declaration we need to make first.

| A               | B               | Rule      |
| --------------- | --------------- | --------- |
| Mr.Zhang        | Mr.(space)Zhang | Different |
| Mr.Zhang\r\nXXX | Mr.ZhangXXX     | Same      |
| Mr.Zhang\nXXX   | Mr.ZhangXXX     | Same      |
| sub-\r\nway     | subway          | Same      |
| sub-\nway       | subway          | Same      |
| well-known      | well known      | Different |
| No.             | No              | Different |
| 123-\n456       | 123456          | Same      |
| 135-1353-1353   | 13513531353     | Same      |
| 10.25           | 1025            | Different |
| 10.25%          | 10.25           | Different |
| ¥10.25          | 10.25           | Different |

## Doc

Function signature:

```go
func CoolStr(input string) int;
```

- input string

  The input string get all words you want to count

- output int

  The word count

## Example

see `coolstr_test.go`