# Slack Status CLI
Lightning fast Slack status updates!⚡

<details>
<summary><strong>  See it in action :fire: </strong></summary>
<img src="https://user-images.githubusercontent.com/13623913/187346619-213125ef-6ece-4a73-9c8b-5792a240ca7c.gif" width="600" />
</details>


## Setting things up :hammer:
- [Create a new Slack app](https://api.slack.com/apps)
- Select from an **app manifest**
- Select your workspace
- Paste the following YAML app manifest:

```yaml
display_information:
    name: slack-status-cli
    description: Lightning fast Slack status updates!
oauth_config:
    scopes:
        user:
        - dnd:write
        - emoji:read
        - users.profile:read
        - users.profile:write
        - users:read
        - users:write
```

- Create :rocket:
- Install to workspace
- Go to OAuth and permissions
- Copy User Auth Token

### Download :inbox_tray:

> Get the latest version from [**Releases**](https://github.com/yankeexe/slack-status-cli/releases)

Make it executable 

```sh
chmod +x st_<version_platform_arch>
```

Move to your `$PATH`

```sh
mv st_<version_platform_arch> <$PATH>/st
```

---
### Create new profile :bear:

```bash
st profile --create  # shorthand -c
```

<details>
<summary><strong>Add your profile name and OAuth token:</strong></summary>
<img width="469" alt="image" src="https://user-images.githubusercontent.com/13623913/187342281-96d95ef2-f69b-4587-83df-034f6e07477d.png">
</details>

## Usage :rainbow:

### Add new status :headphones:

Add your status, duration and emoji

```bash
st add
```


### Select status :eyes:

```bash
st set
```


### Set yourself away :sleeping:

```bash
st away
```

**Enable Do Not Disturb (DND) while away** :mask:

```bash
st away --dnd <time duration>
```

### Set yourself active :surfer:

Remove any status, or DND settings

```bash
st active
```

### Profile management :wrench:

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

### Setting time duration :alarm_clock:

Valid durations for time includes:
> minutes, hours or days.
DEFAULTS to minutes

**NOTE: use single or double quotes around the time duration values.**

OPTIONS for the duration:

- minute: m, min, mins       :: Example: "10 m", "10 mins", 10minute, "10 minutes"

- hour:   h, hr, hour, hours :: Example: "1 h", 1hr, "1 hour", "1 hours"

- day:    d, day, days       :: Example: 2d, "2 day", 2days
