# ecdchk

**Project:**

Eric Claus' Demo Check "ecdchk" (current working title), licensed GPL3 see license file

**Description:**

Quick Note: ecdchk -lump *filename* for quick use

Recording demos in the past I often used XDRE to check to make sure difficulty and such had been set correctly after I recorded a demo after a mistaken submission where I had not set the difficulty properly so as not to waste time with a bad submission. Later on I decided it would be cool to write a little command line tool that could be used to check difficulty, respawn and other flags quickly without having to load up something like XDRE just to check a demo header (this is before I learned about Kraflab's lmp more on that later)

At first I went through this with Python as I had messed around with it before but after finishing my prototype I wanted something compiled for easier distribution where Python required the interpreter. After messing around with C which felt too heavy handed and looking at a couple of other languages I ultimately settled on Go which felt comfortable as a language for me and met my wants and needs.

I later found out about Kraflab's lmp Ruby script which did what I was looking for and it also possesses analysis functions, but I still wanted to continue on with my own project as it is mostly for fun and simplicity, and I find the Ruby code less fun for me to read and try and figure out. I did use lmp to help verify some output along with digging in a hex editor to compare bytes, so with it being some source of inspiration I recommend checking out Kraflab's work because it helps to lay out demo headers in a nice way and provides analysis functions. You can find lmp here https://github.com/kraflab/lmp.

I hope in the future to maybe follow in Kraf's footsteps on deeper info and maybe analysis functions, or go down a different route and maybe even include some ways to manipulate the demo file by changing the bytes for fun or TAS purposes on the command line depending on where my whims take me.

**Install and Build Instructions:**

The program is small and straightforward you can grab the binary release for the proper OS (any recently modern version of Go should work), or to build you will need to install the go compiler on your machine for whatever appropriate operating system see , and then simply do the following in the main file "go build main.go vanillademo.go boomdemo.go mbfdemo.go mbf21demo.go sliceread.go" from your command line in the main file and it will build a binary for your system (I intend to update or improve this later but still learning a few details, but this will produce a binary).

Once you have the binary you can place it in your OS PATH somewhere or just in a folder you keep your demos in, use it from you command line from there

**How to Use:**

Just run the binary from the directory you put it in, or from anywhere if its in your PATH, you will need to do "ecdchk -lump *filename*" it will read the lump and give you some basic info about it.

**Credits and Inspirations:**

Author: Eric Claus (not real name)

<p>Help, Support, Inspiration: Ludi https://doomwiki.org/wiki/Ludi.  </br>
Inspiration: Kraflab https://github.com/kraflab/lmp, check out DSDA Doom as well here https://github.com/kraflab/dsda-doom  </br>
Moral Support: Alaux, check out his Nugget Doom port based on Woof https://github.com/MrAlaux/Nugget-Doom  </br>
Moral Support: Mikolah https://doomwiki.org/wiki/Mikolah  </br>
Shoutout: my friends in the Afternoon Crew, love you guys!  </br>
Help on Slice Reading: this post on Stack Overflow was very informative and the code helpful to customize to my use...  </br>
https://stackoverflow.com/questions/70115047/convert-binary-file-to-array-of-bytes more info can be found in the comments in sliceread.go in main  </br>
The Go Documentation was very helpful, and left a nice starting point to find other articles for, especially when I looked into the switch statement syntax, good article here that clarifies on the Go documentation a bit for them https://www.eternaldev.com/blog/go-switch-6-ways-of-using-switch-in-go/  </p>
