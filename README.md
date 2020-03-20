# compositelogger
My complete "composite" logger.

"Logs" helper for private project (composite logging) can be found in its own repo (don't know if this is correct practice) as well as here.

For a demonstration of module functionality, it's just necessary to download the driver code, then do

go mod init "foldername"
go run main.go

it will import my "logs" helper (and everything else) outside of $GOPATH.

NOTE: Database is named "LOGGER" on mysql server, table is named "LOGS". (root@localhost) on my end.
