module thinknetica_golang/task_16/pkg/api

go 1.15

require github.com/gorilla/mux v1.8.0

require pkg/crawler v1.0.0

require pkg/storage v1.0.0

replace pkg/storage => ../../pkg/storage

require pkg/storage/memstore v1.0.0

replace pkg/storage/memstore => ../../pkg/storage/memstore

require pkg/index v1.0.0

replace pkg/index => ../../pkg/index

require pkg/index/hash v1.0.0

replace pkg/crawler => ../../pkg/crawler

require pkg/crawler/stubcrw v1.0.0

require pkg/engine v1.0.0

replace pkg/engine => ../../pkg/engine

require pkg/search v1.0.0

replace pkg/search => ../../pkg/search

require pkg/search/btree v1.0.0

replace pkg/search/btree => ../../pkg/search/btree

replace pkg/crawler/stubcrw => ../crawler/stubcrw

replace pkg/index/hash => ../index/hash
