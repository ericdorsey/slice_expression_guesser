package main

import (
    "testing"
    "fmt"
    "math/rand"
    "time"
)

// Provide a random seed so test slices and test randNums aren't the same every time
func init() {
    rand.Seed(time.Now().UnixNano())
}

// TestRandNum tests the randNum function
func TestRandNum(t *testing.T) {
    min := 1
    max := 10
    v := randNum(min, max)
    fmt.Printf("value v generated is %v\n", v)
    if v < min {
        t.Errorf("Expected minimum value of %d but have %d instead", min, v)
    }
    
    if v > max {
        t.Errorf("Expected maximum value of %d but have %d instead", max, v)
    }
}

// TestRandSliceMaker tests the randSliceMaker function
func TestRandSliceMaker(t *testing.T) {
    minLength := 3
    maxLength := 12
    s := randSliceMaker(minLength, maxLength)
    fmt.Printf("slice generated s is %v\n", s)
    if len(s) < minLength {
        t.Errorf("Expected minimum slice length of %d but have %d instead", minLength, len(s))
    }
    if len(s) > maxLength {
        t.Errorf("Expected maximum slice length of %d but have %d instead", maxLength, len(s))
    }
}

// TestSliceExpr test the sliceExpr function
func TestSliceExpr(t *testing.T) {
    // create a sliceExpr with max length of 8
    length := 8
    s := sliceExpr(length)
    fmt.Printf("slice expression s generated is %v\n", s)
    possibles := [][]int{
        []int{0, length}, 
        []int{length, 0}, 
        []int{length, length},
    }
    var correctFormat bool 
    for _, v := range possibles {
        // matches [0:someNumber]
        if (s[0] == v[0]) && (s[1] <= v[1]) {
            correctFormat = true
        } 
        // matches [someNumber:0]
        if (s[1] == v[1]) && (s[0] <= v[0]) {
            correctFormat = true
        }
        // matches [someNumber:someNumber]
        if (s[0] <= v[0]) && (s[1] <= v[1]) {
            correctFormat = true
        }
    }  
    if !correctFormat {
        t.Errorf("Expected one of [0:someNumber <= %d], [someNumber <= %d:0], [someNumber <= %d:someNumber <= %d] but got %v", length, length, length, length, s)
    }
}

// TestConvertToIntSlice tests the convertTo IntSlice function
func TestConvertToIntSlice(t *testing.T) {
    s := "[5, 3, 2]"
    intS := convertToIntSlice(s)
    if intS[0] != 5 || intS[1] != 3 || intS[2] != 2 {
    //if (intS[0] != 5) {
        t.Errorf("Expected int slice [5, 3, 2] but got %v", intS)
    }
}

// TestCompareIntSlices tests the compareIntSlices function
func TestCompareIntSlices(t *testing.T) {
    s1 := []int{2, 5, 9}
    s2 := []int{2, 5, 9}
    result := compareIntSlices(s1, s2)
    if !result {
        t.Errorf("Expected %v and %v to be the same, but comparison result was %v", s1, s1, result)
    }
}
