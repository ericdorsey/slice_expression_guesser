# About

A Go slice expression quiz game. Go slice expressions are explained in the docs here:

https://go.dev/ref/spec#Slice_expressions

The quiz game helps get familiar with slice expresison syntax like [2:3], [:4], or [1:]

## Example Output

The game generates a random slice and a random slice expression, then asks you to guess it:

```
[9 4 7 8 0 4 1]
What is the result of [3:5] (Use '?' to skip [and see answer]):
80
Nice! You guessed [8 0] correctly!
```

Answers can just be a series of numbers typed in, representing whatever part of the slice is represented by the slice expression.

If you guess wrong it will prompt you to try again:

```
Starting again!

[0 1 4 7 9 7 4 2 1 0]
What is the result of [3:4] (Use '?' to skip [and see answer]):
79
Hmm, [7 9] wasn't right. Try again:
[0 1 4 7 9 7 4 2 1 0]
```

Entering `?` as the answer will skip the current question, display the answer, then generate a new question.

```
What is the result of [3:4] (Use '?' to skip [and see answer]):
?
Answer was [7]
```

Quit with `CTRL + C`:

```
Starting again!

[1 1 4]
What is the result of [1:2] (Use '?' to skip [and see answer]):
^Csignal: interrupt
```
