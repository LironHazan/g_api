### G-API - NX based polyglot monorepo

---
This is just another research project...
I use the NX tooling a lot in production but with single stack, let's make it interesting!

Exploring the Github gql API and NX go-lang / rust support in a polyglot monorepo.

###### Goals

- Test how affected / general caching works for golang / rust projects tasks.
- How easy it is to generate standard libs in golang/rust using NX.
- Containers integration for local development.
- Lib generators in a non-nx opinionated structure (I don't like the concept of `src/lib`, should be just `/lib`)
- Github workflows in polyglot NX project. 
- State of vite/vitest in NX

** Using Local plugin generators (not the workspace one)

###### Checking out nx-go

`npm install -D @nx-go/nx-go
`

Includes basic generators for app + lib 

`nx g @nx-go/nx-go:app api`

`nx g @nx-go/nx-go:lib api`

Adds the go.mod at the root of the project that should be used for all go projects.

Using the generators will add the test and build tasks executors. 
Missing some [conventions](https://github.com/golang-standards/project-layout), e.g. `/cmd`  `/pkg` `/internal`

The `cmd` directory is where the app main goes 

The `pkg` directory is where you put your public libraries

The `internal` directory is the place to put your private packages

At least when generating a lib we can add the proper layout, for app not sure we need the cmd as in nx its under apps anyway.

Let's extend it:

---

### Main Frontend: gapi app
Simple React app, uses MUI and apoloClient,
You'll need to setup your own GITHUB_TOKEN in order to work with the Github API directly. 

Development server:
`nx run gapi:serve` (TBD)

#### boundaries / scopes / tags / types

See the `enforce-module-boundaries` rule in main eslintrc.json

[enforce-project-boundaries docs](https://nx.dev/core-features/enforce-project-boundaries) 

[tag-multiple-dimensions docs](https://nx.dev/recipes/other/tag-multiple-dimensions) 

### Common commands:

Generate a **publishable** npm lib run with default plugin:

``nx g lib test-lib --publishable --importPath=@my-scope/test-lib``

It will prompt you with a list of choices to choose from (react lib, nodejs lib etc)
** When importing code from another library use alias imports.

** When accessing a file from self lib should not use alias imports, instead use relative paths.

#### Local plugins / generators

Generate a local plugin lib

e.g. `nx g @nrwl/nx-plugin:plugin gapi-plugin
`

Generate a new generator as part of the plugin 

e.g. `nx g @nrwl/nx-plugin:generator lirons-ts-lib-gen --project=gapi-plugin`

Customize the generator as you want.

Use the local generator:

e.g. `nx g @g_api/gapi-plugin:lirons-ts-lib-gen mylib
`

e.g. `nx g @g_api/gapi-plugin:lirons-ts-lib-gen foo`

[Extend community plugins](https://nx.dev/community#plugin-directory) 
