package pagination

import "testing"

//todo

func TestPagination_PagesCount(t *testing.T) {
	type args struct {
		totalCount, perPage int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"less than perPage",
			args{totalCount: 5, perPage: 10},
			1,
		},
		{
			"equals with perPage",
			args{totalCount: 10, perPage: 10},
			1,
		},
		{
			"division with remainder",
			args{totalCount: 11, perPage: 10},
			2,
		},
		{
			"division without remainder",
			args{totalCount: 20, perPage: 10},
			2,
		},
		{
			"division with remainder (2)",
			args{totalCount: 21, perPage: 10},
			3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pagination := Pagination{perPage: tt.args.perPage, totalCount: tt.args.totalCount}
			if got := pagination.PagesCount(); got != tt.want {
				t.Errorf("Pagination.PagesCount() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPagination_PagesRange(t *testing.T) {
	type args struct {
		perPage, totalCount, page, pageButtonsCount int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"first page",
			args{10, 150, 1, 6},
			[]int{1, 6},
			// 1 2 3 4 5 6
			// ^
		},
		{
			"middle border",
			args{10, 150, 4, 6},
			[]int{1, 6},
			// 1 2 3 4 5 6
			//       ^
		},
		{
			"next middle border",
			args{10, 150, 5, 6},
			[]int{2, 7},
			// 2 3 4 5 6 7
			//       ^
		},
		{
			"paginate",
			args{10, 150, 10, 6},
			[]int{7, 12},
			// 7, 8, 9, 10, 11, 12
			//          ^
		},
		{
			"before last pages",
			args{10, 150, 13, 6},
			[]int{10, 15},
			// 10, 11, 12, 13, 14, 15
			// 			   ^
		},
		{
			"last page",
			args{10, 150, 15, 6},
			[]int{10, 15},
			// 10, 11, 12, 13, 14, 15
			// 			           ^
		},

		{
			"first page odd",
			args{10, 150, 1, 7},
			[]int{1, 7},
			// 1 2 3 4 5 6 7
			// ^
		},
		{
			"middle border odd",
			args{10, 150, 4, 7},
			[]int{1, 7},
			// 1 2 3 4 5 6 7
			//       ^
		},
		{
			"next middle border odd",
			args{10, 150, 5, 7},
			[]int{2, 8},
			// 2 3 4 5 6 7 8
			//       ^
		},
		{
			"paginate odd",
			args{10, 150, 10, 7},
			[]int{7, 12},
			// 7, 8, 9, 10, 11, 12
			//          ^
		},
		{
			"before last pages odd",
			args{10, 150, 12, 7},
			[]int{9, 15},
			// 9, 10, 11, 12, 13, 14, 15
			// 			  ^
		},
		{
			"last page odd",
			args{10, 150, 15, 7},
			[]int{9, 15},
			// 9, 10, 11, 12, 13, 14, 15
			// 			              ^
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pagination := Pagination{
				perPage:          tt.args.perPage,
				totalCount:       tt.args.totalCount,
				pageButtonsCount: tt.args.pageButtonsCount,
			}
			if from, to := pagination.PagesRange(); from != tt.want[0] || to != tt.want[1] {
				t.Errorf("Pagination.PagesRange() = %v, %v, want %v", from, to, tt.want)
			}
		})
	}
}
