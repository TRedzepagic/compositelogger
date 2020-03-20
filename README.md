# compositelogger
My complete "composite" logger.

The "Logs" helper can be found in its own repo (don't know if this is correct practice) as well as here.

If errors occurr, please go get the necessary packages.

For a demonstration of module functionality, it's just necessary to download the driver code (main.go), then change

logs "github.com/TRedzepagic/compositelogger/logs" (I did my coding inside $GOPATH) to
"github.com/TRedzepagic/logs"

then do :

"go mod init "foldername"
"go run main.go"

in the terminal.

It will import my "logs" helper (and everything else) outside of $GOPATH.

NOTE: Database is named "LOGGER" on mysql server, table is named "LOGS". (root@localhost) on my end.
