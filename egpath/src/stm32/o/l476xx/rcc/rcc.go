// Peripheral: RCC_Periph  Reset and Clock Control.
// Instances:
//  RCC  mmap.RCC_BASE
// Registers:
//  0x00 32  CR          Clock control register.
//  0x04 32  ICSCR       Internal clock sources calibration register.
//  0x08 32  CFGR        Clock configuration register.
//  0x0C 32  PLLCFGR     System PLL configuration register.
//  0x10 32  PLLSAI1CFGR PLL SAI1 configuration register.
//  0x14 32  PLLSAI2CFGR PLL SAI2 configuration register.
//  0x18 32  CIER        Clock interrupt enable register.
//  0x1C 32  CIFR        Clock interrupt flag register.
//  0x20 32  CICR        Clock interrupt clear register.
//  0x28 32  AHB1RSTR    AHB1 peripheral reset register.
//  0x2C 32  AHB2RSTR    AHB2 peripheral reset register.
//  0x30 32  AHB3RSTR    AHB3 peripheral reset register.
//  0x38 32  APB1RSTR1   APB1 peripheral reset register 1.
//  0x3C 32  APB1RSTR2   APB1 peripheral reset register 2.
//  0x40 32  APB2RSTR    APB2 peripheral reset register.
//  0x48 32  AHB1ENR     AHB1 peripheral clocks enable register.
//  0x4C 32  AHB2ENR     AHB2 peripheral clocks enable register.
//  0x50 32  AHB3ENR     AHB3 peripheral clocks enable register.
//  0x58 32  APB1ENR1    APB1 peripheral clocks enable register 1.
//  0x5C 32  APB1ENR2    APB1 peripheral clocks enable register 2.
//  0x60 32  APB2ENR     APB2 peripheral clocks enable register.
//  0x68 32  AHB1SMENR   AHB1 peripheral clocks enable in sleep and stop modes register.
//  0x6C 32  AHB2SMENR   AHB2 peripheral clocks enable in sleep and stop modes register.
//  0x70 32  AHB3SMENR   AHB3 peripheral clocks enable in sleep and stop modes register.
//  0x78 32  APB1SMENR1  APB1 peripheral clocks enable in sleep mode and stop modes register 1.
//  0x7C 32  APB1SMENR2  APB1 peripheral clocks enable in sleep mode and stop modes register 2.
//  0x80 32  APB2SMENR   APB2 peripheral clocks enable in sleep mode and stop modes register.
//  0x88 32  CCIPR       Peripherals independent clock configuration register.
//  0x90 32  BDCR        Backup domain control register.
//  0x94 32  CSR         Clock control & status register.
// Import:
//  stm32/o/l476xx/mmap
package rcc

// DO NOT EDIT THIS FILE. GENERATED BY stm32xgen.

const (
	MSION      CR = 0x01 << 0  //+ Internal Multi Speed oscillator (MSI) clock enable.
	MSIRDY     CR = 0x01 << 1  //+ Internal Multi Speed oscillator (MSI) clock ready flag.
	MSIPLLEN   CR = 0x01 << 2  //+ Internal Multi Speed oscillator (MSI) PLL enable.
	MSIRGSEL   CR = 0x01 << 3  //+ Internal Multi Speed oscillator (MSI) range selection.
	MSIRANGE   CR = 0x0F << 4  //+ Internal Multi Speed oscillator (MSI) clock Range.
	HSION      CR = 0x01 << 8  //+ Internal High Speed oscillator (HSI16) clock enable.
	HSIKERON   CR = 0x01 << 9  //+ Internal High Speed oscillator (HSI16) clock enable for some IPs Kernel.
	HSIRDY     CR = 0x01 << 10 //+ Internal High Speed oscillator (HSI16) clock ready flag.
	HSIASFS    CR = 0x01 << 11 //+ HSI16 Automatic Start from Stop.
	HSEON      CR = 0x01 << 16 //+ External High Speed oscillator (HSE) clock enable.
	HSERDY     CR = 0x01 << 17 //+ External High Speed oscillator (HSE) clock ready.
	HSEBYP     CR = 0x01 << 18 //+ External High Speed oscillator (HSE) clock bypass.
	CSSON      CR = 0x01 << 19 //+ HSE Clock Security System enable.
	PLLON      CR = 0x01 << 24 //+ System PLL clock enable.
	PLLRDY     CR = 0x01 << 25 //+ System PLL clock ready.
	PLLSAI1ON  CR = 0x01 << 26 //+ SAI1 PLL enable.
	PLLSAI1RDY CR = 0x01 << 27 //+ SAI1 PLL ready.
	PLLSAI2ON  CR = 0x01 << 28 //+ SAI2 PLL enable.
	PLLSAI2RDY CR = 0x01 << 29 //+ SAI2 PLL ready.
)

