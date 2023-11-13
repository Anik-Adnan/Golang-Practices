package main

import (
	"fmt"
	"strings"
)

var pl = fmt.Println

func main() {
	s1 := "A word of string"
	replacer := strings.NewReplacer("o", "oo")
	s2 := replacer.Replace(s1)
	pl(s2)
	pl("length : ", len(s2))

	// find string contains a substring or not
	pl("Contain another : ", strings.Contains(s2, "Ano"))
	pl("o index : ", strings.Index(s2, "o"))           // find index
	pl("replace: ", strings.Replace(s2, "oo", "o", 2)) // string replaced

	// trim space
	trim_str := "\n\t\t  Hello !!\t\n"
	trim_str = strings.TrimSpace(trim_str)
	pl("trim space:", trim_str)

	mail := "anikadnanict17@gmail.com"
	pl("Split: ", strings.Split(mail, "@")) // split by adelimiter @
	pl("lower: ", strings.ToLower(mail))
	pl("Upper: ", strings.ToUpper(mail))
	pl("Title: ", strings.Title(s1))
	pl("prefix: ", strings.HasPrefix(mail, "anik"))
	pl("suffix: ", strings.HasSuffix(mail, ".com"))
	pl("cmpare : ", strings.Compare("aRe", "are"))
	pl("count : ", strings.Count(mail, "a"))          // count a
	pl("Last index : ", strings.LastIndex(mail, "c")) // last index of c

}
