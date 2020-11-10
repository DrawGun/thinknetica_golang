module thinknetica_golang/task_6/main

go 1.15

require pkg/engine v1.0.0
replace pkg/engine => ../../task_6/pkg/engine

require pkg/crawler v1.0.0
replace pkg/crawler => ../../task_6/pkg/crawler

require pkg/crawler/spider v1.0.0
replace pkg/crawler/spider => ../../task_6/pkg/crawler/spider

require pkg/storage v1.0.0
replace pkg/storage => ../../task_6/pkg/storage

require pkg/storage/file v1.0.0
replace pkg/storage/file => ../../task_6/pkg/storage/file

require pkg/index v1.0.0
replace pkg/index => ../../task_6/pkg/index

require pkg/btree v1.0.0
replace pkg/btree => ../../task_6/pkg/btree

require (
	golang.org/x/net v0.0.0-20201031054903-ff519b6c9102 // indirect
	pkg/membot v1.0.0
)

replace pkg/membot => ../../task_6/pkg/crawler/membot
