resize.go
=========

Fork of the image resizing library for Go that used to be in Gorilla
(at
http://code.google.com/p/gorilla/source/browse/lib/appengine/example/moustachio/resize/resize.go?r=3dbce6e267e9d497dffbce31220a059f02c4e99d)

but seems to have been taken out. 

What I've done:
--------------

* updated it to work with Go 1.0's image API

What I plan on doing:
---------------------

* improving the quality of the resizing by adding antialiasing or bicubic interpolation
* making the API a little nicer for some common cases. ie, more like Python's PIL.
* experimenting with ways of resizing large images with constrained memory (ie, not ever keeping an entire massive bitmap
in memory all at once).

I may end up dumping this code entirely and just implementing image
resizing by directly porting the relevant parts of Python's PIL to
Go. I haven't really decided yet.

Ultimately, I'm looking to create a library that will give me, for Go,
when I'm used to from Python in terms of resizing/cropping images for
the web. What I'm going for is an image handling foundation for Go
that will let me port my Apomixis distributed image server
(https://github.com/thraxil/apomixis) from Python/Django to Go.

License remains BSD.