const (
	MSIONn      = 0
	MSIRDYn     = 1
	MSIPLLENn   = 2
	MSIRGSELn   = 3
	MSIRANGEn   = 4
	HSIONn      = 8
	HSIKERONn   = 9
	HSIRDYn     = 10
	HSIASFSn    = 11
	HSEONn      = 16
	HSERDYn     = 17
	HSEBYPn     = 18
	CSSONn      = 19
	PLLONn      = 24
	PLLRDYn     = 25
	PLLSAI1ONn  = 26
	PLLSAI1RDYn = 27
	PLLSAI2ONn  = 28
	PLLSAI2RDYn = 29
)

const (
	MSICAL  ICSCR = 0xFF << 0  //+ MSICAL[7:0] bits.
	MSITRIM ICSCR = 0xFF << 8  //+ MSITRIM[7:0] bits.
	HSICAL  ICSCR = 0xFF << 16 //+ HSICAL[7:0] bits.
	HSITRIM ICSCR = 0x1F << 24 //+ HSITRIM[4:0] bits.
)

const (
	MSICALn  = 0
	MSITRIMn = 8
	HSICALn  = 16
	HSITRIMn = 24
)

const (
	SW           CFGR = 0x03 << 0  //+ SW[1:0] bits (System clock Switch).
	SW_MSI       CFGR = 0x00 << 0  //  MSI oscillator selection as system clock.
	SW_HSI       CFGR = 0x01 << 0  //  HSI16 oscillator selection as system clock.
	SW_HSE       CFGR = 0x02 << 0  //  HSE oscillator selection as system clock.
	SW_PLL       CFGR = 0x03 << 0  //  PLL selection as system clock.
	SWS          CFGR = 0x03 << 2  //+ SWS[1:0] bits (System Clock Switch Status).
	SWS_MSI      CFGR = 0x00 << 2  //  MSI oscillator used as system clock.
	SWS_HSI      CFGR = 0x01 << 2  //  HSI16 oscillator used as system clock.
	SWS_HSE      CFGR = 0x02 << 2  //  HSE oscillator used as system clock.
	SWS_PLL      CFGR = 0x03 << 2  //  PLL used as system clock.
	HPRE         CFGR = 0x0F << 4  //+ HPRE[3:0] bits (AHB prescaler).
	HPRE_DIV1    CFGR = 0x00 << 4  //  SYSCLK not divided.
	HPRE_DIV2    CFGR = 0x08 << 4  //  SYSCLK divided by 2.
	HPRE_DIV4    CFGR = 0x09 << 4  //  SYSCLK divided by 4.
	HPRE_DIV8    CFGR = 0x0A << 4  //  SYSCLK divided by 8.
	HPRE_DIV16   CFGR = 0x0B << 4  //  SYSCLK divided by 16.
	HPRE_DIV64   CFGR = 0x0C << 4  //  SYSCLK divided by 64.
	HPRE_DIV128  CFGR = 0x0D << 4  //  SYSCLK divided by 128.
	HPRE_DIV256  CFGR = 0x0E << 4  //  SYSCLK divided by 256.
	HPRE_DIV512  CFGR = 0x0F << 4  //  SYSCLK divided by 512.
	PPRE1        CFGR = 0x07 << 8  //+ PRE1[2:0] bits (APB2 prescaler).
	PPRE1_DIV1   CFGR = 0x00 << 8  //  HCLK not divided.
	PPRE1_DIV2   CFGR = 0x04 << 8  //  HCLK divided by 2.
	PPRE1_DIV4   CFGR = 0x05 << 8  //  HCLK divided by 4.
	PPRE1_DIV8   CFGR = 0x06 << 8  //  HCLK divided by 8.
	PPRE1_DIV16  CFGR = 0x07 << 8  //  HCLK divided by 16.
	PPRE2        CFGR = 0x07 << 11 //+ PRE2[2:0] bits (APB2 prescaler).
	PPRE2_DIV1   CFGR = 0x00 << 11 //  HCLK not divided.
	PPRE2_DIV2   CFGR = 0x04 << 11 //  HCLK divided by 2.
	PPRE2_DIV4   CFGR = 0x05 << 11 //  HCLK divided by 4.
	PPRE2_DIV8   CFGR = 0x06 << 11 //  HCLK divided by 8.
	PPRE2_DIV16  CFGR = 0x07 << 11 //  HCLK divided by 16.
	STOPWUCK     CFGR = 0x01 << 15 //+ Wake Up from stop and CSS backup clock selection.
	MCOSEL       CFGR = 0x07 << 24 //+ MCOSEL [2:0] bits (Clock output selection).
	MCOPRE       CFGR = 0x07 << 28 //+ MCO prescaler.
	MCOPRE_DIV1  CFGR = 0x00 << 28 //  MCO is divided by 1.
	MCOPRE_DIV2  CFGR = 0x01 << 28 //  MCO is divided by 2.
	MCOPRE_DIV4  CFGR = 0x02 << 28 //  MCO is divided by 4.
	MCOPRE_DIV8  CFGR = 0x03 << 28 //  MCO is divided by 8.
	MCOPRE_DIV16 CFGR = 0x04 << 28 //  MCO is divided by 16.
)

