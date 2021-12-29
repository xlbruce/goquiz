# GoQuiz
Simple program to read a csv file containing 2 fields: a question and a answer.

## Usage
```
Usage of ./goquiz:
  -filename string
        path to CSV file containing the quiz (default "problems.csv")
  -timeout int
        time in seconds to wait for every question (default 30)
```

## Example
`goquiz -filename /home/user/quiz.csv -timeout 15`

Above command will read the quiz file and wait for 15 seconds for every question.
