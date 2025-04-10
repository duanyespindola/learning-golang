# what i discovered

- I cannot have two packages in the same folder
- I can have a second package with the same name with "_test" in the end
- If I put the same package name in different folders:
    ```
    /pack01
       - p1.go (pack01)
    /pack02
       - p1.go (pack01)
    ```
    - I could only use the two packages importing with custom names as:
    ```
    import (
        p1 "me.com/tp/pack01"
        p2 "me.com/tp/pack02"
    )
    ```
- How to files under the same package works together ??
   - Its was like to have one single file. Each file can see everthing the other file have declared.
      - Even variables ??
         - Yeah! Even variables!!!!