package packet

import (
	"testing"
)

func generatePacketAF(t *testing.T, AFString string) (Packet, AdaptationField) {
	p := createPacketEmptyAdaptationField(t, "47000030"+AFString)
	a := p.AdaptationField()
	if a == nil {
		t.Errorf("failed to get adaptation field")
	}
	return p, a
}

func TestDiscontinuity(t *testing.T) {
	_, a := generatePacketAF(t, "0180")
	if !a.Discontinuity() {
		t.Errorf("failed to read discontinuity correctly.")
	}
	_, a = generatePacketAF(t, "0190")
	if !a.Discontinuity() {
		t.Errorf("failed to read discontinuity correctly.")
	}
	_, a = generatePacketAF(t, "0170")
	if a.Discontinuity() {
		t.Errorf("failed to read discontinuity correctly.")
	}
}

func TestAdaptationField(t *testing.T) {
	p := createPacketEmptyBody(t, "470000300102")
	a := p.AdaptationField()
	if a == nil {
		t.Errorf("Error getting adaptation field.")
	}

	p = createPacketEmptyBody(t, "470000100002")
	a = p.AdaptationField()
	if a != nil {
		t.Errorf("Adaptation field does not exist but something was returned.")
	}
}

func TestAll(t *testing.T) {
	generated, a := generatePacketAF(t, "0100")
	target, _ := generatePacketAF(t, "B710000000007E01")
	a.SetHasPCR(true)
	a.SetPCR(1)
	assertPacket(t, target, generated)

	target, _ = generatePacketAF(t, "B718000000007E01000000007E02")
	a.SetHasOPCR(true)
	a.SetOPCR(2)
	assertPacket(t, target, generated)

	target, _ = generatePacketAF(t, "B71A000000007E01000000007E020188")
	a.SetHasTransportPrivateData(true)
	a.SetTransportPrivateData([]byte{0x88})
	assertPacket(t, target, generated)

	target, _ = generatePacketAF(t, "B71B000000007E01000000007E0201880177")
	a.SetHasAdaptationFieldExtension(true)
	a.SetAdaptationFieldExtension([]byte{0x77})
	assertPacket(t, target, generated)

	target, _ = generatePacketAF(t, "B71B000000007E01000000007E020266660177")
	a.SetTransportPrivateData([]byte{0x66, 0x66})
	assertPacket(t, target, generated)

	target, _ = generatePacketAF(t, "B713000000007E010266660177")
	a.SetHasOPCR(false)
	assertPacket(t, target, generated)

}
