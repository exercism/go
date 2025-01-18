# Instructions append

## Implementation Notes

Tests are separated into 3 steps.

Run all tests with `go test` or run specific tests with the -tags option.

Examples:

```bash
go test                      # run all tests
go test -tags step1          # run just step 1 tests.
go test -tags 'step1 step2'  # run step1 and step2 tests
```

You are given the source file defs.go which defines a number of things
the test program requires.  It is organized into three sections by step.

## Step 1

To complete step 1 you will define Right, Left, Advance, N, S, E, W,
and Dir.String.  Complete step 1 before moving on to step 2.

## Step 2

For step 1 you implemented robot movements, but it's not much of a simulation.
For example where in the source code is "the robot"?  Where is "the grid"?
Where are the computations that turn robot actions into grid positions,
in the robot, or in the grid?  The physical world is different.

Step 2 introduces a "room."  It seems a small addition, but we'll make
big changes to clarify the roles of "room", "robot", and "test program"
and begin to clarify the physics of the simulation.  You will define Room
and Robot as functions which the test program "brings into existence" by
launching them as goroutines.  Information moves between test program,
robot, and room over Go channels.

Think of Room as a "physics engine," something that models and simulates
a physical room with walls and a robot.  It should somehow model the
coordinate space of the room, the location of the robot and the walls,
and ensure for example that the robot doesn't walk through walls.
We want Robot to be an agent that performs actions, but we want Room to
maintain a coherent truth.

The test program creates the channels and starts both Room and Robot.
The test program then sends commands to Robot.  When it is done sending
commands, it closes the command channel.  Robot must accept commands and
inform Room of actions it is attempting.  When it senses the command channel
closing, it must shut itself down.  The room must interpret the physical
consequences of the robot actions.  When it senses the robot shutting down,
it sends a final report back to the test program, telling the robot's final
position and direction.

## Step 3

Step 3 has three major changes:

*  Robots run scripts rather than respond to individual commands.
*  A log channel allows robots and the room to log messages.
*  The room allows multiple robots to exist and operate concurrently.

For the final position report sent from Room3, you can return the same slice
received from the robots channel, just with updated positions and directions.

Messages must be sent on the log channel for
*  A robot without a name
*  Duplicate robot names
*  Robots placed at the same place
*  A robot placed outside of the room
*  An undefined command in a script
*  An action from an unknown robot
*  A robot attempting to advance into a wall
*  A robot attempting to advance into another robot
