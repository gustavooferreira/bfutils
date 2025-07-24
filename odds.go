// Package bfutils provides a set of utility functions that help with day to day automation in the Betfair exchange.
package bfutils

import (
	"fmt"
)

// Odd represents an odd.
type Odd struct {
	// Index is the index of the odd in the ladder, starting from 0.
	Index int
	// Fixed is the odd represented as an FixedOdd (integer).
	Fixed FixedOdd
	// Float is the odd represented as a float64.
	Float float64
	// String is the odd represented as a string, e.g., "1.01", "2.5", etc.
	String string
}

// OddFromIndex returns an Odd from the given index.
func OddFromIndex(index int) (Odd, error) {
	if index < 0 || index >= OddsCount {
		return Odd{}, fmt.Errorf("index out of bounds: %d", index)
	}

	return Odd{
		Index:  index,
		Fixed:  OddsFixed[index],
		Float:  Odds[index],
		String: OddsStr[index],
	}, nil
}

// OddFromFixedOdd returns an Odd from the given odd string.
func OddFromString(odd string) (Odd, error) {
	index, ok := FindStrOdd(odd)
	if !ok {
		return Odd{}, fmt.Errorf("string odd does not represent a valid odd")
	}

	return OddFromIndex(index)
}

// OddFromFixed returns an Odd from the given float odd.
func OddFromFloat(odd float64) (Odd, error) {
	match, index, err := FindOdd(odd)
	if err != nil {
		return Odd{}, err
	}

	if !match {
		return Odd{}, fmt.Errorf("odd not valid")
	}

	return OddFromIndex(index)
}

// OddFromFixedOdd returns an Odd from the given fixed odd.
func OddFromFixedOdd(odd FixedOdd) (Odd, error) {
	index, ok := FindFixedOdd(odd)
	if !ok {
		return Odd{}, fmt.Errorf("fixed odd does not represent a valid odd")
	}

	return OddFromIndex(index)
}

// OddRange represents a range of odds.
type OddRange struct {
	// Begin is the first odd in the range of odds.
	Begin Odd
	// End is the last odd in the range of odds.
	End Odd
	// Increment is the increment for the odd in this range, e.g., "0.01", "0.02", "0.05", etc.
	Increment OddIncrement
	// TicksCount is the number of ticks in the range.
	TicksCount int
}

// OddIncrement represents the increment of odds in a range.
// The increments are stored as integers scaled by a factor of 100. An increment of 0.02 in the range 2.02 to 3 is
// represented as the integer 2.
type OddIncrement int

// FixedOdd is an integer type that represents odds as fixed-point numbers.
// Odds are stored as integers scaled by a factor of 100, e.g., 1.01 is represented as 101 integer, and 50.0 is
// represented as 5000 integer.
type FixedOdd int

// ToOdd converts a FixedOdd to an Odd without looking it up to check if it's valid.
func (o FixedOdd) ToOdd() Odd {
	floatOdd := ConvertFixedToFloatOdd(o)
	return Odd{
		Index:  -1,
		Fixed:  o,
		Float:  floatOdd,
		String: fmt.Sprintf("%.2f", floatOdd),
	}
}

// ToOddWithLookup converts a FixedOdd to an Odd, looking it up to check if it's valid.
func (o FixedOdd) ToOddWithLookup() (Odd, error) {
	index, ok := FindFixedOdd(o)
	if !ok {
		return Odd{}, fmt.Errorf("fixed odd does not represent a valid odd")
	}

	return Odd{
		Index:  index,
		Fixed:  o,
		Float:  Odds[index],
		String: OddsStr[index],
	}, nil
}

// ConvertFixedToFloat converts a FixedOdd to an odd represented as float64.
func ConvertFixedToFloatOdd(odd FixedOdd) float64 {
	return float64(odd) / 100.0
}

// OddsCount is a constant defining the number of available odds in the ladder.
const OddsCount = 350

