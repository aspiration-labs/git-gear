# Git Gear: Hooks, ladders, picks and shovels for git, jira, ...

## Download release or build
- Releases: https://github.com/aspiration-labs/git-gear/releases
- Build:
```
git clone git@github.com:aspiration-labs/git-gear.git
cd git-gear
go build
```

## Install
```
mkdir -p $HOME/bin
cp git-gear $HOME/bin
```

## Commands and flags
```
git-gear --help
```

## Use the commit message hook
### Option 1: In a single git repo
In `/path/to/git/repo/.git/hooks/commit-msg` add
```
#!/bin/sh

$HOME/bin/git-gear jira commitCheck "$1" --jiraserver https://username:token@jira.atlassian.net
```
Make sure `commit-msg` is executable
```
chmod +x /path/to/git/repo/.git/hooks/commit-msg
```

### Option 2: For all my repos
```
mkdir ~/.githooks
git config --global core.hooksPath ~/.githooks/
```
and in `~/.githooks/commit-msg` add
```
#!/bin/sh

# check and run local repo hook
if [ -e ./.git/hooks/commit-msg ]; then
    ./.git/hooks/commit-msg "$@" || exit $?
fi

$HOME/bin/git-gear jira commitCheck "$1" --jiraserver https://username:token@jira.atlassian.net
```
Make sure `commit-msg` is executable
```
chmod +x ~/.githooks/commit-msg
```

# FAQ

## It's a pain to put the jiraserver option and spam my token everywhere. Is there a better way?

Instead of --jiraserver option, you can use `$HOME/.git-gear.yml` like
```
jiraserver: https://username:token@jira.atlassian.net
```

## Where do I get a Jira API token?

For Atlassian cloud based Jira, https://id.atlassian.com/manage/api-tokens
