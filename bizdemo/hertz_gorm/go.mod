module github.com/cloudwego/hertz-examples/bizdemo/hertz_gorm

go 1.16

replace github.com/apache/thrift => github.com/apache/thrift v0.13.0

require (
	github.com/apache/thrift v0.0.0-00010101000000-000000000000
	github.com/cloudwego/hertz v0.3.2
	github.com/hertz-contrib/gzip v0.0.1
	gorm.io/driver/mysql v1.3.6
	gorm.io/gorm v1.23.10
)
