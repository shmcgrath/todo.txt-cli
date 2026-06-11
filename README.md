# [![todo.txt-cli](http://todotxt.org/images/todotxt_logo_2012.png)][website]

> A simple and extensible shell script for managing your todo.txt file.

![CI](https://github.com/todotxt/todo.txt-cli/workflows/CI/badge.svg)
[![GitHub issues](https://img.shields.io/github/issues/todotxt/todo.txt-cli.svg)](https://github.com/todotxt/todo.txt-cli/issues)
[![GitHub forks](https://img.shields.io/github/forks/todotxt/todo.txt-cli.svg)](https://github.com/todotxt/todo.txt-cli/network)
[![GitHub stars](https://img.shields.io/github/stars/todotxt/todo.txt-cli.svg)](https://github.com/todotxt/todo.txt-cli/stargazers)
[![GitHub license](https://img.shields.io/github/license/todotxt/todo.txt-cli.svg)](https://raw.githubusercontent.com/todotxt/todo.txt-cli/master/LICENSE)
[![Gitter](https://badges.gitter.im/join_chat.svg)](https://gitter.im/todotxt/todo.txt-cli)

![gif](./.github/example.gif)

*Read our [contributing guide][CONTRIBUTING] if you're looking to contribute (issues/PRs/etc).*


## Installation

### Download
Download the latest stable [release][release] for use on your desktop or server.

### OS X / macOS

```shell
brew install todo-txt

cp -n $(brew --prefix)/opt/todo-txt/todo.cfg ~/.todo.cfg
```

**Note**: The `-n` flag for `cp` makes sure you do not overwrite an existing file.

### Linux

#### From command line

```shell
make
make install
make test
```

*NOTE:* Makefile defaults to several default paths for installed files. Adjust to your system:

- `INSTALL_DIR`: PATH for executables (default `/usr/local/bin`)
- `CONFIG_DIR`: PATH for the `todo/config` configuration template (default `/usr/local/etc`)
- `BASH_COMPLETION`: PATH for autocompletion scripts (default to `/usr/local/share/bash-completion/completions`)

```shell
# Note: Showcasing config overrides for legacy locations; NOT recommended!
make install CONFIG_DIR=/etc INSTALL_DIR=/usr/bin BASH_COMPLETION=/etc/bash_completion.d
```

#### Arch Linux (AUR)

https://aur.archlinux.org/packages/todotxt/


## Configuration

No configuration is required; however, most users tweak the default settings (e.g. relocating the todo.txt directory to a subdirectory of the user's home directory, or onto a cloud drive (via the `TODO_DIR` variable)), modify the colors, add additional highlighting of projects, contexts, dates, and so on. A configuration template with a commented-out list of all available options is included.
It is recommended to _copy_ that template into one of the locations listed by `todo.sh help` on `-d CONFIG_FILE`, even if it is installed in the global configuration location (`/etc/todo/config`).

## Usage
```shell
todo.sh [-fhpantvV] [-d todo_config] action [task_number] [task_description]
```

For example, to add a todo item, you can do:

```shell
todo.sh add "THING I NEED TO DO +project @context"
```
### `replace`
Replaces task on line NR with UPDATED TODO.

```shell
todo.sh replace NR "UPDATED TODO"
```
### `report`
Adds the number of open tasks and done tasks to report.txt.

```shell
todo.sh report
```

Read about all the possible commands in the [USAGE][USAGE] file.


## Release History

See [CHANGELOG.md][CHANGELOG]


## Support

- [Github Discussions](https://github.com/todotxt/todo.txt-cli/discussions)
- [Stack Overflow](https://stackoverflow.com/questions/tagged/todotxt)
- [Twitter](https://twitter.com/todotxt)


## Code of Conduct

[Contributor Code of Conduct][CODE_OF_CONDUCT]. By participating in this project you agree to abide by its terms.

## Contributing

We welcome all contributions. First read our [Contributor Code of Conduct][CODE_OF_CONDUCT] and then get started [contributing][CONTRIBUTING].

## License

GNU General Public License v3.0 © [todo.txt org][github]



[release]: https://github.com/todotxt/todo.txt-cli/releases
[website]: http://todotxt.org/
[github]: https://github.com/todotxt
[USAGE]: ./USAGE.md
[CHANGELOG]: ./CHANGELOG.md
[CODE_OF_CONDUCT]: .github/CODE_OF_CONDUCT.md
[CONTRIBUTING]: .github/CONTRIBUTING.md

## todo.txt README.md from [GitHub](https://github.com/todotxt/todo.txt)
## todo.txt format
[![Gitter](https://img.shields.io/gitter/room/todotxt/todotxt.svg)](https://gitter.im/todotxt/todotxt)

A complete primer on the whys and hows of todo.txt.

The first and most important rule of todo.txt:

> A single line in your todo.txt text file represents a single task.


### Why plain text?

Plain text is software and operating system agnostic. It's searchable, portable, lightweight, and easily manipulated. It's unstructured. It works when someone else's web server is down or your Outlook .PST file is corrupt. There's no exporting and importing, no databases or tags or flags or stars or prioritizing or _insert company name here_-induced rules on what you can and can't do with it.


### The 3 axes of an effective todo list

Using special notation in todo.txt, you can create a list that's sliceable by 3 key axes.


#### Priority
Your todo list should be able to tell you what's the next most important thing for you to get done - either by project or by context or overall. You can optionally assign tasks a priority that'll bubble them up to the top of the list.


#### Project
The only way to move a big project forward is to tackle a small subtask associated with it. Your `todo.txt` should be able to list out all the tasks specific to a project.

In order to move along a project like "Cleaning out the garage," my task list should give me the next logical action to take in order to move that project along. "Clean out the garage" isn't a good todo item; but "Call Goodwill to schedule pickup" in the "Clean out garage" project is.


#### Context
[Getting Things Done](https://en.wikipedia.org/wiki/Getting_Things_Done) author David Allen suggests splitting up your task lists by context - ie, the place and situation where you'll work on the job. Messages that you need to send go in the `@email` context; calls to be made `@phone`, household projects `@home`.

That way, when you've got a few minutes in the car with your cell phone, you can easily check your `@phone` tasks and make a call or two while you have the opportunity.

This is all possible inside `todo.txt`.



### `todo.txt` format rules

<img src="./description.svg" width="100%" height="500">

Your `todo.txt` is a plain text file. To take advantage of structured task metadata like priority, projects, context, creation, and completion date, there are a few simple but flexible file format rules.

Philosophically, the `todo.txt` file format has two goals:

- The file contents should be human-readable without requiring any tools other than a plain text viewer or editor.
- A user can manipulate the file contents in a plain text editor in sensible, expected ways. For example, a text editor that can sort lines alphabetically should be able to sort your task list in a meaningful way.

These two goals are why, for example, lines start with priority and/or dates, so that they are easily sorted by priority or time, and completed items are marked with an `x`, which both sorts at the bottom of an alphabetical list and looks like a filled-in checkbox.

Here are the rest.


### Incomplete Tasks: 3 Format Rules

The beauty of todo.txt is that it's completely unstructured; the fields you can attach to each task are only limited by your imagination. To get started, use special notation to indicate task context (e.g. `@phone` ), project (e.g. `+GarageSale` ) and priority (e.g. `(A)` ).

A todo.txt file might look like the following:

```
(A) Thank Mom for the meatballs @phone
(B) Schedule Goodwill pickup +GarageSale @phone
Post signs around the neighborhood +GarageSale
@GroceryStore pies
```

A search and filter for the `@phone` contextual items would output:

```
(A) Thank Mom for the meatballs @phone
(B) Schedule Goodwill pickup +GarageSale @phone
```

To just see the `+GarageSale` project items would output:

```
(B) Schedule Goodwill pickup +GarageSale @phone
Post signs around the neighborhood +GarageSale
```

There are three formatting rules for current todo's.

#### Rule 1: If priority exists, it ALWAYS appears first.

The priority is an uppercase character from A-Z enclosed in parentheses and followed by a space.

This task has a priority:

```
(A) Call Mom
```

These tasks do not have any priorities:

```
Really gotta call Mom (A) @phone @someday
(b) Get back to the boss
(B)->Submit TPS report
```


#### Rule 2: A task's creation date may optionally appear directly after priority and a space.

If there is no priority, the creation date appears first. If the creation date exists, it should be in the format `YYYY-MM-DD`.

These tasks have creation dates:

```
2011-03-02 Document +TodoTxt task format
(A) 2011-03-02 Call Mom
```

This task doesn't have a creation date:

```
(A) Call Mom 2011-03-02
```


#### Rule 3: Contexts and Projects may appear anywhere in the line _after_ priority/prepended date.

- A *context* is preceded by a single space and an at-sign (`@`).
- A *project* is preceded by a single space and a plus-sign (`+`).
- A *project* or *context* contains any non-whitespace character.
- A *task* may have zero, one, or more than one *projects* and *contexts* included in it.

For example, this task is part of the `+Family` and `+PeaceLoveAndHappiness` projects as well as the `@iphone` and `@phone` contexts:

```
(A) Call Mom +Family +PeaceLoveAndHappiness @iphone @phone
```

This task has no contexts in it:

```
Email SoAndSo at soandso@example.com
```

This task has no projects in it:

```
Learn how to add 2+2
```



### Complete Tasks: 2 Format Rules

Two things indicate that a task has been completed.


#### Rule 1: A completed task starts with a lowercase x character (`x`).

If a task starts with an `x` (case-sensitive and lowercase) followed directly by a space, it is marked as complete.

This is a complete task:

```
x 2011-03-03 Call Mom
```

These are not complete tasks.

```
xylophone lesson
X 2012-01-01 Make resolutions
(A) x Find ticket prices
```

We use a lowercase x so that completed tasks sort to the bottom of the task list using standard sort tools.


#### Rule 2: The date of completion appears directly after the x, separated by a space.

For example:

```
x 2011-03-02 2011-03-01 Review Tim's pull request +TodoTxtTouch @github
```

If you’ve prepended the creation date to your task, on completion it will appear directly after the completion date. This is so your completed tasks sort by date using standard sort tools. Many Todo.txt clients discard priority on task completion. To preserve it, use the `key:value` format described below (e.g. `pri:A`)

With the completed date (required), if you've used the prepended date (optional), you can calculate how many days it took to complete a task.



### Additional File Format Definitions

Tool developers may define additional formatting rules for extra metadata.

Developers should use the format `key:value` to define additional metadata (e.g. `due:2010-01-02` as a due date).

Both `key` and `value` must consist of non-whitespace characters, which are not colons. Only one colon separates the `key` and `value`.

