####Assumptions:

* The database is stored in the form of the json file that can be updated (__operations.json__)
* The local API runs on localhost with port number 8000.
* Every transaction has unique ID (Название), which is not necessarily a number.


####The code run
The code can be tested through POSTMAN or using curl bash command. Several options were given below:  

```bash
   curl -i -X GET http://localhost:8000/transactions
```
* Above command provides the whole db in the form of the json 

```bash
   curl -i -X GET http://localhost:8000/transactions/{id}
```
* This is another command that provides a single transaction within db through __{id}__ . 
***
As an example, you can check the first transaction in given db using:
```bash
   curl -i -X GET http://localhost:8000/transactions/1
```
***

* The command below can insert new transaction by providing ___'{new transaction data}'___ in the form of json data. 
```bash
   curl -i -X POST -d '{new transaction data}' http://localhost:8000/transactions

```
***
As an example, you can add new transaction in given db using:
```bash
    curl -i -X PUT -d '{"id":"15","price":"3256","type":"purchase","comment":"The chair was purchased","category":"Furniture","date":{"year":"2022","month":"Jul","day":"23"}}' http://localhost:8000/transactions

```
Furthermore, check the ___operations.json___ db to check whether new transaction with _id:15_ was added.  
***

* This is delete command can be used to delete single transaction with particular ___{id}___ .
```bash
   curl -i -X DELETE  http://localhost:8000/transactions/{id}

```

***
As an example you can delete second transaction in db as:
```bash
   curl -i -X DELETE  http://localhost:8000/transactions/2

```
***

* Lastly you can update the transaction with specific ___{id}___ by a below cmd: 
```bash
   curl -i -X PUT -d 'updated transaction data' http://localhost:8000/transactions/{id}
```
> It should be noted that ___{id}___ is not going to change in update even if you try in ___'updated transaction data'___ .


***
As an instance, you can update third transaction in db as:
```bash
   curl -i -X PUT -d '{"id":"15","price":"45000","type":"income","comment":"The monthly wage","category":"Wage","date":{"year":"2022","month":"Feb","day":"26"}}' http://localhost:8000/transactions/3

```
***