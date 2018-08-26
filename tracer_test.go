package trace

import (
	"testing"
	"bytes"
)

func TestNew(t *testing.T) {
	var buf bytes.Buffer
	tracer := New(&buf)
	if tracer == nil {
		t.Error("tracer is nil")
	}
	traceStr := "Hello, trace package"
	tracer.Trace(traceStr)
	if buf.String() != traceStr + "\n" {
		t.Errorf("Trace should not write '%s'.", buf.String())
	}
}
