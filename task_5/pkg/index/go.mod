module thinknetica_golang/task_5/pkg/index

require pkg/crawler v1.0.0
replace pkg/crawler => ../../pkg/crawler

require index/bst v1.0.0
replace index/bst => ./bst

go 1.15