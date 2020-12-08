package partition_test

import (
	"github.com/jemisonf/advent_of_code_2020/day5/partition"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Partition", func() {

	Context("Search", func() {
		testCases := map[string]*partition.Seat{
			"FBFBBFFRLR": {
				Row:    44,
				Column: 5,
				SeatID: 357,
			},
			"BFFFBBFRRR": {
				Row:    70,
				Column: 7,
				SeatID: 567,
			},
			"FFFBBBFRRR": {
				Row:    14,
				Column: 7,
				SeatID: 119,
			},
			"BBFFBBFRLL": {
				Row:    102,
				Column: 4,
				SeatID: 820,
			},
		}

		for identifer, testSeat := range testCases {
			It(identifer, func() {
				seat, err := partition.Search(identifer[:7], identifer[7:], 128, 8)
				Expect(err).NotTo(HaveOccurred())
				Expect(seat).To(Equal(testSeat))
			})
		}
	})
})
