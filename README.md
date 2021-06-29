# CRUD_GO_SQL
CRUD operation using go, sql, orm(beego)

```kill -9 $(lsof -t -i:8080)``` // kill a process running at port 


//install golang

MAC (M1 Sillicon)

Install HomeBrew as mentioned in``` https://brew.sh/```
Install GNU Make version 4 or above using brew install make. Update the path variable in your ~/.zshrc file as suggested at the end of make installtion. It will look something like export PATH="/opt/homebrew/opt/make/libexec/gnubin:$PATH".
Install GNU Core Utils using ```brew install coreutils```
Install golang 1.16+ using ```brew install golang```


Linux

```
wget https://dl.google.com/go/go1.14.4.linux-amd64.tar.gz
sudo tar -xvf go1.12.2.linux-amd64.tar.gz
sudo mv go /usr/local
echo 'export GOROOT=/usr/local/go' >>~/.profile
echo 'export GOPATH=$HOME/go_projects' >>~/.profile
echo 'export PATH=$GOPATH/bin:$GOROOT/bin:$PATH' >>~/.profile
source ~/.profile
```
check go version

``` go version```

----------------------------------------------------------------------------------------------------------------------
Anywhere ```mkdir Api```


cd Api



```go mod init com.api.rest``` (creates go.mod shows all dependencies)



```code .```


write code with dependency



1.```go mod tidy``` (install all dependencies and creates go.sum)



2.```go mod vendor``` (shows all dependencied)


to download specific verison do ```go get``` automatically install that version and update in go mod  example -```go get github.com/spf13/cast```


now packages will also be visible like ```"../config"```  or will user ```com.api.rest/config"``` auto

for every new local or git hub package run 1 &  2 then the package will be visible


----------------------------------------------------------------------------------------------------------------------------------------------------------


install mysql-workbench from software or terminal 
``` sudo apt install mysql-workbench```


```sudo apt install mysql-server```

```sudo mysql_secure_installation```

//set up the u and p

```sudo mysql -u root -p```

//root as u name will ask for password give same password 
//this will get in sql console

```show databases;
create database local;
use local;
show tables;
```


```SELECT user,authentication_string,plugin,host FROM mysql.user;```

check for authentication socket it should not be there


```ALTER USER 'root'@'localhost' IDENTIFIED WITH mysql_native_password BY 'Current-Root-Password';```


```FLUSH PRIVILEGES;```


```update user set authentication_string=password('1234') where user='root';```

 ```flush privileges;```

go to workbench enter username passoword and database name which got created 

create a connection string using //“user:password@tcp(Hostname:Port)/dbname?charset=utf8&parseTime=True&loc=Local” logic

connection string

``` "root:root@tcp(127.0.0.1:3306)/local?charset=utf8&parseTime=True" ```

paste the connection string 
----------------------------------------------------------------------------------------------------------------------
Create tables

//for long response keep text not varchar of any size and give as string
```UNIQUE KEY `userId` (`userId`,`bookId`),```  //for constraint as 2 set of keys which means both user and book id together should be unique

```
CREATE TABLE `Books` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `Name` varchar(64) DEFAULT NULL,
  `Author` varchar(64) DEFAULT NULL,
  `Publication` varchar(64) DEFAULT NULL,
  `updatedOn` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `createdOn` datetime DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4971 DEFAULT CHARSET=utf8;
```


```SET SQL_SAFE_UPDATES = 0;```

```
CREATE TABLE `Users` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(64) DEFAULT NULL,
  `bookId`int(11) DEFAULT NULL,
  `updatedOn` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `createdOn` datetime DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `fk_User_Book_1_idx` (`bookId`),
  CONSTRAINT `fk_User_Book_1` FOREIGN KEY (`bookId`) REFERENCES `Books` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3001 DEFAULT CHARSET=utf8;
```
------------------------------------------------------------------------------------------------------------------------------------------------------
```go run main.go```


Post Payload for book
```
{
   "name": "Harry Potter",
   "about": {
       "author" :"J.K",
       "publication":"Potter Head"
   }
} 
```


POST payload for user

```
{
   "name": "Akshansh Ohm",
   "bookId":{
            "Id": 4979,
            "Name": "Harry Potter",
            "Author": "J.K",
            "Publication": "Puttar"
        }
} 
```


