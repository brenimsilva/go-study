package main

import (
	"errors"
	"fmt"
	"math/rand"
	"sort"
	"strings"
	"sync"
	"time"
	"unicode/utf8"
)

func main() {
    // partOne()
    // partTwo()
    // partThree()
    // partFour()
    // partFive()
    // partSix()
    // partEight()
    // partNine()
    partTen()
}


var results = []string{}
var mutex = sync.Mutex{}
var wg = sync.WaitGroup{}
var dbData = []string{"id1","id2","id3","id4","id5"}
//GO Routines

func partTen() {
    t0 := time.Now()
    for i := 0; i < len(dbData); i++ {
        wg.Add(1)
        go dbCall(i)
    }

    wg.Wait()

    fmt.Printf("\nTotal execution time: %v", time.Since(t0))
    fmt.Printf("\nThe results are %v", results)
}

func dbCall(i int) {
    //Simulate dbcall delay
    var delay float64 = rand.Float64()*2000
    time.Sleep(time.Duration(delay)*time.Millisecond)
    fmt.Println("The result from the database is: ", dbData[i])
    mutex.Lock()
    results = append(results, dbData[i])
    mutex.Unlock()
    wg.Done()

}


func partNine() {
    var testUser user = user{"Breno", 30, bankAccount{"Itau", 100}} // one way of declaring
    var testUser2 user = user{age: 27, bank: bankAccount{name: "Nubank", balance: 200}, name: "Thamyres"} // another way of declaring
    fmt.Println(testUser)
    fmt.Println(testUser2)

    var structNotReusable = struct{
        name string
        age int
    }{"Breno",30}
    fmt.Println(structNotReusable)

    var animal1 = carnivoro{"Cachorro"}
    var animal2 = herbivoro{"Elefante"}
    comecome(animal1)
    comecome(animal2)

}

func comecome(a animal) {
    a.comer()
}

type user struct{
    name string
    age int
    bank bankAccount
}

type bankAccount struct{
    name string
    balance float64
}

type carnivoro struct {
    race string
}

type herbivoro struct {
    race string
}

type animal interface{
    comer()
}

func (a carnivoro) comer() {
    fmt.Printf("%s esta comendo carne\n", a.race)
}

func (a herbivoro) comer() {
    fmt.Printf("%s esta comendo folhas\n", a.race)
}

func partEight() {
    var n int = 1000000
    var testSlice = []int{}
    var testSlice2 = make([]int,0,n) // First is values created (len) second is capacity (cap)
    fmt.Printf("Total time without preallocation: %v\n", timeLoop(testSlice, n))
    fmt.Printf("Total time without preallocation: %v\n", timeLoop(testSlice2, n))
}

func timeLoop(slice []int, n int) time.Duration {
    t0 := time.Now()

    for(len(slice) < n) {
        slice = append(slice, 1)
    }

    return time.Since(t0)

}

func partSeven() {
    intArr := [...]int{1,2,3} // if you dont want to write the exact number of elements
    fmt.Println(intArr)
    intSlice := []int{4,5,6}
    intSliceToAppend := make([]int, 3) // using make function gains 
    intSliceToAppend[0] = 7
    intSliceToAppend[1] = 8
    intSliceToAppend[2] = 9
    intSlice = append(intSlice, intSliceToAppend...) // Spread comes after
    fmt.Println(intSlice)

}

func partSix() {
    var myBill bill = newBill("Breno's bill")
    fmt.Println(myBill)
    myBill.setName("Breno's new bill")
    fmt.Println(myBill)

    var newBill bill = createBill()

    fmt.Println(newBill)
}

func createBill() bill {
    var reader myReader = createReader()
    name, _ := reader.getInput("Create bill name: ")
    newBill := newBill(name)
    fmt.Println("Created bill - ", newBill.name)

    return newBill
}

func partFive() {
    name := "Breno"
    // pName := &name
    var pName *string = &name

    fmt.Println("Memory address of name is: ", pName)
    fmt.Println("Value at memory address", *pName)

}

func partFour() {
    menu := map[string]float64{
        "pizza": 55.20,
        "lasagna": 24.75,
    }

    fmt.Println(menu)
    fmt.Println(menu["pizza"])
    fmt.Println(menu["lasagna"])
    fmt.Println(menu["undefined"])

    for key, value := range menu {
        fmt.Println(key, "-", value)
    }
}

