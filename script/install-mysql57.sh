
CREATE DATABASE IF NOT EXISTS mozi CHARSET UTF8;


go build startup_admin.go
go build startup_api.go
go build startup_lotto.go
go build startup_cron.go