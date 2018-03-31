# _murex_ Language Guide

## Command reference: event

> Event driven programming for shell scripts

### Description

Create or destroy an event interrupt

### Usage

    event: event-type name=interrupt { code block }

    !event: event-type name

### Examples

Create an event:

    event: afterSecondsElapsed autoquit=60 {
        out "You're 60 secound timeout has elapsed. Quitting murex"
        exit 1
    }

Destory an event:

    !event afterSecondsElapsed autoquit

### Detail

### Synonyms

* !event

### See also

* `runtime`