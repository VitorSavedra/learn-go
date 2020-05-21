package main

import "fmt"

func f1() {
	fmt.Println("Hello. :)")
}

func f2(p1 string, p2 string) {
	fmt.Printf("%s, %s.\n", p1, p2)
}

func f3() string {
	return "Hello?!\n"
}

func f4(p1, p2 string) string {
	return fmt.Sprintf("%s How are you? %s", p1, p2)
}

func f5() (string, string) {
	return "Toc", "Toc"
}

func main() {
	f1()
	f2("Hello", "again")

	r3, r4 := f3(), f4("Ok...", "<_<")
	fmt.Println(r3, r4)

	r51, r52 := f5()
	fmt.Println(r51)
	fmt.Println(r52)
}
