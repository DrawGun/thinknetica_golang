module thinknetica_golang/task_3/main

require pkg/engine v1.0.0
replace pkg/engine => ../../task_3/pkg/engine

require pkg/crawler v1.0.0
replace pkg/crawler => ../../task_3/pkg/crawler

require pkg/stub v1.0.0
replace pkg/stub => ../../task_3/pkg/stub

go 1.15
