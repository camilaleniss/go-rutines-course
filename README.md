# go-rutines-course

Notes from the Udemy Course Mastering Multithreading in Go: <https://www.udemy.com/course/multithreading-in-go-lang>

## Intro

Thread -> abstraction that allows us to perform parallel computation

Parallel computing -> art of taking one problem, put it in different processing unit to solve it faster

Amdahl's law -> ratio between parallel and sequential parts of a process that will dictate the speed up that we can achieve

Gustafson's law -> ratio between the problem size and the number of processors we can put. The relation it's linear if we increase the problem size, no matter how big is the sequential part of the process

Parallel -> each unit in the processing queue goes to different processor sequentially
Concurrent -> each unit have an amount of times in the processor and goes back to the processing queue

## Creating and using threads

Processes -> isolation in the resources consumed. But consume more memory, since each process needs its own

Threads -> used to solve the problem with processes. No allocate new memory for it

Green threads -> tries to solve the isolation problem. Now we give enough time to each process to execute.

- But then we switch to another (context switch) and the time we spent doing it is the context switch overhead.
- If more processors -> more context switch time.
- User level thread. The program decides which thread to execute != Kernel level threads. The program don't know if you are waiting for a certain resource.
- Golang use both: Have Kernel thread in each CPU and each CPU uses multiple green level threads. Whenever a green thread tries to do an operation when you have to wait, takes all the other threads that don't have to wait to be processed in the Kernel

Thread can communicate (Intern-process communication) using two ways:

1. Message passing: like channels
2. Shared memory: like if you have a white board, and all the threads can see what you are writing. And also write.

To create threads in go we use:

´´´
go funcName()
´´´