// OddsFixed is an array of FixedOdd holding the value of all tradable odds.
var OddsFixed = [...]FixedOdd{
	101, 102, 103, 104, 105, 106, 107, 108, 109, 110,
	111, 112, 113, 114, 115, 116, 117, 118, 119, 120,
	121, 122, 123, 124, 125, 126, 127, 128, 129, 130,
	131, 132, 133, 134, 135, 136, 137, 138, 139, 140,
	141, 142, 143, 144, 145, 146, 147, 148, 149, 150,
	151, 152, 153, 154, 155, 156, 157, 158, 159, 160,
	161, 162, 163, 164, 165, 166, 167, 168, 169, 170,
	171, 172, 173, 174, 175, 176, 177, 178, 179, 180,
	181, 182, 183, 184, 185, 186, 187, 188, 189, 190,
	191, 192, 193, 194, 195, 196, 197, 198, 199, 200,
	202, 204, 206, 208, 210, 212, 214, 216, 218, 220,
	222, 224, 226, 228, 230, 232, 234, 236, 238, 240,
	242, 244, 246, 248, 250, 252, 254, 256, 258, 260,
	262, 264, 266, 268, 270, 272, 274, 276, 278, 280,
	282, 284, 286, 288, 290, 292, 294, 296, 298, 300,
	305, 310, 315, 320, 325, 330, 335, 340, 345, 350,
	355, 360, 365, 370, 375, 380, 385, 390, 395, 400,
	410, 420, 430, 440, 450, 460, 470, 480, 490, 500,
	510, 520, 530, 540, 550, 560, 570, 580, 590, 600,
	620, 640, 660, 680, 700, 720, 740, 760, 780, 800,
	820, 840, 860, 880, 900, 920, 940, 960, 980, 1000,
	1050, 1100, 1150, 1200, 1250, 1300, 1350, 1400, 1450, 1500,
	1550, 1600, 1650, 1700, 1750, 1800, 1850, 1900, 1950, 2000,
	2100, 2200, 2300, 2400, 2500, 2600, 2700, 2800, 2900, 3000,
	3200, 3400, 3600, 3800, 4000, 4200, 4400, 4600, 4800, 5000,
	5500, 6000, 6500, 7000, 7500, 8000, 8500, 9000, 9500, 10000,
	11000, 12000, 13000, 14000, 15000, 16000, 17000, 18000, 19000, 20000,
	21000, 22000, 23000, 24000, 25000, 26000, 27000, 28000, 29000, 30000,
	31000, 32000, 33000, 34000, 35000, 36000, 37000, 38000, 39000, 40000,
	41000, 42000, 43000, 44000, 45000, 46000, 47000, 48000, 49000, 50000,
	51000, 52000, 53000, 54000, 55000, 56000, 57000, 58000, 59000, 60000,
	61000, 62000, 63000, 64000, 65000, 66000, 67000, 68000, 69000, 70000,
	71000, 72000, 73000, 74000, 75000, 76000, 77000, 78000, 79000, 80000,
	81000, 82000, 83000, 84000, 85000, 86000, 87000, 88000, 89000, 90000,
	91000, 92000, 93000, 94000, 95000, 96000, 97000, 98000, 99000, 100000,
}

