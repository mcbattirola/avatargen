# Avatargen

<div align="center">
	<img src="./assets/golang-is-awesome.svg" alt="avatar" />
</div>

This is a simple project to generate unique 300x300 avatars for any input string, havily inspired by [loweisz's generate-avatar](https://github.com/loweisz/generate-avatar/tree/6b48f98117bfa0cbbf38954c98e16b09b0906df8).

### How to use

Get the package:

```
go get github.com/mcbattirola/avatargen
```

Call the method Generate with a string as argument:

```go
import "github.com/mcbattirola/avatargen"


svg := avatargen.Generate("mystring")
fmt.Println(svg) 

/* output:
<svg width="300" height="300" viewBox="0 0 300 300" xmlns="http://www.w3.org/2000/svg"><rect id="bg" width="300" height="300" fill="rgb(22,147,25)" /><path d="m 150 597 Q 282 18 -297 150 Q 282 282 150 -297 Q 18 282 597 150 Q 18 18 150 597 z" fill="rgb(232,249,208)" /><path d="m 150 368 Q 203 97 -68 150 Q 203 203 150 -68 Q 97 203 368 150 Q 97 97 150 368 z" fill="rgb(208,249,232)" /><path d="m 150 265 Q 135 165 35 150 Q 135 135 150 35 Q 165 135 265 150 Q 165 165 150 265 z" fill="rgb(44,50,50)" /></svg> 
*/
```

See the example in ```cmd/avatar-gen/main``` to save it as a .svg file.

If you just want to generate your avatar, without using the package, you can download the source code and run ```main``` with a string argument: 

```
go run cmd/avatar/main.go mystring
```

it will create a file ```mystring.svg```

### Avatars

The SVGs avatars are created based on a MD5 hash from the input string, so it should aways be unique. 

There's no randomness, so each input produce aways the same output. 
This means you don't have to store it and can run everytime you need.

The avatars have currently a fixed size of 300x300.
