# Idea

This project serves the same purpose as (testcontainers)[https://www.testcontainers.org/] but tries to do so as efficiently as possible by levariging the build system's capabilities, in this case Bazel, while being language agnostic and as fast as possible.

# How is this trying to be faster than testcontainers?

Traditionally builds and tests can be made faster by parallelising the work and reducing the amount of work needed. This gives us the following techniques that can be applied to any kind of build or test workload to potentially speed it up:
- Parallelism techniques:
    - High action granularity, where an action is the build system's most granular unit of execution (like invoking a test suite or a compiler invocation). This is what makes it possible to truly use the abilities of a build system to work on multiple things at the same time.
    - High sub-action parallelism, this is parallelism outside of the control of the build system, aka parallelism coming from the underlying tools used in a build or a test. Though at first glance this seems beneficial, it has a major issue. Because the build system is unaware of said parallelism, it will consider everything as a single action and thus an higly concurrent tool may cause other actions to starve for resources and potentially make the whole build slower. This becomes even worse if said tool then spawns more executions.
    - Remote Build Execution, this is an API that allows scheduling arbitrary workloads on a remote system that may or may not be distributed. This system may be bigger than your local machine and if distributed may provide resource pools that are exponentially bigger than what any local setup may ever be able to provide.
- Do-less-work techniques:
    - Caching, probably one of the oldest techniques in CS to avoid doing extra work, if you know the outcome of an execution just return that rather than re-executing the work again
    - Remote Caching, same as above but on a remote system
    - Daemon processes, this helps saving bootstraping time by reusing an already running process rather than creating a new one each and every time. This is extremely useful for tools that do a lot of initial preparation work like tools developed in a language running on a relatively slow VM (like JVM, Node, etc.) or tools that are spawning startup heavy processes like databases. In fact, testcontainers already supports this, but this technique does not play well with Remote Build Execution, since it requires to keep a running process and the Remote Build Execution protocol does not support that. It is also potentially counterproductive for parallelism as it may cause multiple actions to wait on a shared resource to be available.

With that in mind, let's start by analysing what we want in terms of parallelisation. We surely want this system to have as much granularity at the action level and as little granularity below what bazel can control, so not to starve the rest of the build. We also want to support well Remote Build Execution as it will allow this library to be automatically way more scalable.

This implies that we are not going to be able to use Daemon processes but we will be using caching and remote caching for everything as much as possible. Matter of fact, caching can be used to introduce a new technique to reduce startup time of an action: Caching container checkpoints.

Docker supports something called checkpoints, a file based snapshot of the current state of the memory of a container. With this we can start a container and do all the initial heavy lifting and then capture a snapshot that will be usable as any other generated source in the build system. This also makes it usable by Remote Execution and because the same checkpoint + image combination will probably be used by multiple tests (imagine a checkpointed warm mysql container) it will also be cacheable.

# Problems with the approach

Because checkpoints can only be resumed and not used to start a fresh new instance of a container, you only get access to a subset of the features supported by `docker container start`. The most notable of these features are:
- Volume bindings
- Port forwarding
- St(in|out|err) attaching

## Port randomization

Though it is impossible to expose a port after a container is started, it is possible to isolate at the networking layer, each instance of a checkpointed container always exposing the same port mapping, which thanks to network isolation won't collide. This is helpful because networks can be attached at runtime so if by default a container is started without a network (network driver = none), then checkpointed and later resumed we can still attach net network interfaces to it on which we can then our own abstraction to forward ports.

## Volume bindings, file sharing and DB migrations

Though volume bindings seem impossible it is possible to copy from/to a running container. Regardless this may cause the default strategies of containerised DBs rather unusable since those normally use volume binding to bind migrations scripts to a specific location. Nothing stops us though from exposing a migration API that is a bit more explicit, but it will make it harder to create custom DB containers.

# APIs and protocol