(Pseudocode)
```
Define function isPalindrome with input string

    If string is equal to reverse of string
        Return "isPalindrome"
    Else
        Return "is not palindrome"

ask user to enter a string
store input in string

call isPalindrome with string
print the result

```


My Code(Before Asking)
```python
def isPalindrome(string): # define a function with string as input
    if (string == string[::-1]): # check if string equal the reversed
        return "isPalindrome" 
    else:
        return "is not palindrome" 
string = input("enter string: ") # ask user for input 
print(isPalindrome(string)) # print the result
```
   ***Text Result:*** input "racecar" it print out isPalindrome, input "hello" it print out is not palindrome, input "A man a plan a canal Panama" it print out is not palindrome

   ***Comment:*** I used slicing to reverse the string because it’s simple and easy to remember and understand from a tutorial i watched, but for very long strings a two pointer approach would be more efficient.

   # Reflection:
   ***What did i learn from solving it before asking AI?***

   ***Ans:*** I learned how to reverse a string using slicing, compare it to the original, and that trying first helps me understand the problem better.

   ***How is my understanding different now?***
   
   ***Ans:*** I now understand why slicing works, its limitations, and more efficient alternatives like the two pointer approach.
    
   ***Could i now write similar functions (e.g., reverse a string) without help?***

   ***Ans:*** yes, I can write simple functions like reversing a string without help. 
