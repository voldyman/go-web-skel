#Website Skeleton in Go

This is a base which can be used to build websites.


##Requirements 
* mux  from Gorilla toolkit

    $ go get github.com/gorilla/mux


## Additional notes

The files in **/static/** are served by go right now but for production this should be diabled, a webserver like nginx or apache should serve them.

