module thinknetica_golang/task_8/pkg/engine

go 1.15

require pkg/index v1.0.0
replace pkg/index => ../../pkg/index

require pkg/storage v1.0.0
replace pkg/storage => ../../pkg/storage

require pkg/storage/teststore v1.0.0
replace pkg/storage/teststore => ../storage/teststore

require pkg/crawler v1.0.0
replace pkg/crawler => ../../pkg/crawler

require pkg/btree v1.0.0
replace pkg/btree => ../../pkg/btree

require pkg/crawler/stubcrw v1.0.0
replace pkg/crawler/stubcrw => ../crawler/stubcrw

require pkg/index/hash v1.0.0
replace pkg/index/hash => ../index/hash
