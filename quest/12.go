package quest

import (
	"fmt"
	"strings"
)

//字符          数值
//I             1
//V             5
//X             10
//L             50
//C             100
//D             500
//M             1000


func intToRoman(num int) string {
	N1000:=num/1000
	num = num-N1000*1000
	N100:= num/100
	num = num-N100*100
	N10:=num/10
	num = num-N10*10
	N1:=num
	builder := strings.Builder{}
	fmt.Println(N1000,N100,N10,N1)

	for N1000>0{
		builder.WriteRune('M')
		N1000--
	}
	for N100>0{
		if N100==9{
			builder.WriteString("CM")
			N100-=9
		} else if N100>=5{
			builder.WriteRune('D')
			N100-=5
		}else if N100==4{
			builder.WriteString("CD")
			N100-=4
		}else {
			builder.WriteRune('C')
			N100--
		}
	}

	for N10>0{
		if N10==9{
			builder.WriteString("XC")
			N10-=9
		} else if N10>=5{
			builder.WriteRune('L')
			N10-=5
		}else if N10==4{
			builder.WriteString("XL")
			N10-=4
		}else {
			builder.WriteRune('X')
			N10--
		}
	}

	for N1>0{
		if N1==9{
			builder.WriteString("IX")
			N1-=9
		} else if N1>=5{
			builder.WriteRune('V')
			N1-=5
		}else if N1==4{
			builder.WriteString("IV")
			N1-=4
		}else {
			builder.WriteRune('I')
			N1--
		}
	}
	return builder.String()
}
