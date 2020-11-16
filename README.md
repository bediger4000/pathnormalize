# Daily Coding Problem: Problem #713 [Medium]

This problem was asked by Quora.

Given an absolute pathname that may have . or .. as part of it,
return the shortest standardized path.

For example,
given "/usr/bin/../bin/./scripts/../",
return "/usr/bin/".

## Analysis

This may be poorly specified.
Is the input path ASCII (as traditional for Unix) or UTF-8?
I'm not sure this is an important difference, but still in 2020,
you should probably say what you're assuming.

I implemented this problem as an array traverse of path components
(non-zero-length strings between '/' bytes.

My code keeps the normalized path as a stack of non-zero-length
strings which will comprise the path components of desired string.
Each non-zero-length path component causes an action:

* "." gets ignored. The normalized path stays the same.
* ".." causes the code to pop an element off the normalized path stack.
This represents the parent directory of the current path.
This code has the only real trick: a filesystem has a root directory,
"/", which has itself as a parent, and therefore does not cause
a stack pop.
* Any other string gets pushed on the normalized path stack.

The problem specification asks for a leading and trailing slash ('/')
character.
My code checks for a non-zero length normalized path,
since a normalized path of "//" makes no sense.

## Interview Analysis

I'd say this isn't a bad interview question.
It clearly requires a loop over path components,
and some decision making based on each component
as it's encountered.

It does require a bit of insight,
in that a Unix-style path works like a stack when you handle ".."
(parent directory) components.

If I were interviewing candidates,
after the basic algorithm,
I'd see what they proposed as test cases.
Something like "///" or "/././." or "/../../.." seem like good cases,
as does the example path.

I'm really uncertain of the "medium" rating.
I'd actually call it an "easy".
