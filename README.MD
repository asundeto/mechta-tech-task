# Techincal task for Мечта.Kz

## Task information

### Write a program to calculate the sum of numbers in a file

Given a json file as an array of objects as shown below:
[
{
"a": 1,
"b": 3
},
{
"a": 5,
"b": -9
},
{
"a": -2,
"b": 4
}
...
]

Number of objects = 1,000,000
number values ​​in the range [-10,10]

The program must read the file and calculate the sum of all the numbers. To do this, it must:
⁃ Parallelize the calculation by goroutines
⁃ The number of goroutines for parallel processing must be obtained when running the program via an argument
⁃ Output the overall result to the console
⁃ Publish the solution on github.com and provide access

### To run code use: go run . "json file name" | Example go run . test.json