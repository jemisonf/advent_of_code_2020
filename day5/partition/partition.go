package partition

type Seat struct {
	Row    int
	Column int
	SeatID int
}

func Search(rowIdentifier, columnIdentifier string, numRows, numColumns int) (*Seat, error) {
	max, min := numRows, 0

	for _, char := range rowIdentifier {
		if char == 'F' {
			max -= (max - min) / 2
		} else if char == 'B' {
			min += (max - min) / 2
		}
	}

	row := min

	max, min = numColumns, 0

	for _, char := range columnIdentifier {
		if char == 'L' {
			max -= (max - min) / 2
		} else if char == 'R' {
			min += (max - min) / 2
		}
	}

	column := min

	return &Seat{Row: row, Column: column, SeatID: row*8 + column}, nil
}
