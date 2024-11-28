module backend/Web

go 1.23.0

require (
	backend/Model/user v0.0.0
	github.com/gorilla/handlers v1.5.2
	github.com/gorilla/mux v1.8.1
	gorm.io/driver/mysql v1.5.7
	gorm.io/gorm v1.25.12
)

require github.com/felixge/httpsnoop v1.0.3 // indirect

require (
	filippo.io/edwards25519 v1.1.0 // indirect
	github.com/dgrijalva/jwt-go v3.2.0+incompatible // indirect
	github.com/go-sql-driver/mysql v1.8.1 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	golang.org/x/text v0.18.0 // indirect

)

replace backend/Model/user => ../Model/user
