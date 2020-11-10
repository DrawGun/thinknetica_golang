module thinknetica_golang/task_6/pkg/engine

go 1.15

require pkg/crawler v1.0.0
replace pkg/crawler => ../../pkg/crawler

require pkg/storage v1.0.0
replace pkg/storage => ../../pkg/storage

require pkg/storage/file v1.0.0
replace pkg/storage/file => ../../pkg/storage/file

require pkg/index v1.0.0
replace pkg/index => ../../pkg/index

require pkg/btree v1.0.0
replace pkg/btree => ../../pkg/btree
