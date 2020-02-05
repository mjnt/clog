# clog
A CLI for standardizing change logs based on [semver][1].

## Commands

### init [name]

Create a new CHANGELOG.md file
```
clog init
```

Optionally takes a name argument. The containing folder name is used when the name is not specified.
```
clog init myProject
```

Example file contents:
```
# clog Change Log

All notable changes associated with a release will be documented in this file.

This project adheres to [Semantic Versioning](http://semver.org/).

## [Unreleased] - YYYY-MM-DD

### Major

### Minor

### Patch
```
## Objectives
The goal is to port [this project][3] from nodejs to go.

## Format
See basis for format [here][2].

[1]: https://semver.org/
[2]: https://majgis.github.io/change-log/
[3]: https://github.com/majgis/change-log
