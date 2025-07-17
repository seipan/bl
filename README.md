# b(lock)l(og)
Golang block log for debug

## Usage
```go
package main

import(
  github.com/seipan/bl
)


bl.Println("Hello, World!")
bl.Println("Error %w", fmt.Errorf("an error occurred"))
```

### output
```
--------------------
| Hello, World!    |
--------------------

------------------------------
| Error %w an error occurred |
------------------------------
```
