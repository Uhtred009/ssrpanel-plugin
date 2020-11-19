module ssrpanel-plugin

go 1.15

require (
	github.com/Uhtred009/v2ray-core-1 v3.50.2+incompatible
	github.com/dgryski/go-metro v0.0.0-20200812162917-85c65e2d0165 // indirect
	github.com/go-sql-driver/mysql v1.5.0
	github.com/gorilla/websocket v1.4.2 // indirect
	github.com/jinzhu/gorm v1.9.16
	github.com/lucas-clemente/quic-go v0.19.1 // indirect
	github.com/pires/go-proxyproto v0.3.2 // indirect
	github.com/pkg/errors v0.9.1
	github.com/robfig/cron v1.2.0
	github.com/seiflotfy/cuckoofilter v0.0.0-20201009151232-afb285a456ab // indirect
	github.com/shirou/gopsutil v3.20.10+incompatible
	github.com/sirupsen/logrus v1.7.0
	github.com/tidwall/gjson v1.6.3
	github.com/xtls/go v0.0.0-20201118062508-3632bf3b7499 // indirect
	go.starlark.net v0.0.0-20201113214410-e292e66a28cd // indirect
	google.golang.org/grpc v1.33.2
	gopkg.in/resty.v1 v1.12.0
	v2ray.com/core v4.19.1+incompatible
)

replace v2ray.com/core => github.com/Uhtred009/v2ray-core-1 v4.31.2+incompatible
