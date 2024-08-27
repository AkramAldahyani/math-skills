# math-skills
This code is going to act as a statistical snalysis program.
The program will read a file you stored the data set in it and calculate the following:
- Average
- Median
- Variance
- Standard Deviation

Usage:
- First: store the data set in a file and use between them a space or a new line. The code is able to reade the floating point as well, but be care full to use a dot (.) not comma (,) in the number.
    examples: 
    12 34 56 67 78 89 : true.
    12/23/34/45/56/67 : false.
    12.2 23.5 45.7 67.8 : true.
    766,4 4545,5 45,6 : false

- second: To run the code you will run main.go and add the file name next to it. 
    example: 
    go run  main.go file.txt
