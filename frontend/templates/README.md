# Svelte frontend templates

* Files in [public](public) directory are embedded into Go executable
  * See [rollup.config.js](rollup.config.js) file's variable `buildthese` to see which pages are generated 
* Served by Go frontend [server](../server)
* User interacts with Go backend [REST API](../../backend/restapi) used by [Svelte templates](src)
