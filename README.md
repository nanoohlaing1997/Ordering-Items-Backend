# go-order-items

Simple testing project with golang

### Requirement

- go 1.19
- mysql 5.7

### Setup guide for developers

- git clone
- copy .env.example to .env
- run go mod tidy
- run go mod vendor
- create database `order`
- copy and paste table schema from schema.sql file to create `order` database's tables
- go to appliction folder path and run **go run main.go**

### API Doc

- POST Sign_up of User

  `request(json)` {
	"username",
	"email",
	"password",
	"address"
  }

  `response(json)` {
	"token" : "gsdgsdewcf",
	"refresh_token" : "tw523532523"
  }

- GET Sign_in of User

  `request(json)` {
	"email",
	"password"
  }

  `response(json)` {
	"token" : "gsdgsdewcf",
	"refresh_token" : "tw523532523"
  }

- GET token refresh

  `request(json)` {
	"email",
	"refresh_token"
  }

  `response(json)` {
	"token" : "sdgfewtwdsf",
	"refresh_token" : "sdgdsgssdg"
  }

- GET list of categories (all) (for customer)
- GET list of Items {category_id} (for customer)
- POST create Item_categories {"name"} (for admin)
- POST create Items {"name", "price", "quantity", "category_id"} (for admin)
- POST register Customers {"email", "password", "address"} (for customer)
- POST login Customers {"email", "password"} (for customer)
- POST customer order items {"customer_id", "item_id", "quantity"} (for customer)
- GET list of orders {"customer_id", "invoice_id"} (for customer) [to check delivery_status]
- GET list of orders (all) (for both customer and admin)
- POST change delivery status in orders ("invoice_id", "delivery_status") (for admin)

### Database Table Diagram

![order Database Diagram](ordering_item.png)
