Parses a html `io.Reader` and returns slice of `Link`s that contain the href and link text

## Usage

```
link.Parse(r) // r is *io.Reader
```

## Output of example files
```
$ go run examples/ex1/main.go 
[{Href:/other-page Text:A link to another page} {Href:/other-page Text:Bar Baz} {Href:/other-page Text:Foo Bar Baz}]
```
```
$ go run examples/ex2/main.go 
[{Href:https://www.twitter.com/joncalhoun Text:Check me out on twitter} {Href:https://github.com/gophercises Text:Gophercises is on Github!}]
```
```
$ go run examples/ex3/main.go 
[{Href:# Text:Login} {Href:/lost Text:Lost? Need help?} {Href:https://twitter.com/marcusolsson Text:@marcusolsson}]
```
```
$ go run examples/ex4/main.go 
[{Href:/dog-cat Text:dog cat}]
```
