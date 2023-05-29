package goduplicateword

import (
	"fmt"
	"testing"
)

func TestProcessMapReduce(t *testing.T) {

	processMapReduce := ProcessMapReduce("xx vv xx")

	if nil == processMapReduce {
		t.Error("processMapReduce is nil")
	}

	if 2 != processMapReduce["xx"] {
		t.Error("xx is not 2")
	}

	if 1 != processMapReduce["vv"] {
		t.Error("vv  is not 1")
	}

	fmt.Println("map:", processMapReduce)
}
