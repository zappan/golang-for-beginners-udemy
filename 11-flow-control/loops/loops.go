package loops

import (
  "bufio"
  "fmt"
  "math/rand"
  "os"
  "strings"
  "time"
)

// three component loop
// (for? ... well Go only has a for loop, so this is probably the typical for loop from other langs)
func ThreeComponentLoop() {
  for i := 0; i < 10; i++ {
    fmt.Printf("%d ", i)
  }
  fmt.Println()
}

// while loop
func WhileLoop() {
  i := 1000
  rand.Seed(time.Now().UnixNano())

  for i > 100 {
    // i /= 10
    i = rand.Intn(1000)
    fmt.Printf("%d ", i)
  }

  fmt.Println()
  fmt.Println("Got", i, "and broke out of the loop")
}

// do-while loop
func DoWhileLoop() {
  reader := bufio.NewReader(os.Stdin)

  fmt.Println("[DoWhileLoop()] Enter something")
  for input := " "; input != "q" && input != "Q"; {
    fmt.Print("-> ")
    input, _ = reader.ReadString('\n')
    input = strings.Replace(input, "\n", "", -1)
    fmt.Println("You entered:", input)
  }
  fmt.Println("You exited the do-while loop...")
}

// infinite loop
func myLogger(ch chan string) {
  for {
    msg := <- ch
    fmt.Print("[myLogger] User input: ", msg)
  }
}

func InfiniteLoop() {
  // read input from the user 5 times and write it to a log
  reader := bufio.NewReader(os.Stdin)
  ch := make(chan string)

  go myLogger(ch)

  fmt.Println("[InfiniteLoop()] Enter something")
  for i := 0; i < 2; i++ {
    fmt.Print("-> ")
    input, _ := reader.ReadString('\n')
    ch <- input
    time.Sleep(time.Millisecond)  // just so that mylogger thread has the time to do a log output
  }
}

// nested loops
func NestedLoop() {
  fmt.Println("[NestedLoop()]")
  for i := 1; i < 5; i++ {
    fmt.Print("i is ", i)

    for j := 1; j < 4; j++ {
      fmt.Print("   j:", j)
    }
    fmt.Println()
  }

  fmt.Println()
}

// for each .. range
func ForEachRange() {}
