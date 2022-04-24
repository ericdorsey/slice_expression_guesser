package main

import (
    "fmt"
    "math/rand"
    "time"
    "strconv"
    "os"
    "strings"
)

// randNum generates a random number between min and max
func randNum(min int, max int) int {
    v := rand.Intn(max - min) + min
    return v
}

// randSliceMaker creates a random slice, length between minLength and maxLength
func randSliceMaker(minLength, maxLength int) []int {
    length := randNum(minLength, maxLength)
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
    var sliceExprToGuess []int
    switch {
        // One in five chance of slice expr being like [:num] or [num:]
        // We don't actually use the zero values, they're just place holders for
        // nothing on the other side of the : separator in the expr
        case p <= 20:
            leftOrRight := randNum(1, 2) 
            if leftOrRight == 1 {
                // [:num]
                // Append a zero placeholder then an actual number
                sliceExprToGuess = append(sliceExprToGuess, 0) 
                sliceExprToGuess = append(sliceExprToGuess, randNum(1, length))
            } else {
                // [num:]
                // Append an actual number then a zero placeholder
                sliceExprToGuess = append(sliceExprToGuess, randNum(1, length))
                sliceExprToGuess = append(sliceExprToGuess, 0) 
            }
        // Put an actual number on both sides of the : separator
        case p >= 21:
            firstIndex := randNum(2, length)
            zeroIndex := randNum(1, firstIndex)
            sliceExprToGuess = append(sliceExprToGuess, zeroIndex)
            sliceExprToGuess = append(sliceExprToGuess, firstIndex)
            
    }
    return sliceExprToGuess
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

    // vars to track results
    correct := 0
    incorrect := 0
    skipped := 0

    // Run the guessing game forever
    for {
        // Visual starting again cue if not first run
        if !firstRun {
            fmt.Printf("\nStarting again!\n")
        }
        firstRun = false

        // Make a new random slice
        s := randSliceMaker(3, 12)
        var answer []int
        var promptString string
        sliceExprToGuess := sliceExpr(len(s))

        // Handle the different slice expression types to generate the question and answer
        switch {
            // expr looks like [:num]
            case sliceExprToGuess[0] == 0:
                rightSideNumber := sliceExprToGuess[1]
                answer = s[:rightSideNumber]
                promptString = fmt.Sprintf("What is the result of [:%d]", sliceExprToGuess[1])
            // expr looks like [num:]
            case sliceExprToGuess[1] == 0:
                leftSideNumber := sliceExprToGuess[0]
                answer = s[leftSideNumber:]
                promptString = fmt.Sprintf("What is the result of [%d:]", sliceExprToGuess[0])
            // expr looks like [num:num]
            default:
                answer = s[sliceExprToGuess[0]:sliceExprToGuess[1]]
                promptString  = fmt.Sprintf("What is the result of [%d:%d]", sliceExprToGuess[0], sliceExprToGuess[1])
        } 

        // Add helfpul user options to prompt
        promptString += " (?=skip [and see answer], q=quit):\n"

        // To hold answer user types in
        var userInput string

        // Prompt this same question until user skips with ?, quits with q, or gets answer right
    
        // Track if already guessed wrong once
        incorrectOnce := false

        for {

            fmt.Println(s)
            fmt.Printf(promptString) 
            fmt.Scanln(&userInput) 
            userFormattedAnswer := convertToIntSlice(userInput)

            // User wants to quit
            if strings.HasPrefix(strings.ToLower(userInput), "q") {
                fmt.Printf("Final score -- correct %d, incorrect %d, skipped %d\n", correct, incorrect, skipped)
                fmt.Printf("Quitting.\n")
                os.Exit(0)
            }
            // User entered ?, let's skip this one
            if userInput == "?" {
                fmt.Printf("Skipping; answer was %v\n", answer)
                if !incorrectOnce {
                    skipped += 1
                }
                // Break out to outer loop to get next question
                break
            }

            // Check the answer
            same := compareIntSlices(userFormattedAnswer, answer)

            // Guessed right
            if same {
                if !incorrectOnce {
                    correct += 1
                    fmt.Printf("Nice! You guessed %v correctly!\n", answer)
                } else {
                    fmt.Printf("Nice! You guessed %v correctly! (doesn't count towards correct score since you guessed wrong once)\n", answer)
                }
                // Break out to outer loop to get next question
                break
            }
            // Guessed wrong
            if !same {
                if !incorrectOnce {
                    incorrect += 1
                }
                incorrectOnce = true
                fmt.Printf("Hmm, %v wasn't right. Try again:\n", userFormattedAnswer) 
            }
        }
    }
}
