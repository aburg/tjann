# tjann - Timewarrior JSON Annotations

Use timewarrior annotations as JSON stores.

## Usage

```[bash]
tjann set @1 '{"billable":true}'
tjann set @1 '{"billed":true}'
tjann get @1
tjann query '{"billed": true}'
```

## How this works

Not at all...
