module gPRC

require (
	github.com/go-sql-driver/mysql v1.4.0
	github.com/golang/glog v0.0.0-20160126235308-23def4e6c14b // indirect
	github.com/golang/protobuf v1.2.0
	golang.org/x/net v0.0.0-20180906233101-161cd47e91fd
	golang.org/x/sync v0.0.0-20180314180146-1d60e4601c6f // indirect
	golang.org/x/sys v0.0.0-20180907202204-917fdcba135d // indirect
	golang.org/x/text v0.3.0 // indirect
	google.golang.org/genproto v0.0.0-20180831171423-11092d34479b // indirect
	google.golang.org/grpc v1.14.0
	gopkg.in/DATA-DOG/go-sqlmock.v1 v1.3.0
)


http://10.5.245.87:8080/enigma/subscription?si=9419111261&productCode=PRIME
{
  "si": "9419111261",
  "productCode": "PRIME",
  "chargingPrice": 0,
  "subProductCode": "Annual",
  "chargeThroughDate": "2020-01-21",
  "subscriptionStatus": "UNSUBSCRIBED",
  "subscriptionUpdateDate": "2019-04-12 21:16:30.685",
  "unsubscriptionReason": "CustomerRequestedCancel",
  "createDate": "2019-01-22 20:46:29.228",
  "isMarkedForCancel": true,
  "chargingCycle": "YEARLY",
  "productPrice": 999,
  "lob": "Mobility",
  "periodStartDate": "2019-01-22",
  "markedForCancel": true
}



http://10.5.245.87:8080/enigma/subscription/v2?si=9419111261&productCode=PRIME
{
    "si": "9419111261",
    "productCode": "PRIME",
    "chargingPrice": 0,
    "subProductCode": "Annual",
    "chargeThroughDate": "2020-01-21",
    "subscriptionStatus": "SUBSCRIBED",
    "subscriptionUpdateDate": "2019-04-12 21:16:30.685",
    "unsubscriptionReason": "CustomerRequestedCancel",
    "createDate": "2019-01-22 20:46:29.228",
    "isMarkedForCancel": true,
    "lob": "Mobility",
    "periodStartDate": "2019-01-22",
    "markedForCancel": true
  }
]