func partThree() {
    // x := 0
    // for x < 3 {
    //     x++
    //     fmt.Println(x)
    // }
    // for i := 0; i < 3; i++ {
    //     fmt.Println("Value of I is:", i)
    // }

    //Loops
    names := []string{"Breno", "Thamyres", "Peggy", "Snow"}

    // for i := 0; i < len(names); i++ {
    //     fmt.Println(names[i])
    // }

    for index, value := range names {
        fmt.Printf("The name is %v and index %v\n", value, index)
    }

    ages := []int{10,12,18,5,30,17,32,44}
    var above []int = filter(ages, aboveEighteen)
    fmt.Println(above)
    var fn, ln = getInitials("Breno")
    fmt.Println(fn,ln)
    var test string
    fmt.Scanf("%s", test)
}

func getInitials(n string) (string, string) {
    nameUpper := strings.ToUpper(n)
    names := strings.Split(nameUpper, " ")
    var initials []string
    if(len(names) == 1) {
        return names[0][:1], names[0][len(names[0])-1:len(names[0])]
    }
    for _,name := range names {
        initials = append(initials, name[:1])
    }
    return initials[0], initials[1]
}

func aboveEighteen(age int) bool {
    return age >= 18
}

func filter(s []int, f func(int) bool) []int{
    filtered := []int{}
    for _, value := range s {
        if(f(value)) {
            filtered = append(filtered, value)
        }
    }
    return filtered
}

func partTwo() {
    //Arrays
    var ages = [3]int{10,20,30}
    fmt.Println(ages)
    names := [2]string{"Breno","Thamyres"}
    names[1] = "Test" // pode ser alterado posteriormente
    fmt.Println(names)
    fmt.Println("================================")

    //Slices
    scores := []int{100,50,60}
    fmt.Println(scores, len(scores))
    scores = append(scores, 85) // Append retorna um novo slice
    fmt.Println(scores, len(scores))
    fmt.Println("==============================")

    //Strings
    greetings := "Hello, i'm Breno!"
    fmt.Println(strings.Contains(strings.ToLower(greetings), "hello"))
    fmt.Println(strings.ReplaceAll(greetings, "Hello", "Hi"))
    nonFormatedCpf := "12345678910"
    formatedCpf := nonFormatedCpf[0:3] + "." + nonFormatedCpf[3:6] + "." + nonFormatedCpf[6:9] + "-" + nonFormatedCpf[9:]
    println(formatedCpf)
    fmt.Println("==============================")

    //Sort
    namesToSort := []string{"Breno", "Thamyres", "Peggy"}
    sort.Strings(namesToSort)
    fmt.Println(namesToSort)
    agesToSort := []int{30,10,20,5,90,3,44}
    sort.Ints(agesToSort)
    fmt.Println(agesToSort)

    fmt.Println(sort.SearchStrings(namesToSort, "Thamyres"))
    fmt.Println(sort.SearchInts(agesToSort, 30))
    fmt.Println("==============================")
}

func partOne() {
    fmt.Println("Hello World!")
    var intNum int = 32000
    var floatNum float64 = 123.9
    var result float32 = float32(floatNum) + float32(intNum)
    fmt.Println(result)

    fmt.Println(utf8.RuneCountInString("yes")) // use para pegar o length
    var myRune rune = 'a'
    fmt.Println(myRune)

    var1, var2 := 1,"2"
    fmt.Println(var1,var2)

    var stringResult, err = printMe("test")
    print(stringResult, err.Error()) // Handling errors

    var arr []string
    var arr2 []int32 = []int32{4,5,6}
    arr2 = append(arr2, 11)
    var name string
    fmt.Scanf("%s", &name)
    fmt.Println(name)
    arr = append(arr, name)
    fmt.Println(arr)
    fmt.Println(len(arr), cap(arr))

    // I/O
    var inputNameOne string
    var inputNameTwo string
    var inputNameThree string
    fmt.Println("Insira o primeiro nome:")
    fmt.Scanf("%s", &inputNameOne)
    fmt.Println("Insira o segundo nome:")
    fmt.Scanf("%s", &inputNameTwo)
    fmt.Println("Insira o terceiro nome:")
    fmt.Scanf("%s", &inputNameThree)
    fmt.Printf("Your age is %s \n", inputNameTwo)
    fmt.Printf("Test float %0.2f", 223.55)
    returnString := fmt.Sprintf("Test float %0.2f", 223.55) // return the string
    fmt.Println(returnString)
}


func printMe(value string) (string, error){
    var error error

    if(value == "test") {
        error = errors.New("Um teste de erro")
        return value, error
    }
    return value, error
}
