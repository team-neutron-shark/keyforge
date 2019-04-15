# keyforge
An API for Keyforge deck analysis as well as playing / simulating games. :shipit:

## What is this project?
This API is currently under development and is far from feature complete. That said, the aim of this project is to provide
an API for interacting with the Keyforge Vault as well as an API for playing and/or simulating games of Keyforge. The general
idea is that this API could be used to facilitate the creation of deck analysis tools, deck performance simulations, or
fully-fledged Keyforge game servers.

## Why on Earth did you write this in Go?
I decided on Go for a few reasons, but the biggest reasons were native system binaries and cross-platform interoperability.
Since Go produces platform-specific static binaries as program output we end up with fast, efficient code that is supported
by the following operating systems:

* Windows
* Linux (kernel version 2.6.23 or later)
* OSX
* Most popular flavors of BSD

Additionally the following CPU architectures are supported:

* AMD64
* i386
* ARM
* ARM64
* PPC64
* PPC64LE
* MIPS64
* MIPS64LE
* S390X

So in a nutshell Go produces fast programs that are supported on just about any combination of operating system and processor 
architecture.

## How do I use this code?
The easiest way to get started with this code is to clone it directly to your GOPATH. For Linux/OSX systems this is located in ~/go/src and for Windows systems this path is typically C:\Users\<username>\go\src. Once you've done this it's ready to be imported in your Go projects.

As with many personal code projects this project is currently devoid of documentation. I'm not a huge fan of sentiment of code being the documentation, but that's where we're currently at. Much of what's there is currently very straightforward, but if you have a question regarding how to use this API feel free to shoot me a line.
