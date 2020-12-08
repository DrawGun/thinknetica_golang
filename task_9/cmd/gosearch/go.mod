module thinknetica_golang/task_9/gosearch

go 1.15

require pkg/search v1.0.0
replace pkg/search => ../../../task_9/pkg/search

require pkg/search/btree v1.0.0
replace pkg/search/btree => ../../../task_9/pkg/search/btree

require pkg/search/array v1.0.0
replace pkg/search/array => ../../../task_9/pkg/search/array

require pkg/crawler v1.0.0
replace pkg/crawler => ../../../task_9/pkg/crawler

require pkg/crawler/spider v1.0.0
replace pkg/crawler/spider => ../../../task_9/pkg/crawler/spider

require pkg/crawler/stubcrw v1.0.0
replace pkg/crawler/stubcrw => ../../../task_9/pkg/crawler/stubcrw

require pkg/engine v1.0.0
replace pkg/engine => ../../../task_9/pkg/engine

require pkg/index v1.0.0
replace pkg/index => ../../../task_9/pkg/index

require pkg/index/hash v1.0.0
replace pkg/index/hash => ../../../task_9/pkg/index/hash

require pkg/storage v1.0.0
replace pkg/storage => ../../../task_9/pkg/storage

require pkg/storage/memstore v1.0.0
replace pkg/storage/memstore => ../../../task_9/pkg/storage/memstore

require pkg/storage/teststore v1.0.0
replace pkg/storage/teststore => ../../../task_9/pkg/storage/teststore
