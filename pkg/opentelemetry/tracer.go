package opentelemetry

import "go.opentelemetry.io/otel"

var Tracer = otel.Tracer("rinha-backend")

//
//type Tracer struct {
//}
//
//func (t Tracer) Start() *Tracer {
//
//
//
//	return &Tracer{}
//}
