# ggenerics
Go Generics testing repository.


Problems to solve for

A way to provide paremtized types for structs, functions and methods.

A way for parameterized types to use common operators like '==', '<' and '>'.


# Language Features

1. Type parameters for functions and types
   Type Parameters lists
   [P, Q constraint1, R constraint2]
Type parameters lsits look like ordinary parameter lists with square brackets. it is customary to start tpype parameters with upper-case letters to emphasize that they are types.
   
2. Type sets defined by interfaces
3. Type inference




It's required to active golang experimental feature to use generics and use constraints experimental library. In the future, this will be added to the core pkg.


Concepts:

Constraints: are interfaces and can also include new concept called "typesets"
Typesets: the union of types allowed to use by the constraint



Parametized types:

Can be: Functions structs and methods

Interfaces: set of methods


Prior 1.18:

Interfaces defines methods sets
Defines

interface I {
a()
b()
c()
}


Implements 

type P
func(P) a()
func(P) b()
func(P) c()
func(P) d()



type Q
func(Q) a()
func(Q) b()
func(Q) c()
func(Q) z()


So P and Q implements interface I

Now... 1.18 TypeSet is the group of types that implements the interface. So P and Q are members of the interface typeset. 

I typeset: P and Q 

Also, in 1.18 Interfaces are being extended to support typesets

Example:

interface H {
int | int8 | int16| int32 | int64
}

H TypeSet: int, int8, int16, int32 and int64

int implements...


In this scenario, only the specified types in the interface are members of the interface typeset.


Type constraints are interfaces. All interfaces can, nut not necessarily should, be used as constraints.
The inverse is not true, not all constraints can be used as interfaces.

IE: constraints Signed:

type Signed interface {  // Define the set of types. No need to provide methods. Operators like less, plus, minus, etc because works for the types defined by this interface , any parametized type can use it.
~int | ~int8 | ~int16 | ~int32 | ~int64
}

To let open the posibility to custom types based on these parametized types, we  can use ~ operator.

This means type myInt64 is member of the Signet TypeSet because we define ~int64.




A diff between java and go, golang doesn't have type [erasure](https://www.baeldung.com/java-type-erasure). Golang preserve type information at runtime.
Another diff is that Golang doesn't need wildcards.





Generics with simplicity in mind

No specialization: 
No metraprogramming
No covariance/contravarience
No operator overloading


When sould I use generics?

When writing functions that work with container types of any element - including those built into stdlib: maps, slices, channels.

func Equal [ T comparable] (s1,s2 []T) bool {
   if len(s1) != len(s2) {
      return false
   }
   for i, e1 := range s1 {
      if e1 != s2[i] {
         return false
      }
   }
return true 
}

When a method looks the same for all types
Ex:
type SliceFN [T any] struct {
s [T]
cmp func (T,T) bool
}

func (s SliceFN [T]) Len() int { return len(s.s)}
func (s SliceFn[T]) Swap ( i , j int) {
s.s[i], s.s[j]= s.s[j], s.s[i]
}

func (s SliceFn[t]) Less (i, j int) bool {
return s.cmp(s.s[i], s.s[j])
}

func SortFn _[T any] (s [T], cmp func(T,T) bool){
sort.Sort(sliceFn[t]{s,cmp})
}


When should I not use generics?
When a simple interface can be used instead
When the implementation is different for each type
