# About

In Go an application is organized in packages. A package is a collection of source files located in the same folder. All source files in a folder must have the same package name at the top of the file. The package is preferred to be the same as the folder it is located in.

```go
package greeting
```

Standard library packages are imported using their name. Other packages are imported using the url of the package's location. A package can be given an alias when importing. This can be used to avoid package name conflicts:

```go
package greeting

import (
	"errors"

	errs "github.com/pkg/errors"
)
```

An imported package is then addressed with the package name or alias:

```go
// using the internal errors package
errors.New("Connection not established")

// using the errors package located at github.com/pkg/errors by alias
errs.New("Connection not established")
```
