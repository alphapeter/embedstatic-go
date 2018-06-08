
# Example project

This shows how to use the static embedded resources in a web server that serves the files. Feel free
to copy the embededHandler.go file and include it in your project.   
Download/clone the repository and try for your self

## Dependencies 
[Nodejs](https://nodejs.org) and [golang](https://golang.org/)

## Project setup

### Configuration
```
{
  "name": "example-buildscript",
  "version": "0.0.1",
  "dependencies": {
    "embedstatic-go": "git+https://github.com/alphapeter/embedstatic-go.git"
  },
  "embedstatic-go": {
    "files": "index.html,static",
    "package": "main",
    "output": "data.go"
  },
  "scripts": {
    "embed": "embedstatic"
  }
}
```
The embedding can be triggered by running `npm run embed`    
The parameters are configured by creating the `embedstatic-go` object in package.json.  
* `files`: a comma separated list of files and folders to be embedded (no white spaces)
* `package`: the package name for the output go source file
* `output`: the file name for the go source file 

All files in static and index.html will be added to the source file data.go and accessible as a
public property called Data

### http handler
The embeddedHandler.go is an example of how to implement a handler that serves all files using corresponding
content type (using the file extension). If no file is found, for example if the file is a directory, the handler will 
try to serve index.html from that directory.   

_parameters_    

func CreateHandler(data map[string][]byte, defaultResource string) EmbeddedHandler

`data`: The map created by the script
`defaultResource`: The resource to be served when accessing '/' (in most cases index.html)
 