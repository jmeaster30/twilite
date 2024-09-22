module github.com/jmeaster30/twilite

go 1.19

replace twilite/twiutil v0.0.0 => ./twiutil
replace twilite/twilib v0.0.0 => ./twilib

require twilite/twiutil v0.0.0
require twilite/twilib v0.0.0
require github.com/mattn/go-sqlite3 v1.14.23 // indirect
