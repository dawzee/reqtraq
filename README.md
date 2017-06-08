# Reqtraq


Reqtraq is an open source go library for requirement management, as mandated by
DO-178C.
Reqtraq is designed to stay out of your way. It requires no user interaction for day-to-day tasks.
Instead, it parses documents that are required in the certification process and extracts everything
it needs from there.

Reqtraq has 3 main components:
1. Precommit hook: makes sure the documents have the correct structure and linking:
  * Correctly named
  * Requirements correctly formatted and continuous
  * References exist, Parent requirements exist and not DELETED
  * Required attributes are there and correctly formatted
  * Linkifies documents

2. Prepush hook: exports tasks to desired task management tool (supported for Phabricator; JIRA and others need to be added)

3. Standalone binary:
  * Report generation with filtering
  * Phabricator export
  * Web tool



## How to install Reqtraq
### Dependencies
  * go 1.8+

Install Go according to the instructions [here](https://golang.org/doc/install)


### Installation
```
$ go get github.com/daedaleanai/reqtraq
$ cd $GOPATH/src/github.com/daedaleanai/reqtraq
```

## Using Reqtraq
Reqtraq is tightly intergrated with Git and Lyx (a Latex WYWIG). In order to successfully use Reqtraq your requirement documents need to be written in Lyx, and each requirement delimited with a lyx Note. See the example documents in the `certdocs` directory for some good templates.
Reqtraq uses the Git history to figure out the Git commits associated with a requirement and the Phabricator API to assess the completion status of each requirement.

### Usage examples
#### Getting the next available requirement ID
```
$ cd $GOPATH/src/github.com/daedaleanai/reqtraq/
$ reqtraq nextid certdocs/0-DDLN-212-SDD.lyx
REQ-0-DDLN-SWL-019
```

#### Parse and List requirements
```
$ reqtraq list certdocs/0-DDLN-100-ORD.lyx
Requirement REQ-0-DDLN-SYS-001  Bidirectional tracing.
...
```

#### Report generation
In report tags such as 'Changelists' and 'Problem Reports' will not work if not integrated with a task manager such as Phrabricator etc. (currently supported for Phabricator; JIRA and others need to be added)
```
$ reqtraq reportdown
2017/06/06 22:48:12 Creating ./req-down.html (this may take a while)...
...
```
Filtering:
```
$ reqtraq reportdown --id_filter=".*0-DDLN-SYS.*"
2017/06/06 22:51:23 Creating ./req-down.html (this may take a while)...
2017/06/06 22:51:41 Creating ./req-down-filtered.html (this may take a while)...
```

#### Start the web interface
```
$ reqtraq web :8080
Server started on http://localhost:8080
```

## Getting help
```
$ reqtraq help
```
Shows general help and a list of commands
```
$ reqtraq help <command>
```
Displays help on a specific command.
