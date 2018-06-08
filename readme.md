# Embedstatic-go
## About
Embedstatic-go is a nodejs script that lets you embed static files into go source code by creating a source code file 
containing a map (string -> []byte). This lets you create a web server that compiles into a single binary.    
One or several directories can be added and it can be easily added as a build step using npm and node js. Since it is common
to use npm and webpack for building the frontend, this package will be a better fit into the existing front end 
build pipeline than corresponding implementations written in go.     
Checkout the example in the [example directory](example)

## Installation
### using npm
```
npm install git+https://github.com/alphapeter/embedstatic-go.git
```
This will install the command `embedstatic` into the node bin directory. If you don't use the -g (global flag) when
installing, the command will only be available from npm or npx.   

## Running the command
The command takes three arguments,  
`embedstatic` `source files/directories` `go-package-name` `output-file-name`

## Configuration
A better way of automating the build process is using npm and create a script that runs 'embedstatic'. This allows for
defining the configuration using a npm configuration object `embedstatic-go`. 
### package.json example

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
Then the embedding can be triggered by running `npm run embed`    
The parameters are configured by creating the `embedstatic-go` object in package.json.  
* `files`: a comma separated list of files and folders to be embedded (no white spaces)
* `package`: the package name for the output go source file
* `output`: the file name for the go source file 

