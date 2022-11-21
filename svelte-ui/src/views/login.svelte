<script lang="ts">
  import Login from "../lib/Login.svelte";
  import { navigateTo } from "svelte-router-spa";
  import { auth } from "../../store/auth";
  import { onMount } from "svelte";

  onMount(async () => {
    if ($auth.auth) navigateTo("/");
  });
</script>

<Login />