const (
	SWn       = 0
	SWSn      = 2
	HPREn     = 4
	PPRE1n    = 8
	PPRE2n    = 11
	STOPWUCKn = 15
	MCOSELn   = 24
	MCOPREn   = 28
)

const (
	PLLSRC     PLLCFGR = 0x03 << 0  //+
	PLLSRC_MSI PLLCFGR = 0x01 << 0  //  MSI oscillator source clock selected.
	PLLSRC_HSI PLLCFGR = 0x02 << 0  //  HSI16 oscillator source clock selected.
	PLLSRC_HSE PLLCFGR = 0x03 << 0  //  HSE oscillator source clock selected.
	PLLM       PLLCFGR = 0x07 << 4  //+
	PLLN       PLLCFGR = 0x7F << 8  //+
	PLLPEN     PLLCFGR = 0x01 << 16 //+
	PLLP       PLLCFGR = 0x01 << 17 //+
	PLLQEN     PLLCFGR = 0x01 << 20 //+
	PLLQ       PLLCFGR = 0x03 << 21 //+
	PLLREN     PLLCFGR = 0x01 << 24 //+
	PLLR       PLLCFGR = 0x03 << 25 //+
)

const (
	PLLSRCn = 0
	PLLMn   = 4
	PLLNn   = 8
	PLLPENn = 16
	PLLPn   = 17
	PLLQENn = 20
	PLLQn   = 21
	PLLRENn = 24
	PLLRn   = 25
)

const (
	PLLSAI1N   PLLSAI1CFGR = 0x7F << 8  //+
	PLLSAI1PEN PLLSAI1CFGR = 0x01 << 16 //+
	PLLSAI1P   PLLSAI1CFGR = 0x01 << 17 //+
	PLLSAI1QEN PLLSAI1CFGR = 0x01 << 20 //+
	PLLSAI1Q   PLLSAI1CFGR = 0x03 << 21 //+
	PLLSAI1REN PLLSAI1CFGR = 0x01 << 24 //+
	PLLSAI1R   PLLSAI1CFGR = 0x03 << 25 //+
)

const (
	PLLSAI1Nn   = 8
	PLLSAI1PENn = 16
	PLLSAI1Pn   = 17
	PLLSAI1QENn = 20
	PLLSAI1Qn   = 21
	PLLSAI1RENn = 24
	PLLSAI1Rn   = 25
)

const (
	PLLSAI2N   PLLSAI2CFGR = 0x7F << 8  //+
	PLLSAI2PEN PLLSAI2CFGR = 0x01 << 16 //+
	PLLSAI2P   PLLSAI2CFGR = 0x01 << 17 //+
	PLLSAI2REN PLLSAI2CFGR = 0x01 << 24 //+
	PLLSAI2R   PLLSAI2CFGR = 0x03 << 25 //+
)

const (
	PLLSAI2Nn   = 8
	PLLSAI2PENn = 16
	PLLSAI2Pn   = 17
	PLLSAI2RENn = 24
	PLLSAI2Rn   = 25
)

const (
	LSIRDYIE     CIER = 0x01 << 0 //+
	LSERDYIE     CIER = 0x01 << 1 //+
	MSIRDYIE     CIER = 0x01 << 2 //+
	HSIRDYIE     CIER = 0x01 << 3 //+
	HSERDYIE     CIER = 0x01 << 4 //+
	PLLRDYIE     CIER = 0x01 << 5 //+
	PLLSAI1RDYIE CIER = 0x01 << 6 //+
	PLLSAI2RDYIE CIER = 0x01 << 7 //+
	LSECSSIE     CIER = 0x01 << 9 //+
)

const (
	LSIRDYIEn     = 0
	LSERDYIEn     = 1
	MSIRDYIEn     = 2
	HSIRDYIEn     = 3
	HSERDYIEn     = 4
	PLLRDYIEn     = 5
	PLLSAI1RDYIEn = 6
	PLLSAI2RDYIEn = 7
	LSECSSIEn     = 9
)

const (
	LSIRDYF     CIFR = 0x01 << 0 //+
	LSERDYF     CIFR = 0x01 << 1 //+
	MSIRDYF     CIFR = 0x01 << 2 //+
	HSIRDYF     CIFR = 0x01 << 3 //+
	HSERDYF     CIFR = 0x01 << 4 //+
	PLLRDYF     CIFR = 0x01 << 5 //+
	PLLSAI1RDYF CIFR = 0x01 << 6 //+
	PLLSAI2RDYF CIFR = 0x01 << 7 //+
	CSSF        CIFR = 0x01 << 8 //+
	LSECSSF     CIFR = 0x01 << 9 //+
)

