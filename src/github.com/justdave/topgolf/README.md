# Topgolf Code Challenge: Sorting a list of numbers

This challenge is designed primarily to gauge your coding abilities, as well as knowledge of algorithms. Given a string containing a comma-delimited list of numbers (integers), your task is to sort the list in ascending order and output this data to a file with *one number per line*. The caveat? ***You are not allowed to use built-in sort methods! (i.e. Ruby's ".sort()" method and PHP's sort(), asort(), ksort() and so forth are out of commission)***

Example:

If given:

    input = "5,2,3,1,4"

Your output should be:

    1
    2
    3
    4
    5

Your input will consist of a string of 200 numbers between 1-10000. There are no duplicate numbers. The output we are expecting is shown in the output/expected.txt file. You may write your solution in PHP, Ruby, Python, or Javascript (Node.js). For each of these languages, there is some starter code in the src/ directory containing code to write to your output file.

It may not be possible to completely finish the task in time. **We will NOT penalize you for that!** We are not looking to evaluate based on the final product--we are looking to see how you solve problems, lookup issues, ask questions, and write code!

algorithm quicksort(A, lo, hi) is
    if lo < hi then
        p := partition(A, lo, hi)
        quicksort(A, lo, p - 1 )
        quicksort(A, p + 1, hi)

algorithm partition(A, lo, hi) is
    pivot := A[hi]
    i := lo - 1    
    for j := lo to hi - 1 do
        if A[j] < pivot then
            i := i + 1
            swap A[i] with A[j]
    swap A[i + 1] with A[hi]
    return i + 1

######################################

algorithm quicksort(A, lo, hi) is
    if lo < hi then
        p := partition(A, lo, hi)
        quicksort(A, lo, p)
        quicksort(A, p + 1, hi)

algorithm partition(A, lo, hi) is
    pivot := A[lo]
    i := lo - 1
    j := hi + 1
    loop forever
        do
            i := i + 1
        while A[i] < pivot

        do
            j := j - 1
        while A[j] > pivot

        if i >= j then
            return j

        swap A[i] with A[j]