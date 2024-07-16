package cronwithcontext

import (
	"context"
)

type Context struct {
	id  EntryID
	ctx context.Context
}

func NewContext(ctx context.Context, id EntryID) *Context {
	return &Context{ctx: ctx, id: id}
}

func (c *Context) Copy() *Context {
	return &Context{ctx: c.ctx, id: c.id}
}

func (c *Context) SetContext(ctx context.Context) {
	c.ctx = ctx
}

func (c *Context) GetContext() context.Context {
	return c.ctx
}

func (c *Context) GetID() EntryID {
	return c.id
}
