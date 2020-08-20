## Template Practice
This folder shows me using Golang template features to read and write data to files. The result will be shown in index.html 
### libraries  Used:
* text/template
* log
* fmt
* html/template
* os

## Things I always forget

#### The & Operator
**&** goes in front of a variable when you want to get that variable's memory address.

#### The \* Operator
 \* goes in front of a **Varible** that holds a memory address and resolves it (it is therefore the counterpart to the & operator). It goes and gets the thing that the pointer was pointing at.

\* in front of a **type**
When \* is put in front of a type, e.g. *string, it becomes part of the type declaration, so you can say "this variable holds a pointer to a string".