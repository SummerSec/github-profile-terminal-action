# github-profile-terminal-action

在原来项目[liamg/github-profile-terminal-action](https://github.com/liamg/github-profile-terminal-action)的基础上改了一点代码，保持了自己的风格。



This is the code that generates my [profile README.md](https://github.com/SummerSec), including the terminal gif.

You can use it to automatically create your GitHub profile README by adding a GitHub action to your profile repository.

Here's an example config that updates every 4 hours:

```yaml
name: auto-update

on:
  workflow_dispatch:
  schedule:
    - cron:  0 */4 * * *

jobs:
  build:
    name: auto update
    runs-on: ubuntu-latest

    steps:
    - uses: actions/checkout@master
    - uses: SummerSec/github-profile-terminal-action@main
      with:
        feed_url: https://sumsec.me/resources/sitemap.xml
        twitter_username: SecSummers
        theme: dark
        token: ${{ secrets.GITHUB_TOKEN }}
```
