module thinknetica_golang/task_13/gosearch

go 1.15

require pkg/search v1.0.0
replace pkg/search => ../../../task_13/pkg/search

require pkg/search/btree v1.0.0
replace pkg/search/btree => ../../../task_13/pkg/search/btree

require pkg/crawler v1.0.0
replace pkg/crawler => ../../../task_13/pkg/crawler

require pkg/crawler/spider v1.0.0
replace pkg/crawler/spider => ../../../task_13/pkg/crawler/spider

require pkg/crawler/stubcrw v1.0.0
replace pkg/crawler/stubcrw => ../../../task_13/pkg/crawler/stubcrw

require pkg/engine v1.0.0
replace pkg/engine => ../../../task_13/pkg/engine

require pkg/index v1.0.0
replace pkg/index => ../../../task_13/pkg/index

require pkg/index/hash v1.0.0
replace pkg/index/hash => ../../../task_13/pkg/index/hash

require pkg/storage v1.0.0
replace pkg/storage => ../../../task_13/pkg/storage

require pkg/storage/memstore v1.0.0
replace pkg/storage/memstore => ../../../task_13/pkg/storage/memstore

require pkg/storage/array v1.0.0
replace pkg/storage/array => ../../../task_13/pkg/storage/array
