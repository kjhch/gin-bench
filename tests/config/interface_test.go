package config

import (
	"fmt"
	"testing"
)

type MyI interface {
	AddI(string) string
}

func RunMyI(i MyI) {
	fmt.Println(i.AddI("RunMyI"))
}

type MyI1 struct {
	str string
}

func (i *MyI1) AddI(str string) string {
	return str + " " + i.str
}

type MyI2 struct {
	str string
}

func (i MyI2) AddI(str string) string {
	return str + " " + i.str
}

func TestI(t *testing.T) {
	//RunMyI(MyI1{str: "I1"})
	RunMyI(&MyI1{str: "I1"})
	RunMyI(MyI2{str: "I2"})
	RunMyI(&MyI2{str: "I2"})

}
