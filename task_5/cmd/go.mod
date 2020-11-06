module thinknetica_golang/task_5/main

require pkg/crawler v1.0.0
replace pkg/crawler => ../../task_5/pkg/crawler

require pkg/spider v1.0.0
replace pkg/spider => ../../task_5/pkg/crawler/spider

require pkg/membot v1.0.0
replace pkg/membot => ../../task_5/pkg/crawler/membot

require pkg/index v1.0.0
replace pkg/index => ../../task_5/pkg/index

require index/bst v1.0.0
replace index/bst => ../../task_5/pkg/index/bst

go 1.15