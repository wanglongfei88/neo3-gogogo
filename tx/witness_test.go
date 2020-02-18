package tx

import (
	"github.com/joeqian10/neo3-gogogo/helper"
	"github.com/joeqian10/neo3-gogogo/helper/io"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWitness_Deserialize(t *testing.T) {
	s := "41" + "40915467ecd359684b2dc358024ca750609591aa731a0b309c7fb3cab5cd0836ad3992aa0a24da431f43b68883ea5651d548feb6bd3c8e16376e6e426f91f84c58" +
		"23" + "2103322f35c7819267e721335948d385fae5be66e7ba8c748ac15467dcca0693692dac"
	br := io.NewBinaryReaderFromBuf(helper.HexTobytes(s))
	w := Witness{}
	w.Deserialize(br)
	assert.Equal(t, "40915467ecd359684b2dc358024ca750609591aa731a0b309c7fb3cab5cd0836ad3992aa0a24da431f43b68883ea5651d548feb6bd3c8e16376e6e426f91f84c58", helper.BytesToHex(w.InvocationScript))
	assert.Equal(t, "2103322f35c7819267e721335948d385fae5be66e7ba8c748ac15467dcca0693692dac", helper.BytesToHex(w.VerificationScript))
}

func TestWitness_GetScriptHash(t *testing.T) {
	w := Witness{
		InvocationScript:   helper.HexTobytes("40915467ecd359684b2dc358024ca750609591aa731a0b309c7fb3cab5cd0836ad3992aa0a24da431f43b68883ea5651d548feb6bd3c8e16376e6e426f91f84c58"), //65
		VerificationScript: helper.HexTobytes("2103322f35c7819267e721335948d385fae5be66e7ba8c748ac15467dcca0693692dac"),                                                             //35
	}
	scriptHash := w.GetScriptHash()
	assert.Equal(t, "71cb588c8291c18fa87fa07ce16c3fd92ab5aa30", scriptHash.String())
}

func TestWitness_Serialize(t *testing.T) {
	w := Witness{
		InvocationScript:   helper.HexTobytes("40915467ecd359684b2dc358024ca750609591aa731a0b309c7fb3cab5cd0836ad3992aa0a24da431f43b68883ea5651d548feb6bd3c8e16376e6e426f91f84c58"), //65
		VerificationScript: helper.HexTobytes("2103322f35c7819267e721335948d385fae5be66e7ba8c748ac15467dcca0693692dac"),                                                             //35
	}
	bbw := io.NewBufBinaryWriter()
	w.Serialize(bbw.BinaryWriter)
	b := bbw.Bytes()
	assert.Equal(t, "41"+"40915467ecd359684b2dc358024ca750609591aa731a0b309c7fb3cab5cd0836ad3992aa0a24da431f43b68883ea5651d548feb6bd3c8e16376e6e426f91f84c58"+
		"23"+"2103322f35c7819267e721335948d385fae5be66e7ba8c748ac15467dcca0693692dac", helper.BytesToHex(b))
}

func TestWitness_Size(t *testing.T) {
	w := Witness{
		InvocationScript:   helper.HexTobytes("40915467ecd359684b2dc358024ca750609591aa731a0b309c7fb3cab5cd0836ad3992aa0a24da431f43b68883ea5651d548feb6bd3c8e16376e6e426f91f84c58"), //65
		VerificationScript: helper.HexTobytes("2103322f35c7819267e721335948d385fae5be66e7ba8c748ac15467dcca0693692dac"),                                                             //35
	}
	size := w.Size()
	assert.Equal(t, 1+65+1+35, size)
}
