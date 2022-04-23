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

// randSliceMaker creates a random slice of range 3,12
func randSliceMaker() []int {
    length := randNum(3, 12)
    var s []int
    for i := 1; i <= length; i ++ {
        s = append(s, randNum(0, 10))
    }  
    return s
}

// sliceExpr() generates a random slice expression like [2:3], or [:4] or [2:]
func sliceExpr(length int) []int {
    // Generate a percent
    p := randNum(1, 100)
    var sliceToGuess []int
    switch {
        // One in five chance of slice expr being like [:num] or [num:]
        // We don't actually use the zero values, they're just place holders for
        // nothing on the other side of the : separator in the expr
        case p <= 20:
            leftOrRight := randNum(1, 2) 
            if leftOrRight == 1 {
                // [:num]
                // Append a zero placeholder then an actual number
                sliceToGuess = append(sliceToGuess, 0) 
                sliceToGuess = append(sliceToGuess, randNum(1, length))
            } else {
                // [num:]
                // Append an actual number then a zero placeholder
                sliceToGuess = append(sliceToGuess, randNum(1, length))
                sliceToGuess = append(sliceToGuess, 0) 
            }
        // Put an actual number on both sides of the : separator
        case p > 21:
            firstIndex := randNum(2, length)
            zeroIndex := randNum(1, firstIndex)
            sliceToGuess = append(sliceToGuess, zeroIndex)
            sliceToGuess = append(sliceToGuess, firstIndex)
            
    }
    return sliceToGuess
}

// convertToIntSlice converts user answer, a string, to an int slice to match against answer
func convertToIntSlice(userAnswer string) []int {
    var formattedUserAnswer []int
    for _, v := range userAnswer {
        intFromRune, err := strconv.Atoi(string(v))
        if err != nil {
            // It wasn't an int, so let's not add it do our int slice
            // probably a [, ], or comma, or .. whatever
            continue
        }
        formattedUserAnswer = append(formattedUserAnswer, intFromRune)
    }
    return formattedUserAnswer
}

// Compares two int slices to see if they're the same
func compareIntSlices(a, b []int) bool {
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
    firstRun := true

    // Run the guessing game forever
    for {
        // Visual starting again cue if not first run
        if !firstRun {
            fmt.Printf("\nStarting again!\n")
        }
        firstRun = false

        // Make a new random slice
        s := randSliceMaker()
        var answer []int
        var promptString string
        sliceToGuess := sliceExpr(len(s))

        // Handle the different slice expression types to generate the question and answer
        switch {
            // expr looks like [:num]
            case sliceToGuess[0] == 0:
                rightSideNumber := sliceToGuess[1]
                answer = s[:rightSideNumber]
                promptString = fmt.Sprintf("What is the result of [:%d]", sliceToGuess[1])
            // expr looks like [num:]
            case sliceToGuess[1] == 0:
                leftSideNumber := sliceToGuess[0]
                answer = s[leftSideNumber:]
                promptString = fmt.Sprintf("What is the result of [%d:]", sliceToGuess[0])
            // expr looks like [num:num]
            default:
                answer = s[sliceToGuess[0]:sliceToGuess[1]]
                promptString  = fmt.Sprintf("What is the result of [%d:%d]", sliceToGuess[0], sliceToGuess[1])
        } 

        // Add a little spacing between attempts
        if !firstRun {
            fmt.Println()
        }

        // Print the slice and ask user for answer
        fmt.Println(s)
        promptString += " (Use '?' to skip [and see answer]):\n"
        fmt.Printf(promptString)
    
        // var to hold user answer
        var userInput string
        fmt.Scanln(&userInput)

        // Convert user string answer to []int, dumping non ints along the way
        userFormattedAnswer := convertToIntSlice(userInput)

        // Check the answer
        same := compareIntSlices(userFormattedAnswer, answer)

        // Guessed right
        if same {
            fmt.Printf("Nice! You guessed %v correctly!\n", answer)
        }
        
        // User entered ?, let's skip this one
        if userInput == "?" {
            fmt.Printf("Answer was %v\n", answer)
            continue
        }

        // Guessed wrong, keep asking until correct or entered ? to skip
        for !same {
            fmt.Printf("Hmm, %v wasn't right. Try again:\n", userFormattedAnswer)
            // Print the slice and prompt again
            fmt.Println(s)
            fmt.Printf(promptString) 
            fmt.Scanln(&userInput) 
            userFormattedAnswer = convertToIntSlice(userInput)

            // User entered ?, let's skip this one
            if userInput == "?" {
                fmt.Printf("Answer was %v\n", answer)
                break
            }

            // Check the answer
            same := compareIntSlices(userFormattedAnswer, answer)
            // Guessed right
            if same {
                fmt.Printf("Nice! You guessed %v correctly!\n", answer)
                break
            }
        }
    }
}
