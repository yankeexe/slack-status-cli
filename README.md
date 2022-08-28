# Slack Status CLI
Lightning fast Slack status updates!⚡


## Setting things up


### Get Slack token

### Create new profile

```bash
st profile --create  # shorthand -c
```

## Usage

### Add new status

Add your status, duration and emoji

```bash
st add
```


### Select status

```bash
st set
```


### Set yourself away

```bash
st away
```

**Enable Do Not Disturb (DND) while away**

```bash
st away --dnd <time duration>
```

### Set yourself active

Remove any status, or DND settings

```bash
st active
```

### Profile management

**Set default profile**
If you have multiple slack profiles, select default profile using:

```bash
st profile --default  # shorthand -d
```

**Manage profile: change name, token, delete or update status**

Uses default profile:

```bash
st profile --manage  # shorthand -m
```

Select profile to manage:

```bash
st profile --manage --select # shorthand -m -s
```

**Show current/default profile**

```bash
st profile --show
```

### Setting time duration

Valid durations for time includes:
> minutes, hours or days.
DEFAULTS to minutes

NOTE: use single or double quotes around the value if contains space.

OPTIONS for the duration:

- minute: m, min, mins       :: Example: "10 m", "10 mins", 10minute, "10 minutes"

- hour:   h, hr, hour, hours :: Example: "1 h", 1hr, "1 hour", "1 hours"

- day:    d, day, days       :: Example: 2d, "2 day", 2days
