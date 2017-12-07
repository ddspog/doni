package experiment

// import (
// 	"fmt"
// 	"testing"

// 	"gopkg.in/ddspog/mspec.v1/bdd"
// )

// // Feature Create Data with functional Getters
// // - As a developer,
// // - I want to be able to create a Data object and get values with its
// // getters,
// // - So that I could use these getters to manipulate and read data.
// func Test_Create_Data_with_functional_Getters(t *testing.T) {
// 	given, like, s := bdd.Sentences()

// 	given(t, "a Data d with Columns = '%[1]s' and Rows = %[2]v", func(when bdd.When, args ...interface{}) {
// 		tcolumns := args[0].([]string)
// 		tvalues := args[1].([][]interface{})

// 		var d Dater = &Data{
// 			columns: tcolumns,
// 			values:  tvalues,
// 		}

// 		when("d.Columns() is called", func(it bdd.It) {
// 			it(fmt.Sprintf("should return %d", len(tcolumns)), func(assert bdd.Assert) {
// 				assert.Equal(len(tcolumns), d.Columns())
// 			})
// 		})

// 		for i := 0; i < len(tcolumns); i++ {
// 			when(fmt.Sprintf("d.Column(%d) is called", i), func(it bdd.It) {
// 				it(fmt.Sprintf("should return %s", tcolumns[i]), func(assert bdd.Assert) {
// 					assert.Equal(tcolumns[i], d.Column(i))
// 				})
// 			})
// 		}

// 		when("d.Rows() is called", func(it bdd.It) {
// 			it(fmt.Sprintf("should return %d", len(tvalues)), func(assert bdd.Assert) {
// 				assert.Equal(len(tvalues), d.Rows())
// 			})
// 		})

// 		for i := 0; i < len(tvalues); i++ {
// 			for j := 0; j < len(tcolumns); j++ {
// 				when(fmt.Sprintf("d.Value(%d, %d) is called", i, j), func(it bdd.It) {
// 					it(fmt.Sprintf("should return %d", tvalues[i][j]), func(assert bdd.Assert) {
// 						assert.Equal(tvalues[i][j], d.Value(i, j))
// 					})
// 				})
// 			}
// 		}
// 	}, like(
// 		s([]string{"a", "b"}, [][]interface{}{{1, 2}, {2, 3}}),
// 		s([]string{"a", "b", "c"}, [][]interface{}{{1, 2, 3}, {2, 3, 5}}),
// 		s([]string{"a"}, [][]interface{}{{1}, {2}}),
// 	))
// }
