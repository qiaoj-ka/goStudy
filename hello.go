package main

import "fmt"

//=========继承、封装、多态==========
func main() {
	//fmt.Println("hello world!")
	bark(&Dog{"小狗", 6, 0})
	//=======小猫小鸭子继承=======
	cat := Cat{Animal{"小猫"}}
	cat.shout()
	dark := Dark{Animal{"小鸭子"}}
	dark.shout()
	//=========小猫小狗接口======
	var cute AmphibiansAble
	cute = &cat
	cute.Climbing()
	cute = new(Dog)
	cute.Climbing()
}

//Dog===================
type Dog struct {
	Name string
	Age  int64
	Sex  int
}

func bark(dog *Dog) {
	fmt.Println(dog.Name + "汪汪汪")
}

//=======================
//======继承例子=========
type Animal struct {
	name string
}

type Cat struct {
	Animal //伪继承
}
type Dark struct {
	Animal //伪继承
}

func (a *Animal) shout() {
}

//方法重写
func (c *Cat) shout() {
	fmt.Println(c.name + "喵喵喵")
}
func (d *Dark) shout() {
	fmt.Println(d.name + "嘎嘎")
}

//==========加入接口======
type BirdAble interface {
	Flying()
}
type AmphibiansAble interface {
	Climbing()
}

func (d *Dark) Flying() {
	fmt.Println(d.name + "拥有一双翅膀可以飞翔")
}
func (d *Dog) Climbing() {
	fmt.Println("小狗拥有四肢行走")
}
func (c *Cat) Climbing() {
	fmt.Println(c.name + "拥有四肢行走")
}
