// This example demonstrates usage of FTDI EVE based displays.
//
// It seems that FT800CB-HY50B display is unstable with fast SPI. If you have
// problems please reduce SPI speed or better desolder U1 and U2 (74LCX125
// buffers) and short the U1:2-3,5-6,11-2, U2:2-3,5-6 traces.
package main

import (
	"delay"
	"fmt"
	"rtos"

	"display/eve"
	"display/eve/evetest"
	"display/eve/ft80"

	"stm32/evedci"

	"stm32/hal/dma"
	"stm32/hal/exti"
	"stm32/hal/gpio"
	"stm32/hal/irq"
	"stm32/hal/spi"
	"stm32/hal/system"
	"stm32/hal/system/timer/systick"
)

var dci *evedci.SPI

func init() {
	system.Setup168(8)
	systick.Setup(2e6)

	// GPIO

	gpio.B.EnableClock(true)
	spiport, sck, miso, mosi := gpio.B, gpio.Pin13, gpio.Pin14, gpio.Pin15

	gpio.D.EnableClock(true)
	csn := gpio.D.Pin(8)
	irqn := gpio.D.Pin(9)
	pdn := gpio.D.Pin(10)

	// EVE SPI (Use SPI2. SPI1 is faster but used by onboard accelerometer).

	// Use max. speed for SCK (not necessary for ABP2), see Errata sheet.
	spiport.Setup(sck|mosi, &gpio.Config{Mode: gpio.Alt, Speed: gpio.VeryHigh})
	spiport.Setup(miso, &gpio.Config{Mode: gpio.AltIn})
	spiport.SetAltFunc(sck|miso|mosi, gpio.SPI2)
	d := dma.DMA1
	d.EnableClock(true)
	spidrv := spi.NewDriver(spi.SPI2, d.Channel(3, 0), d.Channel(4, 0))
	spidrv.P.EnableClock(true)
	rtos.IRQ(irq.SPI2).Enable()
	rtos.IRQ(irq.DMA1_Stream3).Enable()
	rtos.IRQ(irq.DMA1_Stream4).Enable()

	// EVE control lines

	cfg := gpio.Config{Mode: gpio.Out, Speed: gpio.High}
	pdn.Setup(&cfg)
	csn.Setup(&cfg)
	irqn.Setup(&gpio.Config{Mode: gpio.In})
	irqline := exti.Lines(irqn.Mask())
	irqline.Connect(irqn.Port())
	irqline.EnableFallTrig()
	irqline.EnableIRQ()
	rtos.IRQ(irq.EXTI9_5).Enable()

	dci = evedci.NewSPI(spidrv, csn, pdn)
}

func curFreq(lcd *eve.Driver) uint32 {
	clk1 := lcd.ReadUint32(ft80.REG_CLOCK)
	t1 := rtos.Nanosec()
	delay.Millisec(8)
	clk2 := lcd.ReadUint32(ft80.REG_CLOCK)
	t2 := rtos.Nanosec()
	return uint32(int64(clk2-clk1) * 1e9 / (t2 - t1))
}

func checkErr(err error) {
	if err == nil {
		return
	}
	fmt.Printf("Error: %v\n", err)
	for {
	}
}

func main() {
	spibus := dci.SPI().P.Bus()
	fmt.Printf("\nSPI on %s (%d MHz).\n", spibus, spibus.Clock()/1e6)
	fmt.Printf("SPI speed: %d bps.\n", dci.SPI().P.Baudrate(dci.SPI().P.Conf()))

	lcd := eve.NewDriver(dci, 128)
	checkErr(lcd.Init(&eve.Default480x272, nil))

	fmt.Printf("EVE clock: %d Hz.\n", curFreq(lcd))
	dci.SetBaudrate(21e6) // Max. for SPI2 and SPI3 < EVE max. 30 MHz
	fmt.Printf("SPI speed: %d bps.\n", dci.SPI().P.Baudrate(dci.SPI().P.Conf()))

	checkErr(evetest.Run(lcd))
	fmt.Printf("End.\n")
}

func lcdSPIISR() {
	dci.SPI().ISR()
}

func lcdRxDMAISR() {
	dci.SPI().DMAISR(dci.SPI().RxDMA)
}

func lcdTxDMAISR() {
	dci.SPI().DMAISR(dci.SPI().TxDMA)
}

func exti9_5ISR() {
	pending := exti.Pending()
	pending &= exti.L5 | exti.L6 | exti.L7 | exti.L8 | exti.L9
	pending.ClearPending()
	if pending&exti.L9 != 0 {
		dci.ISR()
	}
}

//emgo:const
//c:__attribute__((section(".ISRs")))
var ISRs = [...]func(){
	irq.SPI2:         lcdSPIISR,
	irq.DMA1_Stream3: lcdRxDMAISR,
	irq.DMA1_Stream4: lcdTxDMAISR,
	irq.EXTI9_5:      exti9_5ISR,
}
