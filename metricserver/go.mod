module metricserver

go 1.12

replace (
	github.com/zkfy/archiver => github.com/mholt/archiver v2.1.0+incompatible
	golang.org/x/crypto => github.com/golang/crypto v0.0.0-20180910181607-0e37d006457b
	golang.org/x/lint => github.com/golang/lint v0.0.0-20190301231843-5614ed5bae6f
	golang.org/x/net => github.com/golang/net v0.0.0-20180911220305-26e67e76b6c3
	golang.org/x/sync => github.com/golang/sync v0.0.0-20180314180146-1d60e4601c6f
	golang.org/x/sys => github.com/golang/sys v0.0.0-20180909124046-d0be0721c37e
	golang.org/x/text => github.com/golang/text v0.3.0
	golang.org/x/time => github.com/golang/time v0.0.0-20180820150726-fbb02b2291d28
	google.golang.org/appengine => github.com/golang/appengine v1.1.0
	google.golang.org/genproto => github.com/ilisin/genproto v0.0.0-20181026194446-8b5d7a19e2d9
	google.golang.org/grpc => github.com/grpc/grpc-go v1.16.0
)

require (
	github.com/fortytw2/leaktest v1.3.0 // indirect
	github.com/go-sql-driver/mysql v1.4.1
	github.com/google/go-cmp v0.3.0 // indirect
	github.com/mailru/easyjson v0.0.0-20190403194419-1ea4449da983 // indirect
	github.com/olivere/elastic v6.2.18+incompatible
	github.com/pkg/errors v0.8.1 // indirect
	github.com/sereiner/library v0.0.0-20190516091247-20671f14dec5
	github.com/sereiner/parrot v0.0.0-20190520024652-7acbe377d0a1
	github.com/shopspring/decimal v0.0.0-20180709203117-cd690d0c9e24
	github.com/ugorji/go v1.1.5-pre // indirect
)
