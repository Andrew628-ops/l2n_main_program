(My Modified Code)
``` python
def isPalindrome(string): # define a function with string as input 
  string = string.lower().replace('',"") # converting string to lowercase
  first, last = 0, len(string)-1 # assign pointers
  while first < last: # loop from both end
    if string[first] == string[last]: # check if both are thesame
       first += 1 # increase first pointer
       last -= 1 # decrease last pointer
    else:
       return "this string is not a palindrome" 
    return "this string is a palindrome" 
string = input("enter string: ") # ask user for input
print(isPalindrome(string)) # print the result
```
***Reflection:***
I realized I needed to normalize the input and move the return statement outside the loop so the entire string is checked. I also learned that the two pointer approach is more space-efficient than slicing.