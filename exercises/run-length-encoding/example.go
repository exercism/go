package run_length_encoding

import (
	"fmt"
	"strconv"
)

func RunLengthEncode(s string)string{
	count:=1
	output:=""
	for i, c := range s{
		if i!=0{
			if rune(s[i-1])==c{
				count++
			}else{
				getEncodedOutput(&count,i,s,&output)
			}
		}
	}
	if len(s) != 0{
		getEncodedOutput(&count,len(s),s,&output)
	}
	return output
}

func RunLengthDecode(s string)string{
	count:=1
	stringCount:=""
	output:=""
	for _, c := range s{
		if _, err := strconv.Atoi(string(c)); err == nil {
			stringCount+=string(c);
		}else{
			if stringCount !=""{
				count, _ = strconv.Atoi(stringCount)
			}
			for j:=0;j<count;j++{
				output+=string(c)
			}
			count = 1
			stringCount = ""
		}
	}
	return output
}

func getEncodedOutput(count *int,i int, s string, output *string){
	if *count > 1{
		*output+= fmt.Sprintf("%d%c",*count,s[i-1])
		*count = 1
	}else{
		*output+= fmt.Sprintf("%c",s[i-1])
		*count = 1
	}
}