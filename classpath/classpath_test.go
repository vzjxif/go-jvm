package classpath

import (
	"fmt"
	"testing"
)

func TestJVM(t *testing.T) {
	cp := Parse("/Library/Java/JavaVirtualMachines/jdk1.8.0_321.jdk/Contents/Home/jre", "/Library/Java/JavaVirtualMachines/jdk1.8.0_321.jdk/Contents/Home/jre/lib/rt.jar")
	_, _, err := cp.ReadClass("java/lang/Object")
	if err != nil {
		fmt.Println(err)
	}
}
