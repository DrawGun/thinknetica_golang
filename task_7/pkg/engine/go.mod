module thinknetica_golang/task_7/pkg/engine

go 1.15

require pkg/index v1.0.0
replace pkg/index => ../../pkg/index

require pkg/storage v1.0.0
replace pkg/storage => ../../pkg/storage

require pkg/crawler v1.0.0
replace pkg/crawler => ../../pkg/crawler

require pkg/btree v1.0.0
replace pkg/btree => ../../pkg/btree