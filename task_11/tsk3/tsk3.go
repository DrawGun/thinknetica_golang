// Package tsk3 реализует решение для задания 3
package tsk3

import "io"

func write(w io.Writer, values ...interface{}) {
	for _, v := range values {
		if str, ok := v.(string); ok {
			w.Write([]byte(str))
		}
	}
}
