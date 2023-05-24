# Challenge 3 - Broadcast

## Thoughts

### Similarity to past projects 
Seems similar to a gossip protocol over a P2P overlay network project I'd done in CMU Networks course.
This was used to communicate changes to the network topology quickly, so each node could calculate the shortest route to every other node for itself.
I remember being surprised how quickly information spread across the cluster, though I had to take care to avoid infinite loops.

### Challenge 3a - Single-Node Broadcast

#### Golang quirks

Spent some time wrangling with Go syntax on this one. Learnt about `interface{}` and learnt that the default JSON tries to marshal and unmarshal `any` types to `float64` instead of `int` if it encounters a number.
Also learnt that Go treats unused variables as a compilation error and the `_ = ...` workaround

#### Testing & debugging

Wish there was a cleaner way for quick testing other than the Maelstrom workload they've given us. I did use the online Go playground for JSON testing but for handler or Maelstrom-related code, we have to run the given workload.
One hack I found was to decrease the time period and the request rate as a sanity check - for faster code-edit-run cycles.

### Solution - 3a

Nothing much to do here other than understand the API we need to implement and throw in a few internal datastructures.
I did find it annoying that Go doesn't have a hashset. The workaround seemed more cumbersome than just using a `map[int]bool` so that's what I did for broadcast IDs.
I'm using a `map[string]interface{}` for network topology though it's technically a `map[string][]string`. 

I'm currently not using any mutexes or RWLocks to protect these datastructures. This might be required in the future, since Go maps are not thread-safe.

### Challenge 3b - Multi-Node Broadcast

Now that we have single-node broadcast working, let's try to see how to make multi-node broadcast without sending the entire dataset on every broadcast message.
What makes most sense is to avoid having state that tracks what has been broadcasted. Rather that that, we should every node should forward broadcasts as soon as it receives them, trusting the gossip protocol to avoid infinite loops.
Note that infinite loops in a gossip protocol is more than a performance issue, it causes a positive feedback loop that leads to message flooding and completely congests the network, eventually even taking down the nodes. So this is a correctness issue, not just a performance one.

## Solution


## Random ideas
