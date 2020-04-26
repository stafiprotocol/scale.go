package scalecodec_test

import (
	"github.com/freehere107/scalecodec"
	"github.com/freehere107/scalecodec/source"
	"github.com/freehere107/scalecodec/types"
	"github.com/freehere107/scalecodec/utiles"
	"testing"
)

func TestEventsDecoder(t *testing.T) {
	m := scalecodec.MetadataDecoder{}
	m.Init(utiles.HexToBytes(Kusama1055))
	_ = m.Process()
	types.RegCustomTypes(source.LoadTypeRegistry("source/base"))

	e := scalecodec.EventsDecoder{}
	option := types.ScaleDecoderOption{Metadata: &m.Metadata}

	eventRaw := "0x10020f001400000000000000000000001027000001010000010000000000102700000001000002000000000040420f00000100"
	e.Init(types.ScaleBytes{Data: utiles.HexToBytes(eventRaw)}, &option)
	e.Process()
}