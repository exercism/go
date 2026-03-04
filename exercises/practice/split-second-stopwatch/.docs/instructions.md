# Instructions

Your task is to build a stopwatch to keep precise track of lap times.

The stopwatch uses four commands (start, stop, lap, and reset) to keep track of:

1. The current lap's tracked time
2. Previously recorded lap times

What commands can be used depends on which state the stopwatch is in:

1. Ready: initial state
2. Running: tracking time
3. Stopped: not tracking time

| Command | Begin state | End state | Effect                                                   |
| ------- | ----------- | --------- | -------------------------------------------------------- |
| Start   | Ready       | Running   | Start tracking time                                      |
| Start   | Stopped     | Running   | Resume tracking time                                     |
| Stop    | Running     | Stopped   | Stop tracking time                                       |
| Lap     | Running     | Running   | Add current lap to previous laps, then reset current lap |
| Reset   | Stopped     | Ready     | Reset current lap and clear previous laps                |
