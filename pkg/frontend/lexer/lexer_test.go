package lexer

import (
	"fmt"
	"testing"
)

func TestString(t *testing.T) {
	sql := "create procedure proc04()" + "\n" +
		"begin" + "\n" +
		"	declare var int(10);" + "\n" +
		"	set var=0;" + "\n" +
		"	while var < 6 do" + "\n" +
		"		insert into t values(var); " + "\n" +
		"		set var=var+1;" + "\n" +
		"	end while;" + "\n" +
		"end;"
	source := NewStringSource(sql)
	fmt.Println(source.sourcStr)

}
