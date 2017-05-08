// +build l476xx

package adc

// DO NOT EDIT THIS FILE. GENERATED BY xgen.

import (
	"bits"
	"mmio"
	"unsafe"

	"stm32/o/l476xx/mmap"
)

type ADC_Common_Periph struct {
	CSR CSR
	_   uint32
	CCR CCR
	CDR CDR
}

func (p *ADC_Common_Periph) BaseAddr() uintptr {
	return uintptr(unsafe.Pointer(p))
}

//emgo:const
var ADC123_COMMON = (*ADC_Common_Periph)(unsafe.Pointer(uintptr(mmap.ADC123_COMMON_BASE)))

type CSR_Bits uint32

func (b CSR_Bits) Field(mask CSR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask CSR_Bits) J(v int) CSR_Bits {
	return CSR_Bits(bits.Make32(v, uint32(mask)))
}

type CSR struct{ mmio.U32 }

func (r *CSR) Bits(mask CSR_Bits) CSR_Bits { return CSR_Bits(r.U32.Bits(uint32(mask))) }
func (r *CSR) StoreBits(mask, b CSR_Bits)  { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *CSR) SetBits(mask CSR_Bits)       { r.U32.SetBits(uint32(mask)) }
func (r *CSR) ClearBits(mask CSR_Bits)     { r.U32.ClearBits(uint32(mask)) }
func (r *CSR) Load() CSR_Bits              { return CSR_Bits(r.U32.Load()) }
func (r *CSR) Store(b CSR_Bits)            { r.U32.Store(uint32(b)) }

type CSR_Mask struct{ mmio.UM32 }

func (rm CSR_Mask) Load() CSR_Bits   { return CSR_Bits(rm.UM32.Load()) }
func (rm CSR_Mask) Store(b CSR_Bits) { rm.UM32.Store(uint32(b)) }

func (p *ADC_Common_Periph) ADRDY_MST() CSR_Mask {
	return CSR_Mask{mmio.UM32{&p.CSR.U32, uint32(ADRDY_MST)}}
}

func (p *ADC_Common_Periph) EOSMP_MST() CSR_Mask {
	return CSR_Mask{mmio.UM32{&p.CSR.U32, uint32(EOSMP_MST)}}
}

func (p *ADC_Common_Periph) EOC_MST() CSR_Mask {
	return CSR_Mask{mmio.UM32{&p.CSR.U32, uint32(EOC_MST)}}
}

func (p *ADC_Common_Periph) EOS_MST() CSR_Mask {
	return CSR_Mask{mmio.UM32{&p.CSR.U32, uint32(EOS_MST)}}
}

func (p *ADC_Common_Periph) OVR_MST() CSR_Mask {
	return CSR_Mask{mmio.UM32{&p.CSR.U32, uint32(OVR_MST)}}
}

func (p *ADC_Common_Periph) JEOC_MST() CSR_Mask {
	return CSR_Mask{mmio.UM32{&p.CSR.U32, uint32(JEOC_MST)}}
}

func (p *ADC_Common_Periph) JEOS_MST() CSR_Mask {
	return CSR_Mask{mmio.UM32{&p.CSR.U32, uint32(JEOS_MST)}}
}

func (p *ADC_Common_Periph) AWD1_MST() CSR_Mask {
	return CSR_Mask{mmio.UM32{&p.CSR.U32, uint32(AWD1_MST)}}
}

func (p *ADC_Common_Periph) AWD2_MST() CSR_Mask {
	return CSR_Mask{mmio.UM32{&p.CSR.U32, uint32(AWD2_MST)}}
}

func (p *ADC_Common_Periph) AWD3_MST() CSR_Mask {
	return CSR_Mask{mmio.UM32{&p.CSR.U32, uint32(AWD3_MST)}}
}

func (p *ADC_Common_Periph) JQOVF_MST() CSR_Mask {
	return CSR_Mask{mmio.UM32{&p.CSR.U32, uint32(JQOVF_MST)}}
}

func (p *ADC_Common_Periph) ADRDY_SLV() CSR_Mask {
	return CSR_Mask{mmio.UM32{&p.CSR.U32, uint32(ADRDY_SLV)}}
}

func (p *ADC_Common_Periph) EOSMP_SLV() CSR_Mask {
	return CSR_Mask{mmio.UM32{&p.CSR.U32, uint32(EOSMP_SLV)}}
}

func (p *ADC_Common_Periph) EOC_SLV() CSR_Mask {
	return CSR_Mask{mmio.UM32{&p.CSR.U32, uint32(EOC_SLV)}}
}

func (p *ADC_Common_Periph) EOS_SLV() CSR_Mask {
	return CSR_Mask{mmio.UM32{&p.CSR.U32, uint32(EOS_SLV)}}
}

func (p *ADC_Common_Periph) OVR_SLV() CSR_Mask {
	return CSR_Mask{mmio.UM32{&p.CSR.U32, uint32(OVR_SLV)}}
}

func (p *ADC_Common_Periph) JEOC_SLV() CSR_Mask {
	return CSR_Mask{mmio.UM32{&p.CSR.U32, uint32(JEOC_SLV)}}
}

func (p *ADC_Common_Periph) JEOS_SLV() CSR_Mask {
	return CSR_Mask{mmio.UM32{&p.CSR.U32, uint32(JEOS_SLV)}}
}

