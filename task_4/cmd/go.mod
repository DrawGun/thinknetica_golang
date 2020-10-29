module thinknetica_golang/task_4/main

require pkg/crawler v1.0.0
replace pkg/crawler => ../../task_4/pkg/crawler

require pkg/spider v1.0.0
replace pkg/spider => ../../task_4/pkg/crawler/spider

require pkg/membot v1.0.0
replace pkg/membot => ../../task_4/pkg/crawler/membot

require pkg/index v1.0.0
replace pkg/index => ../../task_4/pkg/index

go 1.15

require pkg/spider v0.0.0-00010101000000-000000000000
