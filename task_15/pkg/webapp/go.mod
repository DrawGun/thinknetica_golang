module thinknetica_golang/task_15/pkg/webapp

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

replace pkg/index/hash => ../index/hash

replace pkg/crawler => ../../pkg/crawler

require pkg/crawler/stubcrw v1.0.0

replace pkg/crawler/stubcrw => ../crawler/stubcrw