// Odds is an array of float64 holding the value of all tradable odds.
var Odds = [...]float64{
	1.01, 1.02, 1.03, 1.04, 1.05, 1.06, 1.07, 1.08, 1.09, 1.1,
	1.11, 1.12, 1.13, 1.14, 1.15, 1.16, 1.17, 1.18, 1.19, 1.2,
	1.21, 1.22, 1.23, 1.24, 1.25, 1.26, 1.27, 1.28, 1.29, 1.3,
	1.31, 1.32, 1.33, 1.34, 1.35, 1.36, 1.37, 1.38, 1.39, 1.4,
	1.41, 1.42, 1.43, 1.44, 1.45, 1.46, 1.47, 1.48, 1.49, 1.5,
	1.51, 1.52, 1.53, 1.54, 1.55, 1.56, 1.57, 1.58, 1.59, 1.6,
	1.61, 1.62, 1.63, 1.64, 1.65, 1.66, 1.67, 1.68, 1.69, 1.7,
	1.71, 1.72, 1.73, 1.74, 1.75, 1.76, 1.77, 1.78, 1.79, 1.8,
	1.81, 1.82, 1.83, 1.84, 1.85, 1.86, 1.87, 1.88, 1.89, 1.9,
	1.91, 1.92, 1.93, 1.94, 1.95, 1.96, 1.97, 1.98, 1.99, 2,
	2.02, 2.04, 2.06, 2.08, 2.1, 2.12, 2.14, 2.16, 2.18, 2.2,
	2.22, 2.24, 2.26, 2.28, 2.3, 2.32, 2.34, 2.36, 2.38, 2.4,
	2.42, 2.44, 2.46, 2.48, 2.5, 2.52, 2.54, 2.56, 2.58, 2.6,
	2.62, 2.64, 2.66, 2.68, 2.7, 2.72, 2.74, 2.76, 2.78, 2.8,
	2.82, 2.84, 2.86, 2.88, 2.9, 2.92, 2.94, 2.96, 2.98, 3,
	3.05, 3.1, 3.15, 3.2, 3.25, 3.3, 3.35, 3.4, 3.45, 3.5,
	3.55, 3.6, 3.65, 3.7, 3.75, 3.8, 3.85, 3.9, 3.95, 4,
	4.1, 4.2, 4.3, 4.4, 4.5, 4.6, 4.7, 4.8, 4.9, 5, 5.1, 5.2,
	5.3, 5.4, 5.5, 5.6, 5.7, 5.8, 5.9, 6, 6.2, 6.4, 6.6, 6.8,
	7, 7.2, 7.4, 7.6, 7.8, 8, 8.2, 8.4, 8.6, 8.8, 9, 9.2,
	9.4, 9.6, 9.8, 10, 10.5, 11, 11.5, 12, 12.5, 13, 13.5, 14,
	14.5, 15, 15.5, 16, 16.5, 17, 17.5, 18, 18.5, 19, 19.5, 20,
	21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 32, 34, 36, 38, 40,
	42, 44, 46, 48, 50, 55, 60, 65, 70, 75, 80, 85, 90, 95, 100,
	110, 120, 130, 140, 150, 160, 170, 180, 190, 200, 210, 220,
	230, 240, 250, 260, 270, 280, 290, 300, 310, 320, 330, 340,
	350, 360, 370, 380, 390, 400, 410, 420, 430, 440, 450, 460,
	470, 480, 490, 500, 510, 520, 530, 540, 550, 560, 570, 580,
	590, 600, 610, 620, 630, 640, 650, 660, 670, 680, 690, 700,
	710, 720, 730, 740, 750, 760, 770, 780, 790, 800, 810, 820,
	830, 840, 850, 860, 870, 880, 890, 900, 910, 920, 930, 940,
	950, 960, 970, 980, 990, 1000,
}

// OddsStr is an array of strings holding the value of all tradable odds.
var OddsStr = [...]string{
	"1.01", "1.02", "1.03", "1.04", "1.05", "1.06", "1.07", "1.08", "1.09", "1.1",
	"1.11", "1.12", "1.13", "1.14", "1.15", "1.16", "1.17", "1.18", "1.19", "1.2",
	"1.21", "1.22", "1.23", "1.24", "1.25", "1.26", "1.27", "1.28", "1.29", "1.3",
	"1.31", "1.32", "1.33", "1.34", "1.35", "1.36", "1.37", "1.38", "1.39", "1.4",
	"1.41", "1.42", "1.43", "1.44", "1.45", "1.46", "1.47", "1.48", "1.49", "1.5",
	"1.51", "1.52", "1.53", "1.54", "1.55", "1.56", "1.57", "1.58", "1.59", "1.6",
	"1.61", "1.62", "1.63", "1.64", "1.65", "1.66", "1.67", "1.68", "1.69", "1.7",
	"1.71", "1.72", "1.73", "1.74", "1.75", "1.76", "1.77", "1.78", "1.79", "1.8",
	"1.81", "1.82", "1.83", "1.84", "1.85", "1.86", "1.87", "1.88", "1.89", "1.9",
	"1.91", "1.92", "1.93", "1.94", "1.95", "1.96", "1.97", "1.98", "1.99", "2",
	"2.02", "2.04", "2.06", "2.08", "2.1", "2.12", "2.14", "2.16", "2.18", "2.2",
	"2.22", "2.24", "2.26", "2.28", "2.3", "2.32", "2.34", "2.36", "2.38", "2.4",
	"2.42", "2.44", "2.46", "2.48", "2.5", "2.52", "2.54", "2.56", "2.58", "2.6",
	"2.62", "2.64", "2.66", "2.68", "2.7", "2.72", "2.74", "2.76", "2.78", "2.8",
	"2.82", "2.84", "2.86", "2.88", "2.9", "2.92", "2.94", "2.96", "2.98", "3",
	"3.05", "3.1", "3.15", "3.2", "3.25", "3.3", "3.35", "3.4", "3.45", "3.5",
	"3.55", "3.6", "3.65", "3.7", "3.75", "3.8", "3.85", "3.9", "3.95", "4",
	"4.1", "4.2", "4.3", "4.4", "4.5", "4.6", "4.7", "4.8", "4.9", "5", "5.1", "5.2",
	"5.3", "5.4", "5.5", "5.6", "5.7", "5.8", "5.9", "6", "6.2", "6.4", "6.6", "6.8",
	"7", "7.2", "7.4", "7.6", "7.8", "8", "8.2", "8.4", "8.6", "8.8", "9", "9.2",
	"9.4", "9.6", "9.8", "10", "10.5", "11", "11.5", "12", "12.5", "13", "13.5", "14",
	"14.5", "15", "15.5", "16", "16.5", "17", "17.5", "18", "18.5", "19", "19.5", "20",
	"21", "22", "23", "24", "25", "26", "27", "28", "29", "30", "32", "34", "36", "38", "40",
	"42", "44", "46", "48", "50", "55", "60", "65", "70", "75", "80", "85", "90", "95", "100",
	"110", "120", "130", "140", "150", "160", "170", "180", "190", "200", "210", "220",
	"230", "240", "250", "260", "270", "280", "290", "300", "310", "320", "330", "340",
	"350", "360", "370", "380", "390", "400", "410", "420", "430", "440", "450", "460",
	"470", "480", "490", "500", "510", "520", "530", "540", "550", "560", "570", "580",
	"590", "600", "610", "620", "630", "640", "650", "660", "670", "680", "690", "700",
	"710", "720", "730", "740", "750", "760", "770", "780", "790", "800", "810", "820",
	"830", "840", "850", "860", "870", "880", "890", "900", "910", "920", "930", "940",
	"950", "960", "970", "980", "990", "1000",
}

