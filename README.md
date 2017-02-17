# README #

An implementation of the VectorSpace model in Golang. Pass in two strings and get back a number indicating how similar they are.

Usage like so,

```
var concordance1 = Concordance("Go has a lightweight test framework composed of the go test command and the testing package.")
var concordance2 = Concordance("Package testing provides support for automated testing of Go packages. It is intended to be used in concert with the go test command, which automates execution of any function of the form.")

// value of got will be 0.48211825989991874   
got := Relation(concordance1, concordance2)
```


See tests for other examples.