//go:build sam && atsamd21 && gemma_m0

package machine

// Used to reset into bootloader.
const resetMagicValue = 0xf01669ef

// GPIO Pins.
const (
	D0  = PA04 // SERCOM0/PAD[0]
	D1  = PA02
	D2  = PA05 // SERCOM0/PAD[1]
	D3  = PA00 // DotStar LED: SERCOM1/PAD[0]
	D4  = PA01 // DotStar LED: SERCOM1/PAD[1]
	D13 = PA23 // LED: SERCOM3/PAD[1] SERCOM5/PAD[1]
)

// Analog pins.
const (
	A0 = D1
	A1 = D2
	A2 = D0
)

const (
	LED = D13
)

// USBCDC pins.
const (
	USBCDC_DM_PIN = PA24
	USBCDC_DP_PIN = PA25
)

// UART0 pins.
const (
	UART_TX_PIN = D0 // TX: SERCOM0/PAD[0]
	UART_RX_PIN = D2 // RX: SERCOM0/PAD[1]
)

// UART0s on the Trinket M0.
var UART0 = &sercomUSART0

// SPI pins 
const (
	SPI0_SDI_PIN = D0 // SDI: SERCOM0/PAD[0]
	SPI0_SCK_PIN = D2 // SCK: SERCOM0/PAD[1]
)

// SPI on the Gemma M0.
var SPI0 = sercomSPIM0

// SPI pins for DotStar LED.
const (
	SPI1_SDI_PIN = D3 // SDI: SERCOM1/PAD[0]
	SPI1_SCK_PIN = D4 // SCK: SERCOM1/PAD[1]
)

// SPI for DotStar LED.
var SPI1 = sercomSPIM1

// I2C pins
const (
	SDA_PIN = D0 // SDA: SERCOM0/PAD[0]
	SCL_PIN = D2 // SCL: SERCOM0/PAD[1]
)

// I2C on the Gemma M0.
var (
	I2C0 = sercomI2CM0
)

// USB CDC identifiers.
const (
	usb_STRING_PRODUCT      = "Adafruit Gemma M0"
	usb_STRING_MANUFACTURER = "Adafruit"
)

var (
	usb_VID uint16 = 0x239A
	usb_PID uint16 = 0x801E
)

var (
	DefaultUART = UART0
)
