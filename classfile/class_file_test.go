package classfile

import (
	"fmt"
	"testing"

	"github.com/vzjxif/go-jvm/classpath"
)

func TestClassFile(t *testing.T) {
	cf := ClassFile{}
	fmt.Println(cf.interfaces)
}

func TestParser(t *testing.T) {
	//load class
	cp := classpath.Parse("/Library/Java/JavaVirtualMachines/jdk1.8.0_321.jdk/Contents/Home/jre", "")
	classData, _, err := cp.ReadClass("java/lang/String")
	if err != nil {
		panic(err)
	}
	Parse(classData)
}
