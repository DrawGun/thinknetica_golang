module thinknetica_golang/task_9/pkg/storage/memstore

go 1.15

require pkg/crawler v1.0.0
replace pkg/crawler => ../../crawler

require pkg/storage/teststore v1.0.0
replace pkg/storage/teststore => ../teststore