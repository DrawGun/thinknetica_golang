module thinknetica_golang/task_14/gosearch

go 1.15

require pkg/search v1.0.0
replace pkg/search => ../../../task_14/pkg/search

require pkg/search/btree v1.0.0
replace pkg/search/btree => ../../../task_14/pkg/search/btree

require pkg/crawler v1.0.0
replace pkg/crawler => ../../../task_14/pkg/crawler

require pkg/crawler/spider v1.0.0
replace pkg/crawler/spider => ../../../task_14/pkg/crawler/spider

require pkg/crawler/stubcrw v1.0.0
replace pkg/crawler/stubcrw => ../../../task_14/pkg/crawler/stubcrw

require pkg/engine v1.0.0
replace pkg/engine => ../../../task_14/pkg/engine

require pkg/index v1.0.0
replace pkg/index => ../../../task_14/pkg/index

require pkg/index/hash v1.0.0
replace pkg/index/hash => ../../../task_14/pkg/index/hash

require pkg/storage v1.0.0
replace pkg/storage => ../../../task_14/pkg/storage

require pkg/storage/memstore v1.0.0
replace pkg/storage/memstore => ../../../task_14/pkg/storage/memstore

require pkg/netsrv v1.0.0
replace pkg/netsrv => ../../../task_14/pkg/netsrv
