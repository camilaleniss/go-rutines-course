# Barriers

A tool to syncronize different processes together
Like an start line

To create a Barrier

Let's imagine we have these attributes:

- total
- count
- mutex
- condition variable

And we need this as a constructor

New(total) # where total is the number of threads for the barrier

And then when a thread finish their execution needs to call the wait()

count-=1
count > 0 ?
    cv.wait()

When the latest one call wait()

count == 0
    count = total
    cv.broadcast()
