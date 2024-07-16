package cronwithcontext

import (
	"context"
	"fmt"
	"testing"
)

func WithNewContext() JobWrapper {
	return func(j Job) Job {
		return FuncJob(func(cCtx *Context) {
			if cCtx.GetID() == 1 {
				ctx := context.WithValue(cCtx.GetContext(), "aaa", "111")
				cCtx.SetContext(ctx)
			}

			if cCtx.GetID() == 2 {
				ctx := context.WithValue(cCtx.GetContext(), "aaa", "2222")
				cCtx.SetContext(ctx)
			}
			fmt.Println("WithNewContext Id:", cCtx.GetID())
			j.Run(cCtx)
		})
	}
}

func TestCron(t *testing.T) {
	c := New(WithSeconds(), WithChain(WithNewContext()))
	aId, _ := c.AddFunc("* * * * * ?", func(cCtx *Context) {
		v := cCtx.GetContext().Value("aaa")
		t.Log("a Value: ", v)
		t.Log("a END")
	})
	fmt.Println("aId:  ", aId)

	bId, _ := c.AddFunc("* * * * * ?", func(cCtx *Context) {
		v := cCtx.GetContext().Value("aaa")
		t.Log("b Value: ", v)
		t.Log("END")
	})
	fmt.Println("bId:  ", bId)

	c.Start()
	select {}
}
