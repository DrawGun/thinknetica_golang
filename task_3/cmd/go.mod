module thinknetica_golang/task_3/main

require pkg/crawler v1.0.0
replace pkg/crawler => ../../task_3/pkg/crawler

require pkg/spider v1.0.0
replace pkg/spider => ../../task_3/pkg/crawler/spider

require pkg/membot v1.0.0
replace pkg/membot => ../../task_3/pkg/crawler/membot

go 1.15
