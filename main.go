package main

import (
    "fmt"
    "math/rand"
    "time"
    "strconv"
)

// randNum generates a random number between min and max
func randNum(min int, max int) int {
    v := rand.Intn(max - min) + min
    return v
}

func randSliceMaker() []int {
    length := randNum(3, 12)
    var s []int
    for i := 1; i <= length; i ++ {
        s = append(s, randNum(0, 10))
    }  
    return s
}

// sliceExpr() generates a random slice expression like [2:3], or [:3] or [2:]
func sliceExpr(length int) []int {
    // Generate a percent
    p := randNum(1, 100)
    var sliceToGuess []int
    switch {
        case p <= 20:
            leftOrRight := randNum(1, 2) 
            if leftOrRight == 1 {
                // [:num]
                sliceToGuess = append(sliceToGuess, 0) 
                sliceToGuess = append(sliceToGuess, randNum(1, length))
            } else {
                // [num:]
                sliceToGuess = append(sliceToGuess, randNum(1, length))
                sliceToGuess = append(sliceToGuess, 0) 
            }
        case p > 21:
            firstIndex := randNum(2, length)
            zeroIndex := randNum(1, firstIndex)
            sliceToGuess = append(sliceToGuess, zeroIndex)
            sliceToGuess = append(sliceToGuess, firstIndex)
            
    }
    //fmt.Printf("sliceToGuess inside sliceExpr() is %v\n", sliceToGuess)
    return sliceToGuess
}

// formatAnswer converts a string to a slice of strings
func formatAnswer(answer string) []string {
    var formatted []string
    for _, v := range answer {
        formatted = append(formatted, string(v)) 
    }
    return formatted
}

// Compares two string slices to see if they're the same
func compareStringSlices(a, b []string) bool {
    if len(a) != len(b) {
        return false
    }
    
    for i, v := range a {
        if v != b[i] {
            return false
        }
    }
    return true
}

func main() {
    // Provide a random seed
    rand.Seed(time.Now().UnixNano())

    for {
        // Do stuff forever
        fmt.Printf("\nStarting again!\n")
        s := randSliceMaker()
        var answer []int
        var promptString string
        sliceToGuess := sliceExpr(len(s))

        switch {
            // [:num]
            case sliceToGuess[0] == 0:
                rightSideNumber := sliceToGuess[1]
                answer = s[:rightSideNumber]
                //promptString = fmt.Sprintf("What is the result of [:%d]\n", sliceToGuess[1])
                promptString = fmt.Sprintf("What is the result of [:%d]", sliceToGuess[1])
            // [num:]
            case sliceToGuess[1] == 0:
                leftSideNumber := sliceToGuess[0]
                answer = s[leftSideNumber:]
                //promptString = fmt.Sprintf("What is the result of [%d:]\n", sliceToGuess[0])
                promptString = fmt.Sprintf("What is the result of [%d:]", sliceToGuess[0])
            // [num:num]
            default:
                answer = s[sliceToGuess[0]:sliceToGuess[1]]
               //promptString  = fmt.Sprintf("What is the result of [%d:%d]\n", sliceToGuess[0], sliceToGuess[1])
                promptString  = fmt.Sprintf("What is the result of [%d:%d]", sliceToGuess[0], sliceToGuess[1])
        } 

        // Ask user for answer
        fmt.Println(s)
        promptString += " (Use '?' to skip and see answer):\n"
        fmt.Printf(promptString)
        var formattedAnswer []string

        for _, v := range answer {
            formattedAnswer = append(formattedAnswer, strconv.Itoa(v)) 
        }

        //fmt.Println(userFormattedAnswer) 
        var userInput string
        fmt.Scanln(&userInput)
        userFormattedAnswer := formatAnswer(userInput)
        /*
        if userInput == "?" {
            fmt.Printf("Answer was %v\n", answer)
        }
        */
        same := compareStringSlices(userFormattedAnswer, formattedAnswer)
        // Guessed right
        if same {
            fmt.Printf("Nice! You guessed %v correctly!\n", answer)
        }
        // Guessed wrong
        for !same {
            if userInput == "?" {
                fmt.Printf("Answer was %v\n", answer)
                break
            }
            fmt.Printf("Hmm, that wasn't right. Try again?\n")
            // Print the slice again
            fmt.Println(s)
            fmt.Printf(promptString) 
            fmt.Scanln(&userInput) 
            userFormattedAnswer := formatAnswer(userInput)
            same = compareStringSlices(userFormattedAnswer, formattedAnswer)
            // Guessed right
            if same {
                fmt.Printf("Nice! You guessed %v correctly!\n", answer)
                break
            }
        }
    }
}
