### go-api
Simple api with golang and postgres db.

#### Database notes
Table structure is as follows:
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
