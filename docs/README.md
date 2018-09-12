# Documentation

##### Endpoint
List of endpoint

| **Endpoints** | **Method** | **Description** |
|-----------------|:------------:|-------------------|
|`/sanitycheck/submit`| `GET`   | Web view to testing input bill |
|`/sanitycheck/submit`| `POST`  | Web view to testing save bill. It will redirect to `/sanitycheck/submit` if fail, and `/sanitycheck/submit` if success|
|`/sanitycheck/list`   | `GET`       | Web view to display all bills |
|`/bill/list`    | `GET`        | To get all bills |
|`/bill`         | `POST`        | To save bill |

##### GET `/sanitycheck/submit`
Serve as web view to testing input bill

##### POST `/sanitycheck/submit`
Serve as endpoint to testing save bill. It is a little bit different than `POST /bill`. If error happen when input bill will redirect to  `/sanitycheck/submit` and show the error. If success will redirect to `/sanitycheck/list`

##### GET `/sanitycheck/list`
Serve as web view to list all bills that have been recorded in database.

##### GET `/bill/list`
Serve as endpoint to return all bills that have been recorded in database as JSON type.
Example responses
- Success, but empty record
```JSON
{
  "status": 200,
  "result": {
    "bills": null,
    "total_amount": "0",
    "total_tax": "0",
    "total_original_price": "0"
  }
}
```
- Error
```JSON
{
  "status": 500,
  "result": "sql: expected 7 destination arguments in Scan, not 6"
}
```
- Succes with data
```JSON
{
  "status": 200,
  "result": {
    "bills": [
      {
        "bill_id": 1,
        "product_name": "Kitkat",
        "tax": {
          "tax_id": 1,
          "type": "Food"
        },
        "original_price": 10000,
        "tax_amount": 1000,
        "total_amount": 11000
      },
      {
        "bill_id": 2,
        "product_name": "Cigar Nat Sherman Sterling Series",
        "tax": {
          "tax_id": 2,
          "type": "Tobacco"
        },
        "original_price": 550000,
        "tax_amount": 11010,
        "total_amount": 561010
      },
      {
        "bill_id": 3,
        "product_name": "Venom 2018",
        "tax": {
          "tax_id": 3,
          "type": "Entertainment"
        },
        "original_price": 60000,
        "tax_amount": 599,
        "total_amount": 60599
      }
    ],
    "total_amount": "632609",
    "total_tax": "12609",
    "total_original_price": "620000"
  }
}

```
##### POST `/bill`
Serve as endpoint to save bill to database. Only accept one bill per call

Accept form data as the parameter

| Body       | Value    | Desc |
|------------|----------|------|
| `product_name `    | `string` | Must not empty, at least 1 char |
| `tax_code` | `string`    | Must not empty, currently only accept value 1, 2, 3|
| `amount`   | `string`    | Must not empty, and price must bigger than 0 |

Example usage
- From command line
```sh
$ curl -d "product_name=Kitkat&tax_code=1&amount=1000" -X POST http://localhost:9001/bill
```

- From web GUI
Please visit `/sanitycheck/submit`

Example responses
- Error
```JSON
{
    "status":500,
    "result":"product name should not empty"
}
```
- Success
```JSON
{
    "status":200,
    "result":"Success add bill"
}
```

##### Database

See schema [image](https://github.com/jonathanhaposan/taxcalc/docs/schemadb.png)

this application use 2 table : `tax_code` and `bill`, the relationship between these two table are `one to many`
