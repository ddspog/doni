package experiment

// import (
// 	"fmt"
// 	"testing"

// 	"gopkg.in/ddspog/mspec.v1/bdd"
// )

// // Feature Create Experiment with functional Getters
// // - As a developer,
// // - I want to be able to create a Experiment and access data with its
// // getters,
// // - So that I could use these getters to manipulate and read data.
// func Test_Create_Experiment_with_functional_Getters(t *testing.T) {
// 	given, like, s := bdd.Sentences()

// 	given(t, "a Experiment e with Desc = '%[1]s', Data = %[2]v, "+
// 		"Collect = %[3]v, Phase = %[4]v", func(when bdd.When, args ...interface{}) {
// 		tdesc := args[0].(string)
// 		tdata := args[1].(Dater)
// 		tcollect := args[2].([]Collecter)
// 		tphases := args[3].([]Phaser)

// 		var e Experimenter = &Experiment{
// 			desc:    tdesc,
// 			data:    tdata,
// 			collect: tcollect,
// 			phases:  tphases,
// 		}

// 		when("e.Desc() is called", func(it bdd.It) {
// 			it("should return '%[1]s'", func(assert bdd.Assert) {
// 				assert.Equal(args[0].(string), e.Desc())
// 			})
// 		})

// 		when("e.Data().Columns() is called", func(it bdd.It) {
// 			it(fmt.Sprintf("should return %d", tdata.Columns()), func(assert bdd.Assert) {
// 				assert.Equal(tdata.Columns(), e.Data().Columns())
// 			})
// 		})

// 		when("e.Data().Rows() is called", func(it bdd.It) {
// 			it(fmt.Sprintf("should return %d", tdata.Rows()), func(assert bdd.Assert) {
// 				assert.Equal(tdata.Rows(), e.Data().Rows())
// 			})
// 		})

// 		when("e.CollectMethods() is called", func(it bdd.It) {
// 			it(fmt.Sprintf("should return %d", len(tcollect)), func(assert bdd.Assert) {
// 				assert.Equal(len(tcollect), e.CollectMethods())
// 			})
// 		})

// 		when("e.Phases() is called", func(it bdd.It) {
// 			it(fmt.Sprintf("should return %d", len(tphases)), func(assert bdd.Assert) {
// 				assert.Equal(len(tphases), e.Phases())
// 			})
// 		})
// 	}, like(
// 		s("Random test description 01.", genDater(2, 3), genCollecterArr(2), genPhaserArr(5)),
// 		s("Random test description 02.", genDater(1, 2), genCollecterArr(3), genPhaserArr(2)),
// 		s("Random test description 03.", genDater(4, 2), genCollecterArr(1), genPhaserArr(3)),
// 	))
// }
