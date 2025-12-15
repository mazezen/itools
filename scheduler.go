package itools

import (
	"context"
	"sync"
	"time"
)

type task struct {
	ID       int
	Name     string
	ExecTime time.Time
	Job      func(ctx context.Context)
	Ctx      context.Context
	Cancel   context.CancelFunc
}

// Scheduler 调度器结构
type Scheduler struct {
	mu    sync.Mutex
	tasks []*task
}

// NewScheduler 创建新的调度器
func NewScheduler() *Scheduler {
	return &Scheduler{
		tasks: make([]*task, 0),
	}
}

// AddTask 添加新任务
func (s *Scheduler) AddTask(id int, name string, execTime time.Time, job func(ctx context.Context)) {
	s.mu.Lock()
	defer s.mu.Unlock()
	ctx, cancel := context.WithCancel(context.Background())
	task := &task{
		ID:       id,
		Name:     name,
		ExecTime: execTime,
		Job:      job,
		Ctx:      ctx,
		Cancel:   cancel,
	}
	s.tasks = append(s.tasks, task)
	// fmt.Printf("Task %d (%s) added, scheduled at %v\n", id, name, execTime)
}

// CancelTask 取消指定任务
func (s *Scheduler) CancelTask(id int) {
	s.mu.Lock()
	defer s.mu.Unlock()
	for _, task := range s.tasks {
		if task.ID == id {
			task.Cancel()
			// fmt.Printf("Task %d (%s) canceled\n", task.ID, task.Name)
			return
		}
	}
	// fmt.Printf("Task %d not found for cancellation\n", id)
}

// removeTask 移除指定任务
func (s *Scheduler) removeTask(id int) {
	s.mu.Lock()
	defer s.mu.Unlock()
	for i, task := range s.tasks {
		if task.ID == id {
			s.tasks = append(s.tasks[:i], s.tasks[i+1:]...)
			// fmt.Printf("Task %d removed\n", id)
			return
		}
	}
}

// Run 启动调度器
func (s *Scheduler) Run(ctx context.Context) {
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			// fmt.Println("Scheduler stopped:", ctx.Err())
			return
		case <-ticker.C:
			s.mu.Lock()
			now := time.Now()
			var toRemove []int
			for _, task := range s.tasks {
				if now.After(task.ExecTime) || now.Equal(task.ExecTime) {
					if task.Ctx.Err() == nil {
						// fmt.Printf("Executing task %d (%s) at %v\n", task.ID, task.Name, now)
						task.Job(task.Ctx) // 同步执行任务
						toRemove = append(toRemove, task.ID)
					} else {
						// fmt.Printf("Task %d (%s) was canceled, removing\n", task.ID, task.Name)
						toRemove = append(toRemove, task.ID)
					}
				}
			}
			s.mu.Unlock()

			// 移除已执行或取消的任务
			for _, id := range toRemove {
				s.removeTask(id)
			}
		}
	}
}
