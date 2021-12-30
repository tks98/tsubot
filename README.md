# tsubot

tsubot is a discord bot with user role automation and osu! score integration.

## Installation and Usage

### Running locally

You will need to install Go and create the config.yaml file

```bash
$ go build cmd/main.go
$ ./main -c=/path/to/config.yaml
```

### Using Docker

You can also build and run the bot using the provided Dockerfile

```bash
$ docker build -t myCoolBot:1.0 .
$ docker run myCoolBot:1.0
```

### Example Configuration File

```yaml
apiKeys:
  discord: discord_api_key
  osu:
    clientID: osu_client_id
    clientSecret: osu_client_secret

guildID: server_guild_id

# roles you want users in the server to be able to add and remove to themselves
allowedRoles:
  - Pro-Players
  - Regulars
  - Streamers
  - Mappers
  - Developers
  - Youtubers
  - Osu-Players

# list of commands that users on the server can issue the bot
commands:
  - welcome
  - verify
  - ping
  - help
  - choose
  - roles
  - remove
  - rank
  - info
 ```

## Commands

### Role automation

- Available roles can be found with !roles
- Roles can be chosen with !choose role
- Roles can be removed with !remove role
- Pro-Player role can be added with !verfy osu-profile-url

### osu!

- A user's basic info can be queried with !info username
- A user's top score info can be queried with !info -t username
- A user's most recent score info can be queried with !info -r username
- A user's most recent 1st place score can be queried with !info -f username

## Usage Examples
### !roles
![alt-text](https://cdn.discordapp.com/attachments/611191473601511434/926006474336006184/Screen_Shot_2021-12-30_at_12.58.54_AM.png)
### !info -t Cookiezi
![alt text](https://cdn.discordapp.com/attachments/611191473601511434/926005060046053396/Screen_Shot_2021-12-30_at_12.52.21_AM.png)
### !info Woey
![alt-text](https://cdn.discordapp.com/attachments/611191473601511434/926005309045080064/Screen_Shot_2021-12-30_at_12.54.18_AM.png)

