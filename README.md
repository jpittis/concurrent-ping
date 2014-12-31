Concurrent Get Requests in Go
===============
Concurrently send get requests to urls from standard input. Sends status to standard output.

Use with this command on osx to print the results to stdout. ```./rout < ping.txt```

Other operating systems have to build the go source first. ```go build rout.go```

This project is based off of [John Graham's talk](http://youtu.be/woCg2zaIVzQ?list=PLMW8Xq7bXrG58Qk-9QSy2HRh2WVeIrs7e) at dotGo 2014.
