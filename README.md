# About

A Go slice expression quiz game. Go slice expressions are explained in the docs here:

https://go.dev/ref/spec#Slice_expressions

The quiz game helps get familiar with Go slice expresison syntax like `[2:3]`, `[:4]`, or `[1:]`

## Slice Expressions; Inclusive vs Exclusive

Thinking of slice expressions in terms of `[inclusive:exclusive]` is helpful. ie, the value of the index number before the `:` is included in the resultant sub-slice, the value of the index number after the `:` is excluded from the resultant sub-slice.

# Example Output

The game generates a random slice and a random slice expression, then asks you to guess it:

```
[2 0 5 2 8 5 0 0 8 9]
What is the result of [2:9] (?=skip [and see answer], q=quit):
5285008
Nice! You guessed [5 2 8 5 0 0 8] correctly!
```

Answers can just be a series of numbers typed in, representing whatever part of the slice (ie, whatever sub-slice) is represented by the slice expression.

If you guess wrong it will prompt you to try again:

```
[8 0 9 3]
What is the result of [1:3] (?=skip [and see answer], q=quit):
093
Hmm, [0 9 3] wasn't right. Try again:
[8 0 9 3]
What is the result of [1:3] (?=skip [and see answer], q=quit):
09
Nice! You guessed [0 9] correctly!
```

Entering `?` as the answer will skip the current question, display the answer, then generate a new question:

```
[8 0 6 3 7 9 8 6 4 9]
What is the result of [5:9] (?=skip [and see answer], q=quit):
?
Answer was [9 8 6 4]
```

Quit with `q`; final count of correct, incorect and skipped answers is displayed:

```
[2 4 2 0 7 6 8 1]
What is the result of [3:5] (?=skip [and see answer], q=quit):
q
Final score -- correct 2, incorrect 1, skipped 1
Quitting.
```
