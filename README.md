# BARE

A tool to manage all your boilerplate from cli and generate files for you !
![Bare Demo](https://github.com/bare-cli/bare/blob/main/assets/bare-demo.gif)

## Trivia

Bare was created by [@madrix01](https://github.com/madrix01) to manage boilerplate that is created by us for every new project.
He called templates as bare-bones and hence the tool `bare`.

bare-li tries to build upon the existing features of bare, try to simplify the interface and also
add support for more templates.


## About different terms/commands of the project

> info 

Info expects a `recipe.json` which should have following meta-data

#### Roadmap

- Make it `bare-agnostic` so that any template can be fetched and used to get information about

> variants

Above term will come in action when u run, `bare use <user-name>/<repo-name> <destination>` or `bare info <user-name>
<repo-name>` 

Variants are possible implementation of a template, For example if you have a template of `React`, it can be in `javascript`
or it can be in `typescript`, to allow multiple variants of a particular framework/language in a single template we provide
flexibility to select your popular choice.

> barebones

These are templates that our tool uses to generate new projects. You can look at `bare-cli/react-template` to get an understanding
of how you can create a template and use it with `bare`


#### Roadmap

- Make creating a template a bit more intuitive



 

## Installation
Currently available for debian based system will release for other systems later
- From `npm` : 

	For Stable build
	```bash
	npm i -g barego
	bare #to check installation
	```

	For Nightly version
	```bash
	npm i -g madrix01/Bare#dev
	baren #to check installation 
	```
- From source :

	Please see [Install](Install.md)


## For more commands and usage follow the [docs](https://bare.surge.sh)


## Roadmap (Final)


- [ ] Make `info` generic for all the template repos
  - Include a message for non-bare templates, that these
  - may not be exact templates that `bare` supports.
  - You can request for templates by creating an issue on `bare-cli/bare` issue tab.


- [ ] Make `spring-boot`, `maven-project` template 
- [ ] Make `bare` robust 
- [ ] Add package in `brew`
- [ ] Command that will create your own bare template from an already existing template
  - an already existing template might have too much baggage and may not offer customizibility
  with `bare minimize <existing-template>` we will strip it of un-necessary stuff
  ask you for the fields that you want to make configurable and create a new template for you.







### More feature are in development 🏗️ 
### Contribute to make it better
- Give your idea and in report bugs in the `issues` tab
