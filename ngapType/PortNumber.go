package ngapType

import "free5gc/lib/aper"

// Need to import "free5gc/lib/aper" if it uses "aper"

type PortNumber struct {
	Value aper.OctetString `aper:"sizeLB:2,sizeUB:2"`
}
