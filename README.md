# Tiny- Url

## Requirements

- Golang
- Mysql

## Setup

```shell
git clone git@github.com:anujkumarthakur/coinproject.git
```

```shell
cd coinproject
```

```shell
go get
```

```shell
database configure .env 
```

```shell
create database test;
```

```
CREATE TABLE tinyurl (
	id int NOT NULL AUTO_INCREMENT,
	long_url varchar(1000) NOT NULL,
	short_url varchar(500) NOT NULL,
	url_id bigint(20) NOT NULL,
    PRIMARY KEY (id)
);
``` 

```shell
go build -o app
```

```shell
./app
```
