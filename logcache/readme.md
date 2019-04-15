
Write a method that will retrieve time series data from a server, but will cache the results, and
only fetch new data if the data is not in the cache.

Any time series is specified by two parameters, start and end.

You will be evaluated on correctness (think of edge cases), as well as performance of the
implementation.

Please include runtime complexity (Big O-notation) of the your solution

// GIVEN
const requestData = async (timeStart /* Number */, timeEnd /* Number */) => { /* returns a
promise of an Array of Data Objects */ }
let pieceOfData: Data = {
start: Unix timestamp (Number)
end: Unix timestamp (Number)
data: { ... } (Object)
}
const handleRequest = async (from /* Number */, to /* Number */) => {
// YOUR ANSWER HERE
// returns a Promise of an Array of Data Objects
}
Implement your solution using one of the following programming languages: Rust, Go, or
JavaScript.