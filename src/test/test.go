package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"reflect"
	"strconv"
	"strings"

	"github.com/ranjith5/gomoda/tamilgreeting"
)

func main() {
	Greet()
}

func Greet() {
	tamilgreeting.Greet()
}

func Loops() {
	arr := [5]string{"A", "B", "C", "D", "E"}

	fmt.Println("normal for")
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}
	fmt.Println("full range for")
	for i, value := range arr {
		fmt.Println(i, " ", value)
	}

	fmt.Println("index range for")
	for value := range arr {
		fmt.Println(value)
	}

	fmt.Println("just range for")
	for range arr {
		fmt.Println("test")
	}
}

func PointerExample() {
	var myInt int = 10
	var myIntPointer *int = &myInt
	anotherMyIntPointer := &myInt
	fmt.Printf("variable value %v, variable type %v, pointer value %v, pointer type %v \n", myInt, reflect.TypeOf(myInt), *myIntPointer, reflect.TypeOf(myIntPointer))
	*myIntPointer = 15
	fmt.Printf("variable value %v, variable type %v, pointer value %v, pointer type %v \n", myInt, reflect.TypeOf(myInt), *myIntPointer, reflect.TypeOf(myIntPointer))
	fmt.Println(myInt)
	fmt.Println(&myInt)
	fmt.Println(anotherMyIntPointer)
}

//fmt.Print(RepeatLines("Ranjith", 5))
func RepeatLines(line string, times int) (returnText string, err error) {
	if times < 0 {
		return "", fmt.Errorf("times cannot be a number less than 0")
	}
	for x := 0; x <= times; x++ {
		fmt.Println(line)
	}
	return "done", nil
}

func FormatVerbs() {
	fmt.Printf("the number is %.5f \n", .2154214)
	fmt.Printf("the number is %10.5f \n", .2154214)
	fmt.Printf("the string is %s \n", "10.252154")
	fmt.Printf("the type is %v \n", 10.251245)
}

func GuessGame() {
	randomNumber := rand.Intn(100) + 1
	//fmt.Println("type of randomNumber", reflect.TypeOf(randomNumber))
	fmt.Println("GUESS THE NUMBER GAME")
	reader := bufio.NewReader(os.Stdin)
	guessed := false
	for x := 1; x < 4; x++ {
		fmt.Println("Enter your guess no ", x, ":")
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		input = strings.TrimSpace(input)
		inputNumber, err := strconv.Atoi(input)
		if err != nil {
			log.Fatal(err)
		}
		if inputNumber > randomNumber {
			fmt.Println("Guessed number is greater")
			continue
		} else if inputNumber < randomNumber {
			fmt.Println("Guessed number is lesser")
			continue
		} else {
			fmt.Println("Guessed number is correct")
			break
		}
	}
	if !guessed {
		fmt.Println("Sorry. You did not guess.")
	}
}