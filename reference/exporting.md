# Exported Identifiers

Go does not have `public` and `private` keywords. Instead it exports identifiers that start with a capital letter.

> An identifier is exported if both:
> - the first character of the identifier's name is a Unicode upper case letter (Unicode class "Lu"); and
> - the identifier is declared in the package block or it is a field name or method name.
>
> All other identifiers are not exported.
[The Go Programming Language Specification](https://golang.org/ref/spec#Exported_identifiers)
