package main

import (
	"fmt"
	"reflect"

	"github.com/stakkato95/frequency/frequency"
)

func main() {
	result := frequency.Frequency("My choice title Boring Go  properly written Go boring It weird write boring topic I explain Go small feature set step modern programming lan guages Wellwritten Go programs tend straightforward sometimes repetitive Theres inheritance generics  aspectoriented programming function overloading certainly operator overloading Theres pattern matching named parameters exceptions To horror many pointers Gos concurrency model unlike languages its based ideas 1970s algorithm used garbage collector In short Go feels like throwback And thats point Boring trivial Using Go correctly requires understanding features intended fit together While write Go code looks like Java Python youre going unhappy result wonder fuss  Thats comes  It walks features Go explaining best write idiomatic code grow When comes building things  boring great No wants person drive bridge built untested techniques engineer cool The modern depends software depends bridges perhaps  Yet many programming languages add features without thinking impact maintainability codebase Go intended building programs  programs modified dozens developers dozens years Go boring thats fantastic I hope teaches build exciting projects boring code")
	fmt.Println(result)

	// smallRefTypesPlayground()
}

func smallRefTypesPlayground() {
	fmt.Println(reflect.DeepEqual(
		map[string]bool{"d": true, "a": true, "c": true, "b": true},
		map[string]bool{"a": true, "b": true, "c": true, "d": true},
	))

	my := [2]string{"a", "b"}
	ar(my)
	fmt.Println(my)

	mySlc := []string{"a", "b", "c", "d"}
	sl(mySlc)
	fmt.Println(mySlc)

	mySlc1 := []string{"a", "b", "c"}
	slAppend(&mySlc1)
	fmt.Println(mySlc1)
}

func ar(array [2]string) {
	array[0] = "100500"
}

func sl(slc []string) {
	slc[0] = "100500"
}

func slAppend(slc *[]string) {
	(*slc)[1] = "111111111"
	*slc = append(*slc, "900600")
}
