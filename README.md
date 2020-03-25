# compositelogger
My complete "composite" logger (practice project).

## How to run
For a demonstration of module functionality :

```
Download "main.go"
Execute these commands :  

"go mod init "foldername"
"go run main.go"

in the terminal.
``` &nbsp;
## Database configuration
To setup the database you need to install the mysql-server, which you can look up online.

**NOTE:** Database is named "LOGGER" on mysql server, table is named "LOGS".

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