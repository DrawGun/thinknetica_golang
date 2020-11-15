module thinknetica_golang/task_7/gosearch

go 1.15

require pkg/btree v1.0.0
replace pkg/btree => ../../../task_7/pkg/btree

require pkg/crawler v1.0.0
replace pkg/crawler => ../../../task_7/pkg/crawler

require pkg/crawler/spider v1.0.0
replace pkg/crawler/spider => ../../../task_7/pkg/crawler/spider

require pkg/engine v1.0.0
replace pkg/engine => ../../../task_7/pkg/engine

require pkg/index v1.0.0
replace pkg/index => ../../../task_7/pkg/index

require pkg/index/hash v1.0.0
replace pkg/index/hash => ../../../task_7/pkg/index/hash

require pkg/storage v1.0.0
replace pkg/storage => ../../../task_7/pkg/storage

require pkg/storage/memstore v1.0.0
replace pkg/storage/memstore => ../../../task_7/pkg/storage/memstore
