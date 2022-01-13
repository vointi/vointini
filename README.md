# Vointini

## Requirements: frontend

* Web browser (for example Mozilla Firefox)

## Requirements: backend server

* Operating System: GNU/Linux (x64, ARM, ARM 64 bit)
* Database: [PostgreSQL](https://www.postgresql.org/) version 11+

## Install

### Install PostgreSQL
 
* Install PostgreSQL [docker image](https://hub.docker.com/_/postgres) or OS native version 
* Create `vointini` user with secure password 
* Create `vointini` database with `vointini` user as owner 
* Create database schema from `pg-schema.sql`
* Set up backups (recommended), see [Backup](#Backup)
 
### Configure vointini

* Copy `config.json` to a directory
* Update `config.json` with your created database credentials
* Run server

## Run

Start server:

    ./vointini-server -config /path/to/your/config.json

Go to HTTP address given when server is started with your web browser.
Start logging your data.

## Backup

Since the application uses PostgreSQL you can simply take backups with PostgreSQL's own tools (`pg_dump`)

For example:

    pg_dump --format p --if-exists --clean --no-tablespaces --no-privileges --no-owner --port 5432 --host 127.0.0.1 --username vointini vointini --file backup.sql

Use [`.pgpass`](https://www.postgresql.org/docs/current/libpq-pgpass.html) file to automatically fill in password.

## Developing:

* Frontend is written in [Svelte](https://svelte.dev/) and [TypeScript](https://www.typescriptlang.org/)
* Backend is written in [Go](https://go.dev/)

See the various directories for `README.md` files.

### Creating new page

You should always start with frontend. 
Create the initial listing and form with Svelte first. 
Next add the REST API handler(s) and [DTO(s)](https://en.wikipedia.org/wiki/Data_transfer_object).
As the last thing add service and storage (database) implementation.
Service handles input validation.
Storage only can return internal error(s) which happened during a save or fetch.
Storage never validates data.

Hexagonal architecture is used which means that: 
 
* *REST API* has no clue how *Service* handles things
* *Service* has no clue how *Storage* handles things
* Storage uses internal format(s)
* Many conversion between the different layers must be done

REST API should always use **human** parsable data types (for example tags as *strings* and **not** as *integers*).
Service should use **machine** parsable data (for example tags as *integers* and **not** as *strings*).

Example: generate a new page called `foo`:

* Create `Foo` directory
  * Create `Foo/List.svelte` which lists *Foo* items
  * Create `Foo/Update.svelte` which adds new *Foo* items or updates old ones
* Create `Foo.svelte`:

```sveltehtml
<script lang="ts">
    import {default as List} from "./Foo/List.svelte"
    import Header from "./Header.svelte";
</script>

<main>
    <Header/>
    <List/>
</main>

<style>
</style>
``` 

* Create `foo.ts`:

```typescript
  import Foo from './Foo.svelte'
  
  const app = new Foo({
    target: document.body,
    props: {}
  })
  
  export default app
```

Add `foo` to `rollup.config.js`:

```javascript
const buildthese = [
  "main", "tests", "entries", "etc", "foo",
];
```

So you end up with following directory structure:

    foo.ts
    Foo.svelte
    Foo/List.svelte
    Foo/Update.svelte

Call order is `foo.ts` calls `Foo.svelte` ("main" page for Foo) 
which then calls `Foo/List.svelte` 
which calls `Foo/Update.svelte` from modal if necessary.

You can remove temporarily the other unnecessary pages from `rollup.config.js` file's `buildthese` array while developing to speed up JavaScript compilation.

This is then available as `http://localhost:8080/<language>/foo.html` in the frontend Go HTTP server.

Add *backend* Go REST API handler:

* Create `restfoo.go`

Add DTO to `dto.go`:

```go
type DTOFoo struct {
	Id        int    `json:"id"`
	AddedAt   string `json:"added_at"`
	Name      string `json:"name"`
}
```

Add handlers to `restfoo.go`:

```go
func (restapi restAPI) fooList(w http.ResponseWriter, r *http.Request) {
  panic(`not implemented`)
  // Call Service to get a list of Foo(s)
  // Convert Service format to DTO(s)
  // Return JSON DTO(s)
}

func (restapi restAPI) fooUpdate(w http.ResponseWriter, r *http.Request) {
  id, err := getIntParam(r, `id`)
  if err != nil {
    w.WriteHeader(http.StatusBadRequest)
    panic(err)
    return
  }

  var item DTOFoo
  if err := readStruct(r.Body, &item); err != nil {
    panic(err)
  }
  
  // handle possible conversion(s) to Service format
  // call Service to store item
  // return JSON error(s) or "newId" JSON message 
}
```

Add route(s) to `router.go`:
```go
	// Foo
	router.Get(`/foo`, endpoint.fooList) // List
	router.Post(`/foo/{id}`, endpoint.fooUpdate) // Add new or update existing
```

Next add `FooUpdate` and `FooList` to Service and add validation to field(s). You should look at existing implementations of different items.
Finally add `FooUpdate` and `FooList` to Storage which handles storing and fetching internally. You should look at existing implementations of different items.

### How to release:

* UPX is required when compressing executables

Create a new version:

    ./make_release 1.0.0

Generate packages:

    make release

This creates `release/1.0.0/` directory with compressed files for different CPU architects.
