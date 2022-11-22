<script lang="ts">
import type {User} from '../../types/user'
import api from '../../http'
import {auth} from '../../store/auth'
import {navigateTo} from 'svelte-router-spa'
var user: User={email: "",password: ""}

const login=async()=>{
    const res = await api.post(`/login`, {email: user.email,password:user.password});
      if (res) {
        $auth.user=res.data
        $auth.auth=true;
     navigateTo("/")
      
      }
}


</script>

<div class="login-form">
    <h1>Please Login</h1>
      <label><input type="text" bind:value={user.email} placeholder="email" /></label>
    <label><input type="password" bind:value={user.password} placeholder="password" /></label>
  
    <div class="buttons">
      <button
        on:click={login}
        disabled={!user?.email || !user?.password}>Login</button
      >
    </div>
  </div>


  <style>
     input {
    display: block;
    margin: 0 0 0.5em 0;
  }

  button {
    width: 100px;
    height: 40px;
    text-align: center;
  }
  .buttons {
    clear: both;
  }
  .login-form {
    display: flex;
    align-items: center;
    flex-direction: column;
  }
  </style>