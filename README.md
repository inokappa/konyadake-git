# konyadake-git [![CircleCI](https://circleci.com/gh/inokappa/konyadake-git.svg?style=svg)](https://circleci.com/gh/inokappa/konyadake-git)

## これなに

* [go-git](https://github.com/src-d/go-git) の勉強の為に作った
* `git clone` と `git checkout` が出来る

## 使い方

### clone

```sh
konyadake -repo=git@github.com:inokappa/example.git
```

```sh
konyadake -repo=https://github.com/inokappa/example.git
```

```sh
konyadake -repo=user@user.git.backlog.jp:YOUR-PJ/example.git
```

```sh
konyadake -repo=https://user.git.backlog.jp/YOUR-PJ/example.git -username=username
```
