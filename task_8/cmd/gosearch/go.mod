module thinknetica_golang/task_8/gosearch

go 1.15

require pkg/btree v1.0.0
replace pkg/btree => ../../../task_8/pkg/btree

require pkg/crawler v1.0.0
replace pkg/crawler => ../../../task_8/pkg/crawler

require pkg/crawler/spider v1.0.0
replace pkg/crawler/spider => ../../../task_8/pkg/crawler/spider

require pkg/crawler/stubcrw v1.0.0
replace pkg/crawler/stubcrw => ../../../task_8/pkg/crawler/stubcrw

require pkg/engine v1.0.0
replace pkg/engine => ../../../task_8/pkg/engine

require pkg/index v1.0.0
replace pkg/index => ../../../task_8/pkg/index

require pkg/index/hash v1.0.0
replace pkg/index/hash => ../../../task_8/pkg/index/hash

require pkg/storage v1.0.0
replace pkg/storage => ../../../task_8/pkg/storage

require pkg/storage/memstore v1.0.0
replace pkg/storage/memstore => ../../../task_8/pkg/storage/memstore

require pkg/storage/teststore v1.0.0
replace pkg/storage/teststore => ../../../task_8/pkg/storage/teststore
