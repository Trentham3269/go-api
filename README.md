# go-api
Simple api with golang and postgres db.

### Database notes
* Table structure is as follows:
```
CREATE TABLE honourboard (  
  name VARCHAR (50),  
  board VARCHAR (50),  
  championship VARCHAR (50),  
  year INTEGER,  
  decade CHAR (9),  
  place VARCHAR (6),  
  type VARCHAR (10),  
  location VARCHAR (20),  
  country VARCHAR (3)  
);
```
* Copied from csv (no headers) into honourboard table:
```
COPY honourboard FROM '/path/to/data.csv' WITH (FORMAT csv);
```
### Compile
* ```$ cd $GOPATH```
* ```$ go install honourboard/go-api```

### Run 
* ```$ cd src/honourboard/go-api```
* ```$ export DBHOST=localhost```
* ```$ export DBPORT=5432```
* ```$ export DBUSER=***********```
* ```$ export DBPASS=*********```
* ```$ export DBNAME=***********```
* ```$ go run app.go```
* ```$ curl localhost:8000/api/```
