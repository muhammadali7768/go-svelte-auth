<script lang="ts">
    import Register from "../lib/Register.svelte";
    import { navigateTo } from "svelte-router-spa";
    import { auth } from "../../store/auth";
    import { onMount } from "svelte";
  
    onMount(async () => {
      if ($auth.auth) navigateTo("/");
    });
  </script>
  
  <Register />
  