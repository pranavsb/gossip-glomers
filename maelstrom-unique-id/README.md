# Challenge 2 - Unique ID generation

## Thoughts

From the challenge:  "IDs may be of any type--strings, booleans, integers, floats, arrays, etc."
Though IDs can be of any type, let us stick with integer since they are easier to reason about (and can be converted to string or array if required). Booleans can be ruled out cuz duh, and floats are black magic to a computer so let's stay away as any sane person would.

Since our system has to be partition tolerant, we have to leverage something unique that every node has,which it can obtain without depending on any other node: *its ID!*

We trust the underlying system to return unique IDs on node init (a cursory look at the `maelstrom.NewNode` code seems to do so). We also have  to assume that adding new nodes to the system (which is not in scope for the provided Jepsen tests) relies on a central service that always returns unique IDs (eg. global monotonically increasing  int)

So the node's ID forms one component of the unique ID we can return to the client. For the second (and last) part, we can go with two approaches:
* randomly generated large int
* UNIX timestamp

Though timestamp should work well in practice, as long as your system (OS, programming language) relies on a nanosecond precision timestamp like [utimensat](https://linux.die.net/man/2/utimensat). 
The Jepsen test uses 3 nodes at 1000 requests/second rate, so microsecond should be fine as well. 

This simplifies things as random number generation might have taken a lot more time, and we would have to store all past generated numbers in memory to guarantee that it hasn't been generated before. Though we could make the length of the random int long enough to get around it.

## Solution

I will use: node ID + timestamp as the unique ID. Node ID is already available as a string, so we can convert timestamp to a string and concat it.

## Random ideas 
Another thought, if we were restricted to microsecond/millisecond timestamps for whatever reason, or we did not trust the resolution of our timestamp to be unique, we could maintain the past generated timestamp in an in-memory data structure (a hashmap makes sense) to make sure we aren't duplicating it.
Then, a background thread can periodically perform cleanup (since we expect timestamps to be monotonically increasing) of the map - deleting old enough timestamps.