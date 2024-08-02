- reset to default value https://github.com/cocopon/tweakpane/issues/508
- get latest value: once you get the event from sse, /get backend again to get the latest value (otherwise you'd get default value)
- restore (checkbox, defaults to false): rather than /get to get latest value we'd store last item in tweaked-values from local storage (if present) and we'll check if a value is present there and we will use the value from that place. Any value restored is not marked as tweaked.

- qlocktwo
- text shimmer
- coverflow
- rope physics (twitter)
- `A*` algo (hackernews, query term: raylib)