// OddsRange is a slice with all the odd ranges available.
var OddsRange = []OddRange{
	{
		Begin:      Odd{Fixed: 101, Float: 1.01, String: "1.01"},
		End:        Odd{Fixed: 200, Float: 2.0, String: "2"},
		Increment:  1, // 0.01
		TicksCount: 100,
	},
	{
		Begin:      Odd{Fixed: 202, Float: 2.02, String: "2.02"},
		End:        Odd{Fixed: 300, Float: 3.0, String: "3"},
		Increment:  2, // 0.02
		TicksCount: 50,
	},
	{
		Begin:      Odd{Fixed: 305, Float: 3.05, String: "3.05"},
		End:        Odd{Fixed: 400, Float: 4.0, String: "4"},
		Increment:  5, // 0.05
		TicksCount: 20,
	},
	{
		Begin:      Odd{Fixed: 410, Float: 4.10, String: "4.1"},
		End:        Odd{Fixed: 600, Float: 6.0, String: "6"},
		Increment:  10, // 0.1
		TicksCount: 20,
	},
	{
		Begin:      Odd{Fixed: 620, Float: 6.20, String: "6.2"},
		End:        Odd{Fixed: 1000, Float: 10.0, String: "10"},
		Increment:  20, // 0.2
		TicksCount: 20,
	},
	{
		Begin:      Odd{Fixed: 1050, Float: 10.50, String: "10.5"},
		End:        Odd{Fixed: 2000, Float: 20.0, String: "20"},
		Increment:  50, // 0.5
		TicksCount: 20,
	},
	{
		Begin:      Odd{Fixed: 2100, Float: 21.0, String: "21"},
		End:        Odd{Fixed: 3000, Float: 30.0, String: "30"},
		Increment:  100, // 1
		TicksCount: 10,
	},
	{
		Begin:      Odd{Fixed: 3200, Float: 32.0, String: "32"},
		End:        Odd{Fixed: 5000, Float: 50.0, String: "50"},
		Increment:  200, // 2
		TicksCount: 10,
	},
	{
		Begin:      Odd{Fixed: 5500, Float: 55.0, String: "55"},
		End:        Odd{Fixed: 10000, Float: 100.0, String: "100"},
		Increment:  500, // 5
		TicksCount: 10,
	},
	{
		Begin:      Odd{Fixed: 11000, Float: 110.0, String: "110"},
		End:        Odd{Fixed: 100000, Float: 1000.0, String: "1000"},
		Increment:  1000, // 10
		TicksCount: 90,
	},
}
