package nature

// IRSignal represents IR signal.
type IRSignal struct {
	Freq   Freq   `json:"freq"`
	Data   Data   `json:"data"`
	Format Format `json:"format"`
}

// Freq represents IR sub carrier frequency.
type Freq int

// Data for IR signal. IR signal consists of on/off of sub carrier frequency.
// When receiving IR, Remo measures on to off, off to on time intervals and
// records the time interval sequence. When sending IR, Remo turns on/off the
// sub carrier frequency with the provided time interval sequence.
type Data []int

// Format represents format and unit of values in the data array.
// Fixed to "us", which means each integer of data array is in microseconds.
type Format string
