# I Ain't Waiting

## Intro

IAW is a tool designed to facility debugging of long-term stability problems. Although, designed for that kind of problems, it can also work for other bugs.

IAW consists of 3 main layers.
1. Wrapper: A wrapper around the binary in question, used to capture output, send input,...
2. Trigers: A trigger that starts the capturing of data.
3. Handlers: Once a trigger has activated, all handlers associated with that trigger is called.

So a standard IAW flow looks like:

Wrapper -> Starts binary
Triggers -> Added to Wrapper. For example a string trigger on the stdout or sterr of the binary
Handlers -> Added to Wrapper. A file trigger, capture the output of a file at the time of trigger. This is mainly useful for system files. Such as temperature readings.

## Types of triggers
|Name|Description|
|---|---|
|ioreader| Trigger on a string found in an ioreader's input.|

## Types of handlers
|Name|Description|
|---|---|
|filehandler| Capture the output of a given file.|
