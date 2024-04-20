## The elements
There are 4 elements that make up backtracking:

1. The visit
  - This is an act of visiting a node in a tree
  - This function 'records' the visit by saving it into memory
  - In my solution, that is the `placeQueen()` function

2. The abandon
  - This is an act of abandoning a node in a tree
  - This function deletes the record of the visit by erasing it from memory
  - In my solution, that is the `removeQueen()` function

3. The dive
  - This is an act of going deeper down the tree of nodes
  - The dive is the backtrack function itself
  - In my solution, that is the `backtrackNQueen()` function

4. The exit
  - This is an act of deciding to stop diving because there are no solutions ahead

5. The arbiter
  - The arbiter decides whether it is worth to visit the node

## The pattern
```
dive() {
    arbiter() {
        visit()

        if !exit() {
            dive()
        }

        abandon()
    }
}
```


