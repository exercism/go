# Introduction

In Go an application is organized in packages.
A package is a collection of source files located in the same folder.
All source files in a folder must have the same package name at the top of the file.
The package is preferred to be the same as the folder it is located in.

```go
package greeting
```

Go provides an extensive standard library of packages which you can use in your program using the `import` keyword.
Standard library packages are imported using their name.

```go
package greeting

import "fmt"
```