const (
	LSIRDYFn     = 0
	LSERDYFn     = 1
	MSIRDYFn     = 2
	HSIRDYFn     = 3
	HSERDYFn     = 4
	PLLRDYFn     = 5
	PLLSAI1RDYFn = 6
	PLLSAI2RDYFn = 7
	CSSFn        = 8
	LSECSSFn     = 9
)

const (
	LSIRDYC     CICR = 0x01 << 0 //+
	LSERDYC     CICR = 0x01 << 1 //+
	MSIRDYC     CICR = 0x01 << 2 //+
	HSIRDYC     CICR = 0x01 << 3 //+
	HSERDYC     CICR = 0x01 << 4 //+
	PLLRDYC     CICR = 0x01 << 5 //+
	PLLSAI1RDYC CICR = 0x01 << 6 //+
	PLLSAI2RDYC CICR = 0x01 << 7 //+
	CSSC        CICR = 0x01 << 8 //+
	LSECSSC     CICR = 0x01 << 9 //+
)

const (
	LSIRDYCn     = 0
	LSERDYCn     = 1
	MSIRDYCn     = 2
	HSIRDYCn     = 3
	HSERDYCn     = 4
	PLLRDYCn     = 5
	PLLSAI1RDYCn = 6
	PLLSAI2RDYCn = 7
	CSSCn        = 8
	LSECSSCn     = 9
)

const (
	DMA1RST  AHB1RSTR = 0x01 << 0  //+
	DMA2RST  AHB1RSTR = 0x01 << 1  //+
	FLASHRST AHB1RSTR = 0x01 << 8  //+
	CRCRST   AHB1RSTR = 0x01 << 12 //+
	TSCRST   AHB1RSTR = 0x01 << 16 //+
)

const (
	DMA1RSTn  = 0
	DMA2RSTn  = 1
	FLASHRSTn = 8
	CRCRSTn   = 12
	TSCRSTn   = 16
)

const (
	GPIOARST AHB2RSTR = 0x01 << 0  //+
	GPIOBRST AHB2RSTR = 0x01 << 1  //+
	GPIOCRST AHB2RSTR = 0x01 << 2  //+
	GPIODRST AHB2RSTR = 0x01 << 3  //+
	GPIOERST AHB2RSTR = 0x01 << 4  //+
	GPIOFRST AHB2RSTR = 0x01 << 5  //+
	GPIOGRST AHB2RSTR = 0x01 << 6  //+
	GPIOHRST AHB2RSTR = 0x01 << 7  //+
	OTGFSRST AHB2RSTR = 0x01 << 12 //+
	ADCRST   AHB2RSTR = 0x01 << 13 //+
	RNGRST   AHB2RSTR = 0x01 << 18 //+
)

const (
	GPIOARSTn = 0
	GPIOBRSTn = 1
	GPIOCRSTn = 2
	GPIODRSTn = 3
	GPIOERSTn = 4
	GPIOFRSTn = 5
	GPIOGRSTn = 6
	GPIOHRSTn = 7
	OTGFSRSTn = 12
	ADCRSTn   = 13
	RNGRSTn   = 18
)

const (
	FMCRST  AHB3RSTR = 0x01 << 0 //+
	QSPIRST AHB3RSTR = 0x01 << 8 //+
)

const (
	FMCRSTn  = 0
	QSPIRSTn = 8
)

const (
	TIM2RST   APB1RSTR1 = 0x01 << 0  //+
	TIM3RST   APB1RSTR1 = 0x01 << 1  //+
	TIM4RST   APB1RSTR1 = 0x01 << 2  //+
	TIM5RST   APB1RSTR1 = 0x01 << 3  //+
	TIM6RST   APB1RSTR1 = 0x01 << 4  //+
	TIM7RST   APB1RSTR1 = 0x01 << 5  //+
	LCDRST    APB1RSTR1 = 0x01 << 9  //+
	SPI2RST   APB1RSTR1 = 0x01 << 14 //+
	SPI3RST   APB1RSTR1 = 0x01 << 15 //+
	USART2RST APB1RSTR1 = 0x01 << 17 //+
	USART3RST APB1RSTR1 = 0x01 << 18 //+
	UART4RST  APB1RSTR1 = 0x01 << 19 //+
	UART5RST  APB1RSTR1 = 0x01 << 20 //+
	I2C1RST   APB1RSTR1 = 0x01 << 21 //+
	I2C2RST   APB1RSTR1 = 0x01 << 22 //+
	I2C3RST   APB1RSTR1 = 0x01 << 23 //+
	CAN1RST   APB1RSTR1 = 0x01 << 25 //+
	PWRRST    APB1RSTR1 = 0x01 << 28 //+
	DAC1RST   APB1RSTR1 = 0x01 << 29 //+
	OPAMPRST  APB1RSTR1 = 0x01 << 30 //+
	LPTIM1RST APB1RSTR1 = 0x01 << 31 //+
)

