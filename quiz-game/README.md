# Quiz Game

This is a small program that read a quiz provided via a CSV file (more details below) and will then give the quiz to a user keeping track of how
many questions they get right and how many they get incorrect. Regardless of whether the answer is correct or wrong the next question asked immediately afterwards.

The CSV file default to **problems.csv** (example shown below), but the user should be able to customize the filename via a flag ` -csv `.

The CSV file will be in a format like below, where the first column is a question and the second column in the same row is the answer to that question.

The Quiz time limit defaul to **30 seconds**, but the user can customize the limit via a flag ` -limit `.

```CSV
5+5,10
5*5,25
10*10,100
50-10,40
45-33,12
```
