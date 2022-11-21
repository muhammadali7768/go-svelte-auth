<script lang="ts">
  import Header from "./Header.svelte";
  import { Route } from "svelte-router-spa";
 
  export let currentRoute;
  const params = {};

</script>

<Header />

<section class="section">
  <Route {currentRoute} {params} />
 
</section>
