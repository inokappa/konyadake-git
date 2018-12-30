# konyadake-git [![CircleCI](https://circleci.com/gh/inokappa/konyadake-git.svg?style=svg)](https://circleci.com/gh/inokappa/konyadake-git)

## これなに

<iframe width="560" height="315" src="https://www.youtube.com/embed/vhWaC9i7Eko" frameborder="0" allow="accelerometer; autoplay; encrypted-media; gyroscope; picture-in-picture" allowfullscreen></iframe>

今夜だけきっと, 今夜だけ git...

* [go-git](https://github.com/src-d/go-git) の勉強の為に作った
* `git clone` と `git checkout` が出来る

## 使い方

### clone

```sh
konyadake -repo=git@github.com:your-name/example.git
```

```sh
konyadake -repo=https://github.com/your-name/example.git
```

```sh
konyadake -repo=user@user.git.backlog.jp:YOUR-PJ/example.git
```

```sh
konyadake -repo=https://user.git.backlog.jp/YOUR-PJ/example.git -username=username
```

### checkout

```sh
konyadake -repo=git@github.com:your-name/example.git -branch=develop
```

## example

```sh
$ konyadake -repo=git@github.com:inokappa/circleci-docker-test-and-build.git -branch=develop
✅  対象リポジトリ git@github.com:inokappa/circleci-docker-test-and-build.git を circleci-docker-test-and-build に clone します.
Enumerating objects: 70, done.
Counting objects: 100% (70/70), done.
Compressing objects: 100% (34/34), done.
Total 70 (delta 24), reused 64 (delta 18), pack-reused 0
✅  対象リポジトリの clone に成功しました.
✅  対象ブランチの checkout に成功しました.
```
