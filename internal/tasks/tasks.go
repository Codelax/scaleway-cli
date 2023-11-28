//go:build !wasm

package tasks

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"reflect"
)

type TaskFunc[T any, U any] func(t *Task, args T) (nextArgs U, err error)
type CleanupFunc func(ctx context.Context) error

type Task struct {
	Name string
	Ctx  context.Context
	Logs *os.File

	taskFunction   TaskFunc[any, any]
	argType        reflect.Type
	returnType     reflect.Type
	cleanFunctions []CleanupFunc
}

type Tasks struct {
	tasks []*Task
}

func Begin() *Tasks {
	return &Tasks{}
}

func Add[TaskArg any, TaskReturn any](ts *Tasks, name string, taskFunc TaskFunc[TaskArg, TaskReturn]) {
	var argValue TaskArg
	var returnValue TaskReturn
	argType := reflect.TypeOf(argValue)
	returnType := reflect.TypeOf(returnValue)

	tasksAmount := len(ts.tasks)
	if tasksAmount > 0 {
		lastTask := ts.tasks[tasksAmount-1]
		if argType != lastTask.returnType {
			panic(fmt.Errorf("invalid task declared, wait for %s, previous task returns %s", argType.Name(), lastTask.returnType.Name()))
		}
	}

	ts.tasks = append(ts.tasks, &Task{
		Name:       name,
		argType:    argType,
		returnType: returnType,
		taskFunction: func(t *Task, i interface{}) (passedData interface{}, err error) {
			if i == nil {
				var zero TaskArg
				passedData, err = taskFunc(t, zero)
			} else {
				passedData, err = taskFunc(t, i.(TaskArg))
			}
			return
		},
	})
}

func (t *Task) AddToCleanUp(cleanupFunc CleanupFunc) {
	t.cleanFunctions = append(t.cleanFunctions, cleanupFunc)
}

// setupContext return a contextWithCancel that will cancel on os interrupt (Ctrl-C)
func setupContext(ctx context.Context) (context.Context, func()) {
	return signal.NotifyContext(ctx, os.Interrupt)
}

// Cleanup execute all tasks cleanup function before failed one in reverse order
func (ts *Tasks) Cleanup(ctx context.Context, failed int) {
	cancelableCtx, cleanCtx := setupContext(ctx)
	defer cleanCtx()
	loader := setupLoader()

	for i := failed; i >= 0; i-- {
		task := ts.tasks[i]

		select {
		case <-cancelableCtx.Done():
			fmt.Println("cleanup has been cancelled, there may be dangling resources")
			return
		default:
		}

		if len(task.cleanFunctions) != 0 {
			var err error
			for i, cleanUpFunc := range task.cleanFunctions {
				loader.Start()
				fmt.Sprintf("Cleaning task %q %d/%d", task.Name, i+1, len(task.cleanFunctions))
				err = cleanUpFunc(cancelableCtx)
				loader.Stop()
				if err != nil {
					break
				}
			}
		}
	}
}

// Execute tasks with interactive display and cleanup on fail
func (ts *Tasks) Execute(ctx context.Context, data interface{}) (interface{}, error) {
	var err error
	loader := setupLoader()
	cancelableCtx, cleanCtx := setupContext(ctx)
	defer cleanCtx()

	for i := range ts.tasks {
		task := ts.tasks[i]

		// Add context and reset cleanup functions, allows to execute multiple times
		task.Ctx = cancelableCtx
		task.cleanFunctions = []CleanupFunc(nil)

		fmt.Printf("[%d/%d] %s\n", i+1, len(ts.tasks), task.Name)
		loader.Start()

		data, err = task.taskFunction(task, data)
		if err != nil {
			loader.Stop()
			ts.Cleanup(ctx, i)

			return nil, fmt.Errorf("task %d %q failed: %w", i+1, task.Name, err)
		}

		select {
		case <-ctx.Done():
			loader.Stop()
			ts.Cleanup(ctx, i)
			return nil, fmt.Errorf("task %d %q failed: context canceled", i+1, task.Name)
		default:
		}

		loader.Stop()
	}

	return data, nil
}