func (p *ADC_Common_Periph) AWD1_SLV() CSR_Mask {
	return CSR_Mask{mmio.UM32{&p.CSR.U32, uint32(AWD1_SLV)}}
}

func (p *ADC_Common_Periph) AWD2_SLV() CSR_Mask {
	return CSR_Mask{mmio.UM32{&p.CSR.U32, uint32(AWD2_SLV)}}
}

func (p *ADC_Common_Periph) AWD3_SLV() CSR_Mask {
	return CSR_Mask{mmio.UM32{&p.CSR.U32, uint32(AWD3_SLV)}}
}

func (p *ADC_Common_Periph) JQOVF_SLV() CSR_Mask {
	return CSR_Mask{mmio.UM32{&p.CSR.U32, uint32(JQOVF_SLV)}}
}

type CCR_Bits uint32

func (b CCR_Bits) Field(mask CCR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask CCR_Bits) J(v int) CCR_Bits {
	return CCR_Bits(bits.Make32(v, uint32(mask)))
}

type CCR struct{ mmio.U32 }

func (r *CCR) Bits(mask CCR_Bits) CCR_Bits { return CCR_Bits(r.U32.Bits(uint32(mask))) }
func (r *CCR) StoreBits(mask, b CCR_Bits)  { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *CCR) SetBits(mask CCR_Bits)       { r.U32.SetBits(uint32(mask)) }
func (r *CCR) ClearBits(mask CCR_Bits)     { r.U32.ClearBits(uint32(mask)) }
func (r *CCR) Load() CCR_Bits              { return CCR_Bits(r.U32.Load()) }
func (r *CCR) Store(b CCR_Bits)            { r.U32.Store(uint32(b)) }

type CCR_Mask struct{ mmio.UM32 }

func (rm CCR_Mask) Load() CCR_Bits   { return CCR_Bits(rm.UM32.Load()) }
func (rm CCR_Mask) Store(b CCR_Bits) { rm.UM32.Store(uint32(b)) }

func (p *ADC_Common_Periph) DUAL() CCR_Mask {
	return CCR_Mask{mmio.UM32{&p.CCR.U32, uint32(DUAL)}}
}

func (p *ADC_Common_Periph) DELAY() CCR_Mask {
	return CCR_Mask{mmio.UM32{&p.CCR.U32, uint32(DELAY)}}
}

func (p *ADC_Common_Periph) MDMACFG() CCR_Mask {
	return CCR_Mask{mmio.UM32{&p.CCR.U32, uint32(MDMACFG)}}
}

func (p *ADC_Common_Periph) MDMA() CCR_Mask {
	return CCR_Mask{mmio.UM32{&p.CCR.U32, uint32(MDMA)}}
}

func (p *ADC_Common_Periph) CKMODE() CCR_Mask {
	return CCR_Mask{mmio.UM32{&p.CCR.U32, uint32(CKMODE)}}
}

func (p *ADC_Common_Periph) PRESC() CCR_Mask {
	return CCR_Mask{mmio.UM32{&p.CCR.U32, uint32(PRESC)}}
}

func (p *ADC_Common_Periph) VREFEN() CCR_Mask {
	return CCR_Mask{mmio.UM32{&p.CCR.U32, uint32(VREFEN)}}
}

func (p *ADC_Common_Periph) TSEN() CCR_Mask {
	return CCR_Mask{mmio.UM32{&p.CCR.U32, uint32(TSEN)}}
}

func (p *ADC_Common_Periph) VBATEN() CCR_Mask {
	return CCR_Mask{mmio.UM32{&p.CCR.U32, uint32(VBATEN)}}
}

type CDR_Bits uint32

func (b CDR_Bits) Field(mask CDR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask CDR_Bits) J(v int) CDR_Bits {
	return CDR_Bits(bits.Make32(v, uint32(mask)))
}

type CDR struct{ mmio.U32 }

func (r *CDR) Bits(mask CDR_Bits) CDR_Bits { return CDR_Bits(r.U32.Bits(uint32(mask))) }
func (r *CDR) StoreBits(mask, b CDR_Bits)  { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *CDR) SetBits(mask CDR_Bits)       { r.U32.SetBits(uint32(mask)) }
func (r *CDR) ClearBits(mask CDR_Bits)     { r.U32.ClearBits(uint32(mask)) }
func (r *CDR) Load() CDR_Bits              { return CDR_Bits(r.U32.Load()) }
func (r *CDR) Store(b CDR_Bits)            { r.U32.Store(uint32(b)) }

type CDR_Mask struct{ mmio.UM32 }

func (rm CDR_Mask) Load() CDR_Bits   { return CDR_Bits(rm.UM32.Load()) }
func (rm CDR_Mask) Store(b CDR_Bits) { rm.UM32.Store(uint32(b)) }

func (p *ADC_Common_Periph) RDATA_MST() CDR_Mask {
	return CDR_Mask{mmio.UM32{&p.CDR.U32, uint32(RDATA_MST)}}
}

func (p *ADC_Common_Periph) RDATA_SLV() CDR_Mask {
	return CDR_Mask{mmio.UM32{&p.CDR.U32, uint32(RDATA_SLV)}}
}