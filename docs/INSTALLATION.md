Follow the instructions for your system on the installation page at [golang.org](http://golang.org/doc/install).

Exercism supports Go 1.2 and higher.

Note about differing Go and Exercism paths: If you have Exercism installed outside of where you setup your $GOPATH to point to, and you want to keep it that way, you will need to do an extra step to ensure the Go tooling will be aware of that directory.

The way to handle cases like this is currently a topic of discussion in the Go community. One way this can be accomplished for the simple case of running Exercism code is to symlink your Exercism directory into your GOPATH. Here is a template you may follow:

`ln -s /Users/myusername/exercism/go $GOPATH`

Source:
https://groups.google.com/forum/#!msg/golang-nuts/f5ZYztyHK5I/H5K46Eete3gJ
