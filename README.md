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

**NOTE:** Database is named "LOGGER" on mysql server, table is named "LOGS".

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
While on the server, you can create a user with this command :

```
"CREATE USER 'compositelogger'@'localhost' IDENTIFIED BY 'Mystrongpassword1234$';"

```
Then you need to grant the user access to our logging table, or else we will get an error :

```
"GRANT ALL PRIVILEGES ON LOGGER.LOGS TO 'compositelogger'@'localhost';"

```
Here we granted all privileges on our "LOGS" table to our user named "compositelogger".