package main

import (
	"fmt"
	"golang.org/x/exp/constraints"
)

/* Basic min function for float64 arguments
func min (x,y float64) float64 {
	 if x< y{
	 	return x
	 }
	 return y
}

*/

/* Generic min function constraint
The type parameter T, declared in a type parameter list, takes the place of float64
*/
func min[T constraints.Ordered](x, y T) T {
	if x < y {
		return x
	}
	return x
}

/* Type & Methods Constraints
any equals to empty interface, meaning every type. sintax: any= interface{}

T is type parameter
any is type constraint. any is an alias for interface{}
*/
type Vector [T any] []T


//only for testing purposes
func (v *Vector[T]) PushOne(x T) {
	*v = append(*v, x)
}

func (v *Vector[T]) PushMany(x ...T) {
	*v = append(*v, x...)
}


type myInt64 int64


type Number[T constraints.Signed] struct{
	number T
}

func (n * Number[T]) isPositive() bool{
	return n.number > 0
}

func (n * Number[T]) isNegative() bool{
	return !n.isPositive()
}

func (n * Number[T]) isZero() bool{
	return n.number==0
}




type OrderedHeritage interface {
	MyOrderedInterface
}

type MyOrderedInterface interface {
	hello()
	constraints.Integer
}


type myOrdered struct{
	name string
}

func (m *myOrdered) hello(){
	fmt.Println(fmt.Sprintf("Hello, I'm %s",m.name))
}




func main() {

	//Calling generic min . We provide type and ordinary arguments
	//m := min[int](2, 3) or
	m := min[int](2, 3)

	fmt.Println(fmt.Sprintf("Min %d", m))


	vector:= Vector[int]{}
	vector.PushOne(1)
	vector.PushMany(2,3,4,5,6)

	for i,e := range vector{
		fmt.Println(fmt.Sprintf("Vector item %d: value %d", i,e))

	}


	// We can use myInt64 as constraints.Signed 'cause ~int64 is defined as a set type.
	positiveNumber:= Number[myInt64]{number: 10}
	negativeNumber:= Number[myInt64]{number: -10}
	zeroNumber := Number[myInt64]{number: 0}

	fmt.Println(fmt.Sprintf("isPositive %v, isNegative %v, isZero %v",positiveNumber.isPositive(), positiveNumber.isNegative(), positiveNumber.isZero()))
	fmt.Println(fmt.Sprintf("isPositive %v, isNegative %v, isZero %v",negativeNumber.isPositive(), negativeNumber.isNegative(), negativeNumber.isZero()))
	fmt.Println(fmt.Sprintf("isPositive %v, isNegative %v, isZero %v", zeroNumber.isPositive(), zeroNumber.isNegative(), zeroNumber.isZero()))

	ordered1:= myOrdered{name: "Gabo"}

	ordered1.hello()

}
