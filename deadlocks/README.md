# Deadlocks

When the program gets blocked due to the threads are waiting for each other to release.
So they end up in an infinite loop of waiting, they are blocking each other

There is no a way to truly omit this. But there're techniques to prevent this to happen

## Techniques to prevent deadlocks

### Resource Hierarchy

In this technique we need to define an order in which the locks will be acquire.
For example, if always needs to adquire the lower one first.
So, if it comes to the bucle, it will try to adquire the first one and since
it's already locked it will wait.

### Arbitrator

A piece of code that will choose which thread is going to use which resources.

Pseudo code of the arbiter

lock(mutex)
while (selected resources are not free)
    wait(condition) -> conditionl variable
mark(all resources as busy)
unlock(mutex)

Pseudo code to unlock the resources

lock(mutex)
mark(all resources as free)
broadcast
unlock(mutex)
