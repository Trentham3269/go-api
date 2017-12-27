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
### Port forwarding
To be able to view in Chrome on my computer:
* ```$ ssh -v -L 8000:localhost:8000 user@remoteserver``` on the local machine
* ```$ go-api``` on the remote server
* View in the browser on the local machine at ```localhost:8000/api/```
