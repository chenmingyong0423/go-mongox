// Generated by [optioner] command-line tool; DO NOT EDIT
// If you have any questions, please create issues and submit contributions at:
// https://github.com/chenmingyong0423/go-optioner

package finder

import "go.mongodb.org/mongo-driver/mongo"

type OpContextOption func(*OpContext)

func NewOpContext(col *mongo.Collection, filter any, opts ...OpContextOption) *OpContext {
	opContext := &OpContext{
		Col:    col,
		Filter: filter,
	}

	for _, opt := range opts {
		opt(opContext)
	}

	return opContext
}

func WithUpdates(updates any) OpContextOption {
	return func(opContext *OpContext) {
		opContext.Updates = updates
	}
}

func WithMongoOptions(mongoOptions any) OpContextOption {
	return func(opContext *OpContext) {
		opContext.MongoOptions = mongoOptions
	}
}
