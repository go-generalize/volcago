//go:build internal
// +build internal

package tests

import (
	"context"
	"fmt"
	"os"
	"reflect"
	"strings"
	"testing"
	"time"

	"cloud.google.com/go/firestore"
	model "github.com/go-generalize/volcago/generator/testfiles/auto"
	"golang.org/x/xerrors"
)

var desc = "Hello, World!"

func initFirestoreClient(t *testing.T) *firestore.Client {
	t.Helper()

	if os.Getenv("FIRESTORE_EMULATOR_HOST") == "" {
		os.Setenv("FIRESTORE_EMULATOR_HOST", "localhost:8000")
	}

	if os.Getenv("FIRESTORE_PROJECT_ID") == "" {
		os.Setenv("FIRESTORE_PROJECT_ID", "project-id-in-google")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := firestore.NewClient(ctx, os.Getenv("FIRESTORE_PROJECT_ID"))

	if err != nil {
		t.Fatalf("failed to initialize firestore client: %+v", err)
	}

	return client
}

func compareTask(t *testing.T, expected, actual *model.Task) {
	t.Helper()

	if actual.ID != expected.ID {
		t.Fatalf("unexpected identity: %s (expected: %s)", actual.ID, expected.ID)
	}

	if !actual.Created.Equal(expected.Created) {
		t.Fatalf("unexpected time: %s(expected: %s)", actual.Created, expected.Created)
	}

	if actual.Desc != expected.Desc {
		t.Fatalf("unexpected desc: %s(expected: %s)", actual.Desc, expected.Created)
	}

	if actual.Done != expected.Done {
		t.Fatalf("unexpected done: %v(expected: %v)", actual.Done, expected.Done)
	}
}

func TestFirestore(t *testing.T) {
	client := initFirestoreClient(t)

	taskRepo := model.NewTaskRepository(client)

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	var ids []string
	defer func() {
		defer cancel()
		if err := taskRepo.DeleteMultiByIDs(ctx, ids); err != nil {
			t.Fatal(err)
		}
	}()

	now := time.Unix(0, time.Now().UnixNano()).UTC()

	t.Run("Multi", func(tr *testing.T) {
		tks := make([]*model.Task, 0)
		for i := int64(1); i <= 10; i++ {
			tk := &model.Task{
				Desc:       fmt.Sprintf("%s%d", desc, i),
				Created:    now,
				Done:       true,
				Done2:      false,
				Count:      int(i),
				Count64:    0,
				Proportion: 0.12345 + float64(i),
				NameList:   []string{"a", "b", "c"},
				Flag: map[string]float64{
					"1": 1.1,
					"2": 2.2,
					"3": 3.3,
				},
			}
			tks = append(tks, tk)
		}
		idList, err := taskRepo.InsertMulti(ctx, tks)
		if err != nil {
			tr.Fatalf("%+v", err)
		}
		ids = append(ids, idList...)

		tks2 := make([]*model.Task, 0)
		for i := int64(1); i <= 10; i++ {
			tk := &model.Task{
				ID:         ids[i-1],
				Desc:       fmt.Sprintf("%s%d", desc, i),
				Created:    now,
				Done:       true,
				Done2:      false,
				Count:      int(i),
				Count64:    i,
				Proportion: 0.12345 + float64(i),
				NameList:   []string{"a", "b", "c"},
				Flag: map[string]float64{
					"4": 4.4,
					"5": 5.5,
					"6": 6.6,
				},
			}
			tks2 = append(tks2, tk)
		}
		if err := taskRepo.UpdateMulti(ctx, tks2); err != nil {
			tr.Fatalf("%+v", err)
		}

		if tks[0].ID != tks2[0].ID {
			tr.Fatalf("unexpected identity: %s (expected: %s)", tks[0].ID, tks2[0].ID)
		}
	})

	t.Run("Single", func(tr *testing.T) {
		tk := &model.Task{
			Desc:       fmt.Sprintf("%s%d", desc, 1001),
			Created:    now,
			Done:       true,
			Done2:      false,
			Count:      11,
			Count64:    11,
			Proportion: 11.12345,
			NameList:   []string{"a", "b", "c"},
			Flag: map[string]float64{
				"1": 1.1,
				"2": 2.2,
				"3": 3.3,
			},
		}
		id, err := taskRepo.Insert(ctx, tk)
		if err != nil {
			tr.Fatalf("%+v", err)
		}
		ids = append(ids, id)

		tk.Count++
		tk.Flag["4"] = 4.4
		if err = taskRepo.Update(ctx, tk); err != nil {
			tr.Fatalf("%+v", err)
		}

		tsk, err := taskRepo.Get(ctx, tk.ID)
		if err != nil {
			tr.Fatalf("%+v", err)
		}

		if tsk.Count != 12 {
			tr.Fatalf("unexpected Count: %d (expected: %d)", tsk.Count, 12)
		}

		if _, ok := tsk.Flag["4"]; !ok {
			tr.Fatalf("unexpected Flag: %v (expected: %v)", ok, true)
		}

		tr.Run("UpdateBuilder", func(ttr *testing.T) {
			desc1002 := fmt.Sprintf("%s%d", desc, 1002)

			updateParam := &model.TaskUpdateParam{
				Desc:       desc1002,
				Created:    firestore.ServerTimestamp,
				Done:       false,
				Count:      firestore.Increment(1),
				Count64:    firestore.Increment(2),
				Proportion: firestore.Increment(0.1),
			}

			if err = taskRepo.StrictUpdate(ctx, tsk.ID, updateParam); err != nil {
				ttr.Fatalf("%+v", err)
			}

			tsk, err = taskRepo.Get(ctx, tk.ID)
			if err != nil {
				ttr.Fatalf("%+v", err)
			}

			if tsk.Desc != desc1002 {
				ttr.Fatalf("unexpected Desc: %s (expected: %s)", tsk.Desc, desc1002)
			}

			if tsk.Created.Before(now) {
				ttr.Fatalf("unexpected Created > now: %t (expected: %t)", tsk.Created.Before(now), tsk.Created.After(now))
			}

			if tsk.Done {
				ttr.Fatalf("unexpected Done: %t (expected: %t)", tsk.Done, false)
			}

			if tsk.Count != 13 {
				ttr.Fatalf("unexpected Count: %d (expected: %d)", tsk.Count, 13)
			}

			if tsk.Count64 != 13 {
				ttr.Fatalf("unexpected Count64: %d (expected: %d)", tsk.Count64, 13)
			}

			if tsk.Proportion != 11.22345 {
				ttr.Fatalf("unexpected Proportion: %g (expected: %g)", tsk.Proportion, 11.22345)
			}
		})
	})
}

func TestFirestoreTransaction_Single(t *testing.T) {
	client := initFirestoreClient(t)

	taskRepo := model.NewTaskRepository(client)

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	var ids []string
	defer func() {
		defer cancel()
		if err := taskRepo.DeleteMultiByIDs(ctx, ids); err != nil {
			t.Fatal(err)
		}
	}()

	now := time.Unix(0, time.Now().UnixNano())

	t.Run("Insert", func(tr *testing.T) {
		err := client.RunTransaction(ctx, func(cx context.Context, tx *firestore.Transaction) error {
			tk := &model.Task{
				Desc:       fmt.Sprintf("%s01", desc),
				Created:    now,
				Done:       true,
				Done2:      false,
				Count:      10,
				Count64:    11,
				Proportion: 11.12345,
				NameList:   []string{"a", "b", "c"},
				Flag: map[string]float64{
					"1": 1.1,
					"2": 2.2,
					"3": 3.3,
				},
			}

			id, err := taskRepo.InsertWithTx(cx, tx, tk)
			if err != nil {
				return err
			}

			ids = append(ids, id)
			return nil
		})

		if err != nil {
			tr.Fatalf("error: %+v", err)
		}

		tsk, err := taskRepo.Get(ctx, ids[len(ids)-1])
		if err != nil {
			tr.Fatalf("%+v", err)
		}

		if tsk.Count != 10 {
			tr.Fatalf("unexpected Count: %d (expected: %d)", tsk.Count, 10)
		}
	})

	t.Run("Update", func(tr *testing.T) {
		tkID := ids[len(ids)-1]
		err := client.RunTransaction(ctx, func(cx context.Context, tx *firestore.Transaction) error {
			tk, err := taskRepo.GetWithTx(tx, tkID)
			if err != nil {
				return err
			}

			if tk.Count != 10 {
				return fmt.Errorf("unexpected Count: %d (expected: %d)", tk.Count, 10)
			}

			tk.Count = 11
			if err = taskRepo.UpdateWithTx(cx, tx, tk); err != nil {
				return err
			}
			return nil
		})

		if err != nil {
			tr.Fatalf("error: %+v", err)
		}

		tsk, err := taskRepo.Get(ctx, tkID)
		if err != nil {
			tr.Fatalf("%+v", err)
		}

		if tsk.Count != 11 {
			tr.Fatalf("unexpected Count: %d (expected: %d)", tsk.Count, 11)
		}
	})

	t.Run("UseUpdateBuilder", func(tr *testing.T) {
		tkID := ids[len(ids)-1]
		desc1002 := fmt.Sprintf("%s%d", desc, 1002)
		err := client.RunTransaction(ctx, func(cx context.Context, tx *firestore.Transaction) error {
			tk, err := taskRepo.GetWithTx(tx, tkID)
			if err != nil {
				return err
			}

			if tk.Count != 11 {
				return fmt.Errorf("unexpected Count: %d (expected: %d)", tk.Count, 11)
			}

			updateParam := &model.TaskUpdateParam{
				Desc:       desc1002,
				Created:    firestore.ServerTimestamp,
				Done:       false,
				Count:      firestore.Increment(1),
				Count64:    firestore.Increment(2),
				Proportion: firestore.Increment(0.1),
			}
			if err = taskRepo.StrictUpdateWithTx(tx, tk.ID, updateParam); err != nil {
				return err
			}
			return nil
		})

		if err != nil {
			tr.Fatalf("error: %+v", err)
		}

		tsk, err := taskRepo.Get(ctx, tkID)
		if err != nil {
			tr.Fatalf("%+v", err)
		}

		if tsk.Desc != desc1002 {
			tr.Fatalf("unexpected Desc: %s (expected: %s)", tsk.Desc, desc1002)
		}

		if tsk.Created.Before(now) {
			tr.Fatalf("unexpected Created > now: %t (expected: %t)", tsk.Created.Before(now), tsk.Created.After(now))
		}

		if tsk.Done {
			tr.Fatalf("unexpected Done: %t (expected: %t)", tsk.Done, false)
		}

		if tsk.Count != 12 {
			tr.Fatalf("unexpected Count: %d (expected: %d)", tsk.Count, 12)
		}

		if tsk.Count64 != 13 {
			tr.Fatalf("unexpected Count64: %d (expected: %d)", tsk.Count64, 13)
		}

		if tsk.Proportion != 11.22345 {
			tr.Fatalf("unexpected Proportion: %g (expected: %g)", tsk.Proportion, 11.22345)
		}
	})
}

func TestFirestoreTransaction_Multi(t *testing.T) {
	client := initFirestoreClient(t)

	taskRepo := model.NewTaskRepository(client)

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	var ids []string
	defer func() {
		defer cancel()
		if err := taskRepo.DeleteMultiByIDs(ctx, ids); err != nil {
			t.Fatal(err)
		}
	}()

	now := time.Unix(0, time.Now().UnixNano())

	tks := make([]*model.Task, 0)
	t.Run("InsertMulti", func(tr *testing.T) {
		for i := int64(1); i <= 10; i++ {
			tk := &model.Task{
				ID:         fmt.Sprintf("Task_%d", i),
				Desc:       fmt.Sprintf("%s%d", desc, i),
				Created:    now,
				Done:       true,
				Done2:      false,
				Count:      int(i),
				Count64:    0,
				NameList:   []string{"a", "b", "c"},
				Proportion: 0.12345 + float64(i),
				Flag: map[string]float64{
					"1": 1.1,
					"2": 2.2,
					"3": 3.3,
				},
			}
			tks = append(tks, tk)
		}

		if err := client.RunTransaction(ctx, func(cx context.Context, tx *firestore.Transaction) error {
			idList, err := taskRepo.InsertMultiWithTx(ctx, tx, tks)
			if err != nil {
				return err
			}
			ids = append(ids, idList...)
			return nil
		}); err != nil {
			tr.Fatalf("error: %+v", err)
		}
	})

	tks2 := make([]*model.Task, 0)
	t.Run("UpdateMulti", func(tr *testing.T) {
		for i := int64(1); i <= 10; i++ {
			tk := &model.Task{
				ID:         ids[i-1],
				Desc:       fmt.Sprintf("%s%d", desc, i+1),
				Created:    now,
				Done:       false,
				Done2:      true,
				Count:      int(i),
				Count64:    i,
				NameList:   []string{"a", "b", "c"},
				Proportion: 0.12345 + float64(i),
				Flag: map[string]float64{
					"1": 1.1,
					"2": 2.2,
					"3": 3.3,
				},
			}
			tks2 = append(tks2, tk)
		}

		if err := client.RunTransaction(ctx, func(cx context.Context, tx *firestore.Transaction) error {
			if err := taskRepo.UpdateMultiWithTx(cx, tx, tks2); err != nil {
				return err
			}
			return nil
		}); err != nil {
			tr.Fatalf("error: %+v", err)
		}

		if tks[0].ID != tks2[0].ID {
			tr.Fatalf("unexpected identity: %s (expected: %s)", tks[0].ID, tks2[0].ID)
		}
	})
}

func TestFirestoreQuery(t *testing.T) {
	client := initFirestoreClient(t)

	taskRepo := model.NewTaskRepository(client)

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	var ids []string
	defer func() {
		defer cancel()
		if err := taskRepo.DeleteMultiByIDs(ctx, ids); err != nil {
			t.Fatalf("%+v\n", err)
		}
	}()

	now := time.Unix(0, time.Now().UnixNano())

	tks := make([]*model.Task, 0)
	for i := 1; i <= 10; i++ {
		tk := &model.Task{
			ID:         fmt.Sprintf("%d", i),
			Desc:       fmt.Sprintf("%s%d", desc, i),
			Created:    now,
			Done:       true,
			Done2:      false,
			Count:      i,
			Count64:    int64(i),
			NameList:   []string{"a", "b", "c"},
			Proportion: 0.12345 + float64(i),
			Flag: map[string]float64{
				"1": 1.1,
				"2": 2.2,
				"3": 3.3,
				"4": 4.4,
				"5": 5.5,
			},
			SliceSubTask: []*model.SubTask{
				{
					Name: "slice_nested",
				},
			},
		}
		tks = append(tks, tk)
	}
	ids, err := taskRepo.InsertMulti(ctx, tks)
	if err != nil {
		t.Fatalf("%+v", err)
	}

	t.Run("int(1件)", func(tr *testing.T) {
		param := &model.TaskSearchParam{
			Count: model.NewQueryChainer().Equal(1),
		}

		tasks, err := taskRepo.Search(ctx, param, nil)
		if err != nil {
			tr.Fatalf("%+v", err)
		}

		if len(tasks) != 1 {
			tr.Fatalf("unexpected length: %d (expected: %d)", len(tasks), 10)
		}
	})

	t.Run("int64(5件)", func(tr *testing.T) {
		param := &model.TaskSearchParam{
			Count64: model.NewQueryChainer().LessThanOrEqual(5),
		}

		tasks, err := taskRepo.Search(ctx, param, nil)
		if err != nil {
			tr.Fatalf("%+v", err)
		}

		if len(tasks) != 5 {
			tr.Fatalf("unexpected length: %d (expected: %d)", len(tasks), 5)
		}
	})

	t.Run("float(1件)", func(tr *testing.T) {
		param := &model.TaskSearchParam{
			Proportion: model.NewQueryChainer().Equal(1.12345),
		}

		tasks, err := taskRepo.Search(ctx, param, nil)
		if err != nil {
			tr.Fatalf("%+v", err)
		}

		if len(tasks) != 1 {
			tr.Fatalf("unexpected length: %d (expected: %d)", len(tasks), 1)
		}
	})

	t.Run("bool(10件)", func(tr *testing.T) {
		param := &model.TaskSearchParam{
			Done: model.NewQueryChainer().Equal(true),
		}

		tasks, err := taskRepo.Search(ctx, param, nil)
		if err != nil {
			tr.Fatalf("%+v", err)
		}

		if len(tasks) != 10 {
			tr.Fatalf("unexpected length: %d (expected: %d)", len(tasks), 10)
		}
	})

	t.Run("time.Time(10件)", func(tr *testing.T) {
		param := &model.TaskSearchParam{
			Created: model.NewQueryChainer().Equal(now),
		}

		tasks, err := taskRepo.Search(ctx, param, nil)
		if err != nil {
			tr.Fatalf("%+v", err)
		}

		if len(tasks) != 10 {
			tr.Fatalf("unexpected length: %d (expected: %d)", len(tasks), 10)
		}
	})

	t.Run("[]string(10件)", func(tr *testing.T) {
		param := &model.TaskSearchParam{
			NameList: model.NewQueryChainer().ArrayContainsAny([]string{"a", "b"}),
		}

		tasks, err := taskRepo.Search(ctx, param, nil)
		if err != nil {
			tr.Fatalf("%+v", err)
		}

		if len(tasks) != 10 {
			tr.Fatalf("unexpected length: %d (expected: %d)", len(tasks), 10)
		}
	})

	t.Run("[]Object(10件)", func(tr *testing.T) {
		param := &model.TaskSearchParam{
			SliceSubTask: model.NewQueryChainer().ArrayContainsAny([]*model.SubTask{{Name: "slice_struct"}}),
		}

		tasks, err := taskRepo.Search(ctx, param, nil)
		if err != nil {
			tr.Fatalf("%+v", err)
		}

		if len(tasks) != 10 {
			tr.Fatalf("unexpected length: %d (expected: %d)", len(tasks), 10)
		}
	})

	t.Run("[]Object(0件)", func(tr *testing.T) {
		param := &model.TaskSearchParam{
			SliceSubTask: model.NewQueryChainer().ArrayContains("slice_struct"),
		}

		tasks, err := taskRepo.Search(ctx, param, nil)
		if err != nil {
			tr.Fatalf("%+v", err)
		}

		if len(tasks) != 0 {
			tr.Fatalf("unexpected length: %d (expected: %d)", len(tasks), 10)
		}
	})

	t.Run("Flag(10件)", func(tr *testing.T) {
		param := &model.TaskSearchParam{
			Flag: model.NewQueryChainer().Equal(map[string]float64{
				"1": 1.1,
				"2": 2.2,
				"3": 3.3,
			}),
		}

		tasks, err := taskRepo.Search(ctx, param, nil)
		if err != nil {
			tr.Fatalf("%+v", err)
		}

		if len(tasks) != 10 {
			tr.Fatalf("unexpected length: %d (expected: %d)", len(tasks), 10)
		}
	})

	t.Run("NotEqual(9件)", func(tr *testing.T) {
		description := fmt.Sprintf("%s%d", desc, 1)
		param := &model.TaskSearchParam{
			Desc: model.NewQueryChainer().NotEqual(description),
		}
		tasks, err := taskRepo.Search(ctx, param, nil)
		if err != nil {
			tr.Fatalf("%+v", err)
		}
		if len(tasks) != 9 {
			tr.Fatalf("unexpected length: %d (expected: %d)", len(tasks), 9)
		}
	})

	t.Run("NotIn(8件)", func(tr *testing.T) {
		description1 := fmt.Sprintf("%s%d", desc, 1)
		description2 := fmt.Sprintf("%s%d", desc, 2)
		param := &model.TaskSearchParam{
			Desc: model.NewQueryChainer().NotIn([]string{description1, description2}),
		}
		tasks, err := taskRepo.Search(ctx, param, nil)
		if err != nil {
			tr.Fatalf("%+v", err)
		}
		if len(tasks) != 8 {
			tr.Fatalf("unexpected length: %d (expected: %d)", len(tasks), 8)
		}
	})

	t.Run("UseQueryBuilder", func(tr *testing.T) {
		tr.Run("range query(3<count<8)", func(ttr *testing.T) {
			qb := model.NewQueryBuilder(taskRepo.GetCollection())
			qb.GreaterThan("count", 3)
			qb.LessThan("count", 8)

			tasks, err := taskRepo.Search(ctx, nil, qb.Query())
			if err != nil {
				ttr.Fatalf("%+v", err)
			}

			if len(tasks) != 4 {
				ttr.Fatalf("unexpected length: %d (expected: %d)", len(tasks), 4)
			}
		})
		tr.Run("!=(count!=1)", func(ttr *testing.T) {
			qb := model.NewQueryBuilder(taskRepo.GetCollection())
			qb.NotEqual("count", 1)

			tasks, err := taskRepo.Search(ctx, nil, qb.Query())
			if err != nil {
				ttr.Fatalf("%+v", err)
			}

			if len(tasks) != 9 {
				ttr.Fatalf("unexpected length: %d (expected: %d)", len(tasks), 9)
			}
		})
		tr.Run("not-in(count not-in [1,2,3,4,5])", func(ttr *testing.T) {
			qb := model.NewQueryBuilder(taskRepo.GetCollection())
			qb.NotIn("count", []int{1, 2, 3, 4, 5})

			tasks, err := taskRepo.Search(ctx, nil, qb.Query())
			if err != nil {
				ttr.Fatalf("%+v", err)
			}

			if len(tasks) != 5 {
				ttr.Fatalf("unexpected length: %d (expected: %d)", len(tasks), 5)
			}
		})
	})

	t.Run("Indexes", func(tr *testing.T) {
		tr.Run("Equal", func(ttr *testing.T) {
			param := &model.TaskSearchParam{
				Desc: model.NewQueryChainer().Filters("Hello, World!1"),
			}

			tasks, err := taskRepo.Search(ctx, param, nil)
			if err != nil {
				ttr.Fatalf("%+v", err)
			}

			if len(tasks) != 1 {
				ttr.Fatalf("unexpected length: %d (expected: %d)", len(tasks), 1)
			}
		})
		tr.Run("Prefix", func(ttr *testing.T) {
			chainer := model.NewQueryChainer
			param := &model.TaskSearchParam{
				Desc: chainer().Filters("Hel", model.FilterTypeAddPrefix),
				Done: chainer().Equal(true),
			}

			tasks, err := taskRepo.Search(ctx, param, nil)
			if err != nil {
				ttr.Fatalf("%+v", err)
			}

			if len(tasks) != 10 {
				ttr.Fatalf("unexpected length: %d (expected: %d)", len(tasks), 10)
			}
		})
		tr.Run("Suffix", func(ttr *testing.T) {
			param := &model.TaskSearchParam{
				Desc: model.NewQueryChainer().Filters("10", model.FilterTypeAddSuffix),
			}

			tasks, err := taskRepo.Search(ctx, param, nil)
			if err != nil {
				ttr.Fatalf("%+v", err)
			}

			if len(tasks) != 1 {
				ttr.Fatalf("unexpected length: %d (expected: %d)", len(tasks), 1)
			}
		})
		tr.Run("Biunigrams", func(ttr *testing.T) {
			ttr.Run("1", func(ttrr *testing.T) {
				param := &model.TaskSearchParam{
					Desc: model.NewQueryChainer().Filters("o, Wor", model.FilterTypeAddBiunigrams),
				}

				tasks, err := taskRepo.Search(ctx, param, nil)
				if err != nil {
					ttrr.Fatalf("%+v", err)
				}

				if len(tasks) != 10 {
					ttrr.Fatalf("unexpected length: %d (expected: %d)", len(tasks), 10)
				}
			})
			ttr.Run("2", func(ttrr *testing.T) {
				param := &model.TaskSearchParam{
					Desc: model.NewQueryChainer().Filters("!1", model.FilterTypeAddBiunigrams),
				}

				tasks, err := taskRepo.Search(ctx, param, nil)
				if err != nil {
					ttr.Fatalf("%+v", err)
				}

				if len(tasks) != 2 {
					ttr.Fatalf("unexpected length: %d (expected: %d)", len(tasks), 2)
				}
			})
			ttr.Run("3", func(ttrr *testing.T) {
				param := &model.TaskSearchParam{
					Desc: model.NewQueryChainer().Filters("Hello, W", model.FilterTypeAddBiunigrams),
				}

				tasks, err := taskRepo.Search(ctx, param, nil)
				if err != nil {
					ttr.Fatalf("%+v", err)
				}

				if len(tasks) != 10 {
					ttr.Fatalf("unexpected length: %d (expected: %d)", len(tasks), 10)
				}
			})
			ttr.Run("NG", func(ttrr *testing.T) {
				param := &model.TaskSearchParam{
					Desc: model.NewQueryChainer().Filters("Hello,W", model.FilterTypeAddBiunigrams),
				}

				tasks, err := taskRepo.Search(ctx, param, nil)
				if err != nil {
					ttr.Fatalf("%+v", err)
				}

				if len(tasks) != 0 {
					ttr.Fatalf("unexpected length: %d (expected: %d)", len(tasks), 0)
				}
			})
		})
		tr.Run("Something", func(ttr *testing.T) {
			param := &model.TaskSearchParam{
				Proportion: model.NewQueryChainer().Filters(10.12345, model.FilterTypeAddSomething),
			}

			tasks, err := taskRepo.Search(ctx, param, nil)
			if err != nil {
				ttr.Fatalf("%+v", err)
			}

			if len(tasks) != 1 {
				ttr.Fatalf("unexpected length: %d (expected: %d)", len(tasks), 1)
			}
		})
	})
}

func TestFirestoreError(t *testing.T) {
	client := initFirestoreClient(t)

	taskRepo := model.NewTaskRepository(client)

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	var ids []string
	defer func() {
		defer cancel()
		if err := taskRepo.DeleteMultiByIDs(ctx, ids); err != nil {
			t.Fatalf("%+v\n", err)
		}
	}()

	now := time.Unix(0, time.Now().UnixNano())

	t.Run("Prepare", func(tr *testing.T) {
		tk := &model.Task{
			Desc:       desc,
			Created:    now,
			Done:       true,
			Done2:      false,
			Count:      11,
			Count64:    11,
			Proportion: 0.12345 + 11,
			NameList:   []string{"a", "b", "c"},
			Flag: map[string]float64{
				"1": 1.1,
				"2": 2.2,
				"3": 3.3,
			},
		}
		id, err := taskRepo.Insert(ctx, tk)
		if err != nil {
			tr.Fatalf("%+v", err)
		}
		ids = append(ids, id)
	})

	t.Run("ErrorReadAfterWrite", func(tr *testing.T) {
		tkID := ids[len(ids)-1]
		errReadAfterWrite := xerrors.New("firestore: read after write in transaction")

		if err := client.RunTransaction(ctx, func(cx context.Context, tx *firestore.Transaction) error {
			tk, err := taskRepo.GetWithTx(tx, tkID)
			if err != nil {
				return err
			}

			if tk.Count != 11 {
				return fmt.Errorf("unexpected Count: %d (expected: %d)", tk.Count, 11)
			}

			tk.Count = 12
			if err = taskRepo.UpdateWithTx(cx, tx, tk); err != nil {
				return err
			}

			if _, err = taskRepo.GetWithTx(tx, tkID); err != nil {
				return err
			}
			return nil
		}); err != nil && xerrors.Is(xerrors.Unwrap(err), errReadAfterWrite) {
			tr.Fatalf("error: %+v", err)
		}

		tsk, err := taskRepo.Get(ctx, tkID)
		if err != nil {
			tr.Fatalf("%+v", err)
		}

		if tsk.Count != 11 {
			tr.Fatalf("unexpected Count: %d (expected: %d)", tsk.Count, 11)
		}

		if err = client.RunTransaction(ctx, func(cx context.Context, tx *firestore.Transaction) error {
			id, er := taskRepo.InsertWithTx(cx, tx, new(model.Task))
			if er != nil {
				return er
			}

			if _, err = taskRepo.GetWithTx(tx, id); err != nil {
				return err
			}
			return nil
		}); err != nil && xerrors.Is(xerrors.Unwrap(err), errReadAfterWrite) {
			tr.Fatalf("error: %+v", err)
		}
	})
}

func TestFirestoreOfTaskRepo(t *testing.T) {
	client := initFirestoreClient(t)

	taskRepo := model.NewTaskRepository(client)

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	now := time.Unix(time.Now().Unix(), 0)

	id, err := taskRepo.Insert(ctx, &model.Task{
		Desc:    desc,
		Created: now,
		Done:    true,
	})

	if err != nil {
		t.Fatalf("failed to put item: %+v", err)
	}

	ret, err := taskRepo.Get(ctx, id)

	if err != nil {
		t.Fatalf("failed to get item: %+v", err)
	}

	compareTask(t, &model.Task{
		ID:      id,
		Desc:    desc,
		Created: now,
		Done:    true,
	}, ret)

	returns, err := taskRepo.GetMulti(ctx, []string{id})

	if err != nil {
		t.Fatalf("failed to get item: %+v", err)
	}

	if len(returns) != 1 {
		t.Fatalf("GetMulti should return 1 item: %#v", returns)
	}

	compareTask(t, &model.Task{
		ID:      id,
		Desc:    desc,
		Created: now,
		Done:    true,
	}, returns[0])

	compareTask(t, &model.Task{
		ID:      id,
		Desc:    desc,
		Created: now,
		Done:    true,
	}, ret)

	if err := taskRepo.DeleteByID(ctx, id); err != nil {
		t.Fatalf("delete failed: %+v", err)
	}

	if _, err := taskRepo.Get(ctx, id); err == nil {
		t.Fatalf("should get an error")
	}
}

func TestFirestoreOfLockRepo(t *testing.T) {
	client := initFirestoreClient(t)

	lockRepo := model.NewLockRepository(client)

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	ids := make([]string, 0)
	defer func() {
		defer cancel()
		mode := model.DeleteOption{Mode: model.DeleteModeHard}
		if err := lockRepo.DeleteMultiByIDs(ctx, ids, mode); err != nil {
			t.Fatal(err)
		}
	}()

	text := "hello"

	t.Run("insert_test", func(tr *testing.T) {
		l := &model.Lock{
			Text: text,
			Flag: nil,
			Meta: model.Meta{},
		}

		id, err := lockRepo.Insert(ctx, l)
		if err != nil {
			tr.Fatalf("failed to put item: %+v", err)
		}

		ids = append(ids, id)

		ret, err := lockRepo.Get(ctx, id)

		if err != nil {
			tr.Fatalf("failed to get item: %+v", err)
		}

		if text != ret.Text {
			tr.Fatalf("unexpected Text: %s (expected: %s)", ret.Text, text)
		}
		if ret.CreatedAt.IsZero() {
			tr.Fatal("unexpected createdAt zero")
		}
		if ret.UpdatedAt.IsZero() {
			tr.Fatal("unexpected updatedAt zero")
		}
	})

	t.Run("update_test", func(tr *testing.T) {
		l := &model.Lock{
			Text: text,
			Flag: nil,
			Meta: model.Meta{},
		}

		id, err := lockRepo.Insert(ctx, l)
		if err != nil {
			tr.Fatalf("failed to put item: %+v", err)
		}

		ids = append(ids, id)

		text = "hello!!!"
		l.Text = text
		err = lockRepo.Update(ctx, l)
		if err != nil {
			tr.Fatalf("failed to update item: %+v", err)
		}

		ret, err := lockRepo.Get(ctx, id)
		if err != nil {
			tr.Fatalf("failed to get item: %+v", err)
		}

		if text != ret.Text {
			tr.Fatalf("unexpected Text: %s (expected: %s)", ret.Text, text)
		}
		if ret.CreatedAt.Equal(ret.UpdatedAt) {
			tr.Fatalf("unexpected CreatedAt == updatedAt: %d == %d",
				ret.CreatedAt.Unix(), ret.UpdatedAt.Unix())
		}
	})

	t.Run("soft_delete_test", func(tr *testing.T) {
		l := &model.Lock{
			Text: text,
			Flag: nil,
			Meta: model.Meta{},
		}

		id, err := lockRepo.Insert(ctx, l)
		if err != nil {
			tr.Fatalf("failed to put item: %+v", err)
		}

		ids = append(ids, id)

		l.Text = text
		err = lockRepo.Delete(ctx, l, model.DeleteOption{
			Mode: model.DeleteModeSoft,
		})
		if err != nil {
			tr.Fatalf("failed to soft delete item: %+v", err)
		}

		ret, err := lockRepo.Get(ctx, id, model.GetOption{
			IncludeSoftDeleted: true,
		})
		if err != nil {
			tr.Fatalf("failed to get item: %+v", err)
		}

		if text != ret.Text {
			tr.Fatalf("unexpected Text: %s (expected: %s)", ret.Text, text)
		}
		if ret.DeletedAt == nil {
			tr.Fatalf("unexpected DeletedAt == nil: %+v", ret.DeletedAt)
		}
	})

	t.Run("hard_delete_test", func(tr *testing.T) {
		l := &model.Lock{
			Text: text,
			Flag: nil,
			Meta: model.Meta{},
		}

		id, err := lockRepo.Insert(ctx, l)
		if err != nil {
			tr.Fatalf("failed to put item: %+v", err)
		}

		l.Text = text
		err = lockRepo.Delete(ctx, l)
		if err != nil {
			tr.Fatalf("failed to hard delete item: %+v", err)
		}

		ret, err := lockRepo.Get(ctx, id, model.GetOption{
			IncludeSoftDeleted: true,
		})
		if err != nil && !strings.Contains(err.Error(), "not found") {
			tr.Fatalf("failed to get item: %+v", err)
		}

		if ret != nil {
			tr.Fatalf("failed to delete item (found!): %+v", ret)
		}
	})

	t.Run("UseQueryBuilder", func(tr *testing.T) {
		l := &model.Lock{
			Text: text,
			Flag: nil,
			Meta: model.Meta{},
		}
		id, err := lockRepo.Insert(ctx, l)
		if err != nil {
			tr.Fatalf("failed to put item: %+v", err)
		}

		ids = append(ids, id)

		qb := model.NewQueryBuilder(lockRepo.GetCollection())
		qb.GreaterThanOrEqual("createdAt", model.SetLastThreeToZero(l.CreatedAt).Add(-100))
		qb.LessThanOrEqual("createdAt", model.SetLastThreeToZero(l.CreatedAt).Add(100))
		if err = qb.Check(); err != nil {
			tr.Fatal(err)
		}

		locks, err := lockRepo.Search(ctx, nil, qb.Query())
		if err != nil {
			tr.Fatalf("%+v", err)
		}

		if len(locks) != 1 {
			tr.Fatalf("unexpected length: %d (expected: %d)", len(locks), 1)
		}

		if id != locks[0].ID {
			tr.Fatalf("unexpected length: %s (expected: %s)", locks[0].ID, id)
		}
	})

	t.Run("UseQueryChainer", func(tr *testing.T) {
		l := &model.Lock{
			Text: "Hello",
			Flag: nil,
			Meta: model.Meta{},
		}
		id, err := lockRepo.Insert(ctx, l)
		if err != nil {
			tr.Fatalf("failed to put item: %+v", err)
		}
		ids = append(ids, id)
		l = &model.Lock{
			Text: "World",
			Flag: nil,
			Meta: model.Meta{},
		}
		id, err = lockRepo.Insert(ctx, l)
		if err != nil {
			tr.Fatalf("failed to put item: %+v", err)
		}
		ids = append(ids, id)
		param := &model.LockSearchParam{
			Text:               model.NewQueryChainer().In([]string{"Hello", "World"}),
			IncludeSoftDeleted: true,
		}
		locks, err := lockRepo.Search(ctx, param, nil)
		if err != nil {
			tr.Fatalf("%+v", err)
		}
		if len(locks) != 2 {
			tr.Fatalf("unexpected length: %d (expected: %d)", len(locks), 2)
		}

		now := time.Now()
		param = &model.LockSearchParam{
			CreatedAt:          model.NewQueryChainer().GreaterThanOrEqual(now.Add(time.Second * 5 * -1)).LessThanOrEqual(now.Add(time.Second * 5)),
			IncludeSoftDeleted: true,
		}
		locks, err = lockRepo.Search(ctx, param, nil)
		if err != nil {
			tr.Fatalf("%+v", err)
		}
		if len(locks) != 6 {
			tr.Fatalf("unexpected length: %d (expected: %d)", len(locks), 6)
		}

		param = &model.LockSearchParam{
			CreatedAt:          model.NewQueryChainer().Asc(),
			CursorLimit:        5,
			IncludeSoftDeleted: true,
		}
		locks, err = lockRepo.Search(ctx, param, nil)
		if err != nil {
			tr.Fatalf("%+v", err)
		}

		if len(locks) != 5 {
			tr.Fatalf("unexpected length: %d (expected: %d)", len(locks), 5)
		}

		param = &model.LockSearchParam{
			CreatedAt:          model.NewQueryChainer().Asc().StartAfter(locks[len(locks)-1].CreatedAt),
			CursorLimit:        5,
			IncludeSoftDeleted: true,
		}
		locks, err = lockRepo.Search(ctx, param, nil)
		if err != nil {
			tr.Fatalf("%+v", err)
		}

		if len(locks) != 1 {
			tr.Fatalf("unexpected length: %d (expected: %d)", len(locks), 1)
		}

		param = &model.LockSearchParam{
			CreatedAt:          model.NewQueryChainer().Asc().EndAt(locks[len(locks)-1].CreatedAt),
			IncludeSoftDeleted: true,
		}
		locks, err = lockRepo.Search(ctx, param, nil)
		if err != nil {
			tr.Fatalf("%+v", err)
		}

		if len(locks) != 6 {
			tr.Fatalf("unexpected length: %d (expected: %d)", len(locks), 6)
		}

		param = &model.LockSearchParam{
			CreatedAt:          model.NewQueryChainer().Asc().EndBefore(locks[len(locks)-1].CreatedAt),
			IncludeSoftDeleted: true,
		}
		locks, err = lockRepo.Search(ctx, param, nil)
		if err != nil {
			tr.Fatalf("%+v", err)
		}

		if len(locks) != 5 {
			tr.Fatalf("unexpected length: %d (expected: %d)", len(locks), 5)
		}
	})

	t.Run("update_builder", func(tr *testing.T) {
		l := &model.Lock{
			Text: text,
			Flag: nil,
			Meta: model.Meta{},
		}

		id, err := lockRepo.Insert(ctx, l)
		if err != nil {
			tr.Fatalf("failed to put item: %+v", err)
		}

		ids = append(ids, id)

		flag := map[string]float64{"test": 123.456}
		hello := fmt.Sprintf("%s world", text)

		t := time.NewTicker(1 * time.Millisecond)
		defer t.Stop()
		<-t.C

		updateParam := &model.LockUpdateParam{
			Text:      hello,
			Flag:      flag,
			UpdatedAt: firestore.ServerTimestamp,
			Version:   firestore.Increment(1),
		}

		if err = lockRepo.StrictUpdate(ctx, id, updateParam); err != nil {
			tr.Fatalf("failed to update item: %+v", err)
		}

		ret, err := lockRepo.Get(ctx, id)
		if err != nil {
			tr.Fatalf("failed to get item: %+v", err)
		}

		if ret.Text != hello {
			tr.Fatalf("unexpected Text: %s (expected: %s)", ret.Text, hello)
		}

		if !reflect.DeepEqual(ret.Flag, flag) {
			tr.Fatalf("unexpected Flag: %v (expected: %v)", ret.Flag, flag)
		}

		if ret.CreatedAt.Equal(ret.UpdatedAt) {
			tr.Fatalf("unexpected CreatedAt == UpdatedAt: %d == %d",
				ret.CreatedAt.Unix(), ret.UpdatedAt.Unix())
		}

		if ret.UpdatedAt.Before(ret.CreatedAt) {
			tr.Fatalf("unexpected UpdatedAt > CreatedAt: %t (expected: %t)", ret.UpdatedAt.Before(ret.CreatedAt), ret.UpdatedAt.After(ret.CreatedAt))
		}

		if ret.Version != 2 {
			tr.Fatalf("unexpected Version: %d (expected: %d)", ret.Version, 2)
		}
	})
}
