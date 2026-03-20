# tjann - Timewarrior JSON Annotations

Use timewarrior annotations as JSON stores.

## Usage

```[bash]
# create a time
timew start
sleep 3
timew stop

# fail to read (there is nothing on there)
tjann get @1 myuuid

# write
tjann set @1 myuuid $(uuidgen)

# reading will succeed
tjann get @1 myuuid

# delete
tjann unset @1 myuuid

# fail to read (again, there is nothing left here)
tjann get @1 myuuid
```

Use ```timew export @1``` if you want to inspect what is happening.
