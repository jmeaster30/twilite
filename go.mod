module github.com/jmeaster30/twilite

go 1.19

replace github.com/jmeaster30/twilite/twiutil v0.0.0 => ./twiutil
replace github.com/jmeaster30/twilite/twilib v0.0.0 => ./twilib

require github.com/jmeaster30/twilite/twiutil v0.0.0 // indirect
require github.com/jmeaster30/twilite/twilib v0.0.0
require github.com/mattn/go-sqlite3 v1.14.23 // indirect
