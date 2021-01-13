module thinknetica_golang/task_16/gosearch

go 1.15

require pkg/search v1.0.0

replace pkg/search => ../../../task_16/pkg/search

require pkg/search/btree v1.0.0

replace pkg/search/btree => ../../../task_16/pkg/search/btree

require pkg/crawler v1.0.0

replace pkg/crawler => ../../../task_16/pkg/crawler

require pkg/crawler/spider v1.0.0

replace pkg/crawler/spider => ../../../task_16/pkg/crawler/spider

require pkg/crawler/stubcrw v1.0.0

replace pkg/crawler/stubcrw => ../../../task_16/pkg/crawler/stubcrw

require pkg/engine v1.0.0

replace pkg/engine => ../../../task_16/pkg/engine

require pkg/index v1.0.0

replace pkg/index => ../../../task_16/pkg/index

require pkg/index/hash v1.0.0

replace pkg/index/hash => ../../../task_16/pkg/index/hash

require pkg/storage v1.0.0

replace pkg/storage => ../../../task_16/pkg/storage

require pkg/storage/memstore v1.0.0

replace pkg/storage/memstore => ../../../task_16/pkg/storage/memstore

require pkg/webapp v1.0.0

replace pkg/webapp => ../../../task_16/pkg/webapp

require (
	github.com/gorilla/mux v1.8.0
	pkg/api v1.0.0
)

replace pkg/api => ../../../task_16/pkg/api
