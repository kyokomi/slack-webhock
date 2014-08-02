slack-webhock

[![wercker status](https://app.wercker.com/status/877657cf5fc70edf3821f7c56540c211/m "wercker status")](https://app.wercker.com/project/bykey/877657cf5fc70edf3821f7c56540c211)

=============

slack webhock golang

## Demo

### Input

Created Gitlab new Issue.

![](https://raw.githubusercontent.com/kyokomi/slack-webhock/master/manual/demo1.png)

### Output

PostMessage Slack.



## Install

## Local

`~/.bash_profile`

`{}`は、自分の設定

```sh
# personal-dev
export GITLAB_TOKEN={YOUR_GITLAB_TOKEN}
export SLACK_TOKEN={YOUR_SLACK_TOKEN}
export SLACK_CHANNEL={TARGET_SLACK_CHANNEL}
```

### wercker

![](https://raw.githubusercontent.com/kyokomi/slack-webhock/master/manual/wercker_setup.png)

### heroku

![](https://raw.githubusercontent.com/kyokomi/slack-webhock/master/manual/heroku_setup.png)

### gitlab

![](https://raw.githubusercontent.com/kyokomi/slack-webhock/master/manual/gitlab_setup.png)