const (
	TIM2RSTn   = 0
	TIM3RSTn   = 1
	TIM4RSTn   = 2
	TIM5RSTn   = 3
	TIM6RSTn   = 4
	TIM7RSTn   = 5
	LCDRSTn    = 9
	SPI2RSTn   = 14
	SPI3RSTn   = 15
	USART2RSTn = 17
	USART3RSTn = 18
	UART4RSTn  = 19
	UART5RSTn  = 20
	I2C1RSTn   = 21
	I2C2RSTn   = 22
	I2C3RSTn   = 23
	CAN1RSTn   = 25
	PWRRSTn    = 28
	DAC1RSTn   = 29
	OPAMPRSTn  = 30
	LPTIM1RSTn = 31
)

const (
	LPUART1RST APB1RSTR2 = 0x01 << 0 //+
	SWPMI1RST  APB1RSTR2 = 0x01 << 2 //+
	LPTIM2RST  APB1RSTR2 = 0x01 << 5 //+
)

const (
	LPUART1RSTn = 0
	SWPMI1RSTn  = 2
	LPTIM2RSTn  = 5
)

const (
	SYSCFGRST APB2RSTR = 0x01 << 0  //+
	SDMMC1RST APB2RSTR = 0x01 << 10 //+
	TIM1RST   APB2RSTR = 0x01 << 11 //+
	SPI1RST   APB2RSTR = 0x01 << 12 //+
	TIM8RST   APB2RSTR = 0x01 << 13 //+
	USART1RST APB2RSTR = 0x01 << 14 //+
	TIM15RST  APB2RSTR = 0x01 << 16 //+
	TIM16RST  APB2RSTR = 0x01 << 17 //+
	TIM17RST  APB2RSTR = 0x01 << 18 //+
	SAI1RST   APB2RSTR = 0x01 << 21 //+
	SAI2RST   APB2RSTR = 0x01 << 22 //+
	DFSDM1RST APB2RSTR = 0x01 << 24 //+
)

const (
	SYSCFGRSTn = 0
	SDMMC1RSTn = 10
	TIM1RSTn   = 11
	SPI1RSTn   = 12
	TIM8RSTn   = 13
	USART1RSTn = 14
	TIM15RSTn  = 16
	TIM16RSTn  = 17
	TIM17RSTn  = 18
	SAI1RSTn   = 21
	SAI2RSTn   = 22
	DFSDM1RSTn = 24
)

const (
	DMA1EN  AHB1ENR = 0x01 << 0  //+
	DMA2EN  AHB1ENR = 0x01 << 1  //+
	FLASHEN AHB1ENR = 0x01 << 8  //+
	CRCEN   AHB1ENR = 0x01 << 12 //+
	TSCEN   AHB1ENR = 0x01 << 16 //+
)

const (
	DMA1ENn  = 0
	DMA2ENn  = 1
	FLASHENn = 8
	CRCENn   = 12
	TSCENn   = 16
)

const (
	GPIOAEN AHB2ENR = 0x01 << 0  //+
	GPIOBEN AHB2ENR = 0x01 << 1  //+
	GPIOCEN AHB2ENR = 0x01 << 2  //+
	GPIODEN AHB2ENR = 0x01 << 3  //+
	GPIOEEN AHB2ENR = 0x01 << 4  //+
	GPIOFEN AHB2ENR = 0x01 << 5  //+
	GPIOGEN AHB2ENR = 0x01 << 6  //+
	GPIOHEN AHB2ENR = 0x01 << 7  //+
	OTGFSEN AHB2ENR = 0x01 << 12 //+
	ADCEN   AHB2ENR = 0x01 << 13 //+
	RNGEN   AHB2ENR = 0x01 << 18 //+
)

const (
	GPIOAENn = 0
	GPIOBENn = 1
	GPIOCENn = 2
	GPIODENn = 3
	GPIOEENn = 4
	GPIOFENn = 5
	GPIOGENn = 6
	GPIOHENn = 7
	OTGFSENn = 12
	ADCENn   = 13
	RNGENn   = 18
)

const (
	FMCEN  AHB3ENR = 0x01 << 0 //+
	QSPIEN AHB3ENR = 0x01 << 8 //+
)

const (
	FMCENn  = 0
	QSPIENn = 8
)

