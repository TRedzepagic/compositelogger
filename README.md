# compositelogger
My complete "composite" logger (practice project).

## How to run
For a demonstration of module functionality, it's just necessary to download the driver code (main.go), then change
```
logs "github.com/TRedzepagic/compositelogger/logs" (I did my coding inside $GOPATH) to
"github.com/TRedzepagic/logs"
```
then execute the command :  
```
"go mod init "foldername"
"go run main.go"
```  
in the terminal.

It will import my "logs" helper (and everything else), thus working outside of $GOPATH.

&nbsp;
&nbsp;

**NOTE:** Database is named "LOGGER" on mysql server, table is named "LOGS". (root@localhost) on my end.

To setup the database you need to install the mysql-server, which you can look up online.

To get the exact same table as me, inside the mysql shell, type these commands :
```
CREATE DATABASE LOGGER;
USE LOGGER;
CREATE TABLE LOGS
(
    id int NOT NULL AUTO_INCREMENT,
    PREFIX varchar(255) NOT NULL,
    DATE varchar(255) NOT NULL,
    TIME varchar(255) NOT NULL,
    TEXT varchar(255) NOT NULL,
    PRIMARY KEY (id)
);
```
If you have issues accessing the DB via Go code, you can try:
```
"ALTER USER 'root'@'localhost' IDENTIFIED WITH mysql_native_password BY 'password';"
inside the mysql shell.
```
(Sets root password to "password", while this obviously isn't desired/smart behavior, it's done to make the program work, you can use other passwords and change it in code accordingly).