# Go: Getting Started

Go is the simple yet powerful language that is revolutionizing how web services and devops tools are created. In this course, you'll learn everything you need to get started creating your own Go applications.
![4482741](https://user-images.githubusercontent.com/14364518/116328631-d29bda00-a79f-11eb-9452-8959c6a4c133.png)

## Introduction - Hey! Ho! Let's Go!

![Intro](https://user-images.githubusercontent.com/14364518/116329687-463ee680-a7a2-11eb-98cb-4897a831cd31.png)

## Course Link 

[Pluralsight Course Link](https://www.postman.com/mrj84/workspace/go-language/collection/1830326-dd360df1-73af-46f7-b635-56cbb51ee1fa?ctx=documentation)

## Certificate

[Course Certificate](https://drive.google.com/file/d/13Tdxgr550HpYrGbGQrJ6Jm2QHLCotsf1/view)

## How to test

### Prerequisites
To follow the example in this repo, you will need:

A Go workspace set up by following [How To Install Go and Set Up a Local Programming Environment](https://www.digitalocean.com/community/tutorial_series/how-to-install-and-set-up-a-local-programming-environment-for-go).

### Building Go Binaries With go build
Using go build, you can generate an executable binary for our Go application, allowing you to distribute and deploy the program where you want.

Try this with main.go. In our Go-Getting-Started directory, run the following command:

`go build`
 
If you do not provide an argument to this command, go build will automatically compile the main.go program in your current directory. The command will include all your *.go files in the directory. It will also build all of the supporting code needed to be able to execute the binary on any computer with the same system architecture, regardless of whether that system has the .go source files, or even a Go installation.

In this case, you built your Go-Getting-Started application into an executable file that was added to your current directory. Check this by running the ls command:

`ls`
 
If you are running macOS or Linux, you will find a new executable file that has been named after the directory in which you built your program:


`controllers  `**`Go-Getting-Started`**`  go.mod  go.sum  main.go  models  README.md`

Note: On Windows, your executable will be **Go-Getting-Started.exe**.

By default go build will generate an executable for the current platform and architecture. For example, if built on a linux/386 system, the executable will be compatible with any other linux/386 system, even if Go is not installed. Go supports building for other platforms and architectures, which you can read more about in our Building Go Applications for Different Operating Systems and Architectures article.

Now, that you’ve created your executable, run it to make sure the binary has been built correctly. On macOS or Linux, run the following command:

`./Go-Getting-Started`
 
On Windows, run:

`Go-Getting-Started.exe`
 
Now you can go to the postman and test our endpoint:

[Postman requests](https://www.postman.com/mrj84/workspace/go-language/collection/1830326-dd360df1-73af-46f7-b635-56cbb51ee1fa?ctx=documentation)