const (
	TIM2EN   APB1ENR1 = 0x01 << 0  //+
	TIM3EN   APB1ENR1 = 0x01 << 1  //+
	TIM4EN   APB1ENR1 = 0x01 << 2  //+
	TIM5EN   APB1ENR1 = 0x01 << 3  //+
	TIM6EN   APB1ENR1 = 0x01 << 4  //+
	TIM7EN   APB1ENR1 = 0x01 << 5  //+
	LCDEN    APB1ENR1 = 0x01 << 9  //+
	WWDGEN   APB1ENR1 = 0x01 << 11 //+
	SPI2EN   APB1ENR1 = 0x01 << 14 //+
	SPI3EN   APB1ENR1 = 0x01 << 15 //+
	USART2EN APB1ENR1 = 0x01 << 17 //+
	USART3EN APB1ENR1 = 0x01 << 18 //+
	UART4EN  APB1ENR1 = 0x01 << 19 //+
	UART5EN  APB1ENR1 = 0x01 << 20 //+
	I2C1EN   APB1ENR1 = 0x01 << 21 //+
	I2C2EN   APB1ENR1 = 0x01 << 22 //+
	I2C3EN   APB1ENR1 = 0x01 << 23 //+
	CAN1EN   APB1ENR1 = 0x01 << 25 //+
	PWREN    APB1ENR1 = 0x01 << 28 //+
	DAC1EN   APB1ENR1 = 0x01 << 29 //+
	OPAMPEN  APB1ENR1 = 0x01 << 30 //+
	LPTIM1EN APB1ENR1 = 0x01 << 31 //+
)

const (
	TIM2ENn   = 0
	TIM3ENn   = 1
	TIM4ENn   = 2
	TIM5ENn   = 3
	TIM6ENn   = 4
	TIM7ENn   = 5
	LCDENn    = 9
	WWDGENn   = 11
	SPI2ENn   = 14
	SPI3ENn   = 15
	USART2ENn = 17
	USART3ENn = 18
	UART4ENn  = 19
	UART5ENn  = 20
	I2C1ENn   = 21
	I2C2ENn   = 22
	I2C3ENn   = 23
	CAN1ENn   = 25
	PWRENn    = 28
	DAC1ENn   = 29
	OPAMPENn  = 30
	LPTIM1ENn = 31
)

const (
	LPUART1EN APB1ENR2 = 0x01 << 0 //+
	SWPMI1EN  APB1ENR2 = 0x01 << 2 //+
	LPTIM2EN  APB1ENR2 = 0x01 << 5 //+
)

const (
	LPUART1ENn = 0
	SWPMI1ENn  = 2
	LPTIM2ENn  = 5
)

const (
	SYSCFGEN APB2ENR = 0x01 << 0  //+
	FWEN     APB2ENR = 0x01 << 7  //+
	SDMMC1EN APB2ENR = 0x01 << 10 //+
	TIM1EN   APB2ENR = 0x01 << 11 //+
	SPI1EN   APB2ENR = 0x01 << 12 //+
	TIM8EN   APB2ENR = 0x01 << 13 //+
	USART1EN APB2ENR = 0x01 << 14 //+
	TIM15EN  APB2ENR = 0x01 << 16 //+
	TIM16EN  APB2ENR = 0x01 << 17 //+
	TIM17EN  APB2ENR = 0x01 << 18 //+
	SAI1EN   APB2ENR = 0x01 << 21 //+
	SAI2EN   APB2ENR = 0x01 << 22 //+
	DFSDM1EN APB2ENR = 0x01 << 24 //+
)

const (
	SYSCFGENn = 0
	FWENn     = 7
	SDMMC1ENn = 10
	TIM1ENn   = 11
	SPI1ENn   = 12
	TIM8ENn   = 13
	USART1ENn = 14
	TIM15ENn  = 16
	TIM16ENn  = 17
	TIM17ENn  = 18
	SAI1ENn   = 21
	SAI2ENn   = 22
	DFSDM1ENn = 24
)

const (
	DMA1SMEN  AHB1SMENR = 0x01 << 0  //+
	DMA2SMEN  AHB1SMENR = 0x01 << 1  //+
	FLASHSMEN AHB1SMENR = 0x01 << 8  //+
	SRAM1SMEN AHB1SMENR = 0x01 << 9  //+
	CRCSMEN   AHB1SMENR = 0x01 << 12 //+
	TSCSMEN   AHB1SMENR = 0x01 << 16 //+
)

const (
	DMA1SMENn  = 0
	DMA2SMENn  = 1
	FLASHSMENn = 8
	SRAM1SMENn = 9
	CRCSMENn   = 12
	TSCSMENn   = 16
)

const (
	GPIOASMEN AHB2SMENR = 0x01 << 0  //+
	GPIOBSMEN AHB2SMENR = 0x01 << 1  //+
	GPIOCSMEN AHB2SMENR = 0x01 << 2  //+
	GPIODSMEN AHB2SMENR = 0x01 << 3  //+
	GPIOESMEN AHB2SMENR = 0x01 << 4  //+
	GPIOFSMEN AHB2SMENR = 0x01 << 5  //+
	GPIOGSMEN AHB2SMENR = 0x01 << 6  //+
	GPIOHSMEN AHB2SMENR = 0x01 << 7  //+
	SRAM2SMEN AHB2SMENR = 0x01 << 9  //+
	OTGFSSMEN AHB2SMENR = 0x01 << 12 //+
	ADCSMEN   AHB2SMENR = 0x01 << 13 //+
	RNGSMEN   AHB2SMENR = 0x01 << 18 //+
)

const (
	GPIOASMENn = 0
	GPIOBSMENn = 1
	GPIOCSMENn = 2
	GPIODSMENn = 3
	GPIOESMENn = 4
	GPIOFSMENn = 5
	GPIOGSMENn = 6
	GPIOHSMENn = 7
	SRAM2SMENn = 9
	OTGFSSMENn = 12
	ADCSMENn   = 13
	RNGSMENn   = 18
)

const (
	FMCSMEN  AHB3SMENR = 0x01 << 0 //+
	QSPISMEN AHB3SMENR = 0x01 << 8 //+
)

const (
	FMCSMENn  = 0
	QSPISMENn = 8
)

const (
	TIM2SMEN   APB1SMENR1 = 0x01 << 0  //+
	TIM3SMEN   APB1SMENR1 = 0x01 << 1  //+
	TIM4SMEN   APB1SMENR1 = 0x01 << 2  //+
	TIM5SMEN   APB1SMENR1 = 0x01 << 3  //+
	TIM6SMEN   APB1SMENR1 = 0x01 << 4  //+
	TIM7SMEN   APB1SMENR1 = 0x01 << 5  //+
	LCDSMEN    APB1SMENR1 = 0x01 << 9  //+
	WWDGSMEN   APB1SMENR1 = 0x01 << 11 //+
	SPI2SMEN   APB1SMENR1 = 0x01 << 14 //+
	SPI3SMEN   APB1SMENR1 = 0x01 << 15 //+
	USART2SMEN APB1SMENR1 = 0x01 << 17 //+
	USART3SMEN APB1SMENR1 = 0x01 << 18 //+
	UART4SMEN  APB1SMENR1 = 0x01 << 19 //+
	UART5SMEN  APB1SMENR1 = 0x01 << 20 //+
	I2C1SMEN   APB1SMENR1 = 0x01 << 21 //+
	I2C2SMEN   APB1SMENR1 = 0x01 << 22 //+
	I2C3SMEN   APB1SMENR1 = 0x01 << 23 //+
	CAN1SMEN   APB1SMENR1 = 0x01 << 25 //+
	PWRSMEN    APB1SMENR1 = 0x01 << 28 //+
	DAC1SMEN   APB1SMENR1 = 0x01 << 29 //+
	OPAMPSMEN  APB1SMENR1 = 0x01 << 30 //+
	LPTIM1SMEN APB1SMENR1 = 0x01 << 31 //+
)

const (
	TIM2SMENn   = 0
	TIM3SMENn   = 1
	TIM4SMENn   = 2
	TIM5SMENn   = 3
	TIM6SMENn   = 4
	TIM7SMENn   = 5
	LCDSMENn    = 9
	WWDGSMENn   = 11
	SPI2SMENn   = 14
	SPI3SMENn   = 15
	USART2SMENn = 17
	USART3SMENn = 18
	UART4SMENn  = 19
	UART5SMENn  = 20
	I2C1SMENn   = 21
	I2C2SMENn   = 22
	I2C3SMENn   = 23
	CAN1SMENn   = 25
	PWRSMENn    = 28
	DAC1SMENn   = 29
	OPAMPSMENn  = 30
	LPTIM1SMENn = 31
)

const (
	LPUART1SMEN APB1SMENR2 = 0x01 << 0 //+
	SWPMI1SMEN  APB1SMENR2 = 0x01 << 2 //+
	LPTIM2SMEN  APB1SMENR2 = 0x01 << 5 //+
)

const (
	LPUART1SMENn = 0
	SWPMI1SMENn  = 2
	LPTIM2SMENn  = 5
)

const (
	SYSCFGSMEN APB2SMENR = 0x01 << 0  //+
	SDMMC1SMEN APB2SMENR = 0x01 << 10 //+
	TIM1SMEN   APB2SMENR = 0x01 << 11 //+
	SPI1SMEN   APB2SMENR = 0x01 << 12 //+
	TIM8SMEN   APB2SMENR = 0x01 << 13 //+
	USART1SMEN APB2SMENR = 0x01 << 14 //+
	TIM15SMEN  APB2SMENR = 0x01 << 16 //+
	TIM16SMEN  APB2SMENR = 0x01 << 17 //+
	TIM17SMEN  APB2SMENR = 0x01 << 18 //+
	SAI1SMEN   APB2SMENR = 0x01 << 21 //+
	SAI2SMEN   APB2SMENR = 0x01 << 22 //+
	DFSDM1SMEN APB2SMENR = 0x01 << 24 //+
)

const (
	SYSCFGSMENn = 0
	SDMMC1SMENn = 10
	TIM1SMENn   = 11
	SPI1SMENn   = 12
	TIM8SMENn   = 13
	USART1SMENn = 14
	TIM15SMENn  = 16
	TIM16SMENn  = 17
	TIM17SMENn  = 18
	SAI1SMENn   = 21
	SAI2SMENn   = 22
	DFSDM1SMENn = 24
)

const (
	USART1SEL  CCIPR = 0x03 << 0  //+
	USART2SEL  CCIPR = 0x03 << 2  //+
	USART3SEL  CCIPR = 0x03 << 4  //+
	UART4SEL   CCIPR = 0x03 << 6  //+
	UART5SEL   CCIPR = 0x03 << 8  //+
	LPUART1SEL CCIPR = 0x03 << 10 //+
	I2C1SEL    CCIPR = 0x03 << 12 //+
	I2C2SEL    CCIPR = 0x03 << 14 //+
	I2C3SEL    CCIPR = 0x03 << 16 //+
	LPTIM1SEL  CCIPR = 0x03 << 18 //+
	LPTIM2SEL  CCIPR = 0x03 << 20 //+
	SAI1SEL    CCIPR = 0x03 << 22 //+
	SAI2SEL    CCIPR = 0x03 << 24 //+
	CLK48SEL   CCIPR = 0x03 << 26 //+
	ADCSEL     CCIPR = 0x03 << 28 //+
	SWPMI1SEL  CCIPR = 0x01 << 30 //+
	DFSDM1SEL  CCIPR = 0x01 << 31 //+
)

const (
	USART1SELn  = 0
	USART2SELn  = 2
	USART3SELn  = 4
	UART4SELn   = 6
	UART5SELn   = 8
	LPUART1SELn = 10
	I2C1SELn    = 12
	I2C2SELn    = 14
	I2C3SELn    = 16
	LPTIM1SELn  = 18
	LPTIM2SELn  = 20
	SAI1SELn    = 22
	SAI2SELn    = 24
	CLK48SELn   = 26
	ADCSELn     = 28
	SWPMI1SELn  = 30
	DFSDM1SELn  = 31
)

const (
	LSEON    BDCR = 0x01 << 0  //+
	LSERDY   BDCR = 0x01 << 1  //+
	LSEBYP   BDCR = 0x01 << 2  //+
	LSEDRV   BDCR = 0x03 << 3  //+
	LSECSSON BDCR = 0x01 << 5  //+
	LSECSSD  BDCR = 0x01 << 6  //+
	RTCSEL   BDCR = 0x03 << 8  //+
	RTCEN    BDCR = 0x01 << 15 //+
	BDRST    BDCR = 0x01 << 16 //+
	LSCOEN   BDCR = 0x01 << 24 //+
	LSCOSEL  BDCR = 0x01 << 25 //+
)

const (
	LSEONn    = 0
	LSERDYn   = 1
	LSEBYPn   = 2
	LSEDRVn   = 3
	LSECSSONn = 5
	LSECSSDn  = 6
	RTCSELn   = 8
	RTCENn    = 15
	BDRSTn    = 16
	LSCOENn   = 24
	LSCOSELn  = 25
)

const (
	LSION     CSR = 0x01 << 0  //+
	LSIRDY    CSR = 0x01 << 1  //+
	MSISRANGE CSR = 0x0F << 8  //+
	RMVF      CSR = 0x01 << 23 //+
	FWRSTF    CSR = 0x01 << 24 //+
	OBLRSTF   CSR = 0x01 << 25 //+
	PINRSTF   CSR = 0x01 << 26 //+
	BORRSTF   CSR = 0x01 << 27 //+
	SFTRSTF   CSR = 0x01 << 28 //+
	IWDGRSTF  CSR = 0x01 << 29 //+
	WWDGRSTF  CSR = 0x01 << 30 //+
	LPWRRSTF  CSR = 0x01 << 31 //+
)

const (
	LSIONn     = 0
	LSIRDYn    = 1
	MSISRANGEn = 8
	RMVFn      = 23
	FWRSTFn    = 24
	OBLRSTFn   = 25
	PINRSTFn   = 26
	BORRSTFn   = 27
	SFTRSTFn   = 28
	IWDGRSTFn  = 29
	WWDGRSTFn  = 30
	LPWRRSTFn  = 31
)
