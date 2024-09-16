package network

import (
	"errors"
	"strconv"
)

type Octet struct {
	octet [8]uint8
}

func CreateOctet(octet string) (Octet, error) {
	if len(octet) > 8 {
		return Octet{}, errors.New("octet are 8 bits")
	}
	var bits []uint8
	for _, j := range octet {
		digit, err := strconv.ParseUint(string(j), 10, 8)
		if err != nil {
			return Octet{}, errors.New("octet must contain positive integer, 1s and 0s")
		}
		if digit > 1 {
			return Octet{}, errors.New("octet must not be larger than 1, accepting only 1s and 0s")
		}
		bits = append(bits, uint8(digit))
	}
	return Octet{
		octet: [8]uint8{bits[0], bits[1], bits[2], bits[3], bits[4], bits[5], bits[6], bits[7]},
	}, nil
}

/*
Convert binary octet to decimal
*/
func (o Octet) GetDecimal() int {
	var total int
	p := 128
	for _, j := range o.octet {
		if j == 1 {
			total = total + p
		}
		p = p / 2
	}
	return total
}
