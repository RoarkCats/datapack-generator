package main

import "fmt"

func InputType[T any](prefix string, vartype string) (out T) {

	fmt.Print(prefix)

	_, err := fmt.Scanf(vartype+"\n", &out)
	if err != nil && err.Error() != "unexpected newline" {
		panic(err)
	}
	return out
}

func Input(prefix string) (out string) {

	return fmt.Sprint(InputType[string](prefix, "%s"))
}
