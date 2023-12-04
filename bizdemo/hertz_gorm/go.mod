module github.com/cloudwego/hertz-examples/bizdemo/hertz_gorm

go 1.16

replace github.com/apache/thrift => github.com/apache/thrift v0.13.0

require (
	github.com/apache/thrift v0.16.0
	github.com/cloudwego/hertz v0.7.2
	github.com/hertz-contrib/gzip v0.0.3
	gorm.io/driver/mysql v1.5.2
	gorm.io/gorm v1.25.5
)
