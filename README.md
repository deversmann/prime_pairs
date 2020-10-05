# prime_pairs

This is a program written in go to solve the Matt Parker's [Maths Puzzles](http://www.think-maths.co.uk/primepairs): [Prime Pairs](http://www.think-maths.co.uk/primepairs) puzzle from September 30, 2020. The puzzle is described as:

> Rearrange the numbers from 1-9, such that all adjacent pairs sum to a prime number.

While I came up with my initial solution by building a spreadsheet to check candidate solutions and then randomly trying some swaps till a solution worked (only took a few minutes), I decided to try to use the problem as a tool to help me learn the [Go](http://golang.org) programming language.  After a few hours, I came up with something workable, but spent most of the weekend tweaking and improving.  Time well spent!

The program take a single optional argument which is the top of the range of positive integers. There is a max set at 15 and it defaults to 9. The logic is as follows:

- Separate the list of numbers into evens and odds since no two evens or odds can appear next to each other.
- Compute all of the permutations of each of those sub lists.
- In a doubly-nested loop, interleave and test each combination of permutations, saving the valid solutions.
- If the top number was even, reverse every solution and add to the list (since the above logic will only find the solutions starting odd and ending even and every one of those has a valid mirror that won't be found).
- Sort the final list for prettier display
- Print the list with timings for both computation and sort/display.

Sample output:
```
> go run prime_pairs.go 5
There are 4 solutions for the first 5 positive integers:
[1 4 3 2 5]
[3 4 1 2 5]
[5 2 1 4 3]
[5 2 3 4 1]
End of the 4 solutions for the first 5 positive integers
Compute took 22.58µs not including the additional 54.824µs to sort and display
> 
```
