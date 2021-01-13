module thinknetica_golang/task_16/pkg/engine

go 1.15

require pkg/index v1.0.0
replace pkg/index => ../../pkg/index

require pkg/storage v1.0.0
replace pkg/storage => ../../pkg/storage

require pkg/storage/memstore v1.0.0
replace pkg/storage/memstore => ../../pkg/storage/memstore

require pkg/crawler v1.0.0
replace pkg/crawler => ../../pkg/crawler

require pkg/search v1.0.0
replace pkg/search => ../../pkg/search

require pkg/search/btree v1.0.0
replace pkg/search/btree => ../../pkg/search/btree

require pkg/crawler/stubcrw v1.0.0
replace pkg/crawler/stubcrw => ../crawler/stubcrw

require pkg/index/hash v1.0.0
replace pkg/index/hash => ../index/hash
