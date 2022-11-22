<script lang="ts">
    import type {User} from '../../types/user'
    import api from '../../http'
    import {auth} from '../../store/auth'
    import {navigateTo} from 'svelte-router-spa'
    var user: User={name: "", email: "",password: ""}
    
    const register=async()=>{
        const res = await api.post(`/register`, {...user});
          if (res) {           
         navigateTo("/login")
          
          }
    }
    
    
    </script>
    
    <div class="register-form">
        <h1>Please register</h1>
        <label><input type="text" bind:value={user.name} placeholder="user name" /></label>
          <label><input type="text" bind:value={user.email} placeholder="email" /></label>
        <label><input type="password" bind:value={user.password} placeholder="password" /></label>
      
        <div class="buttons">
          <button
            on:click={register}
            disabled={!user?.email || !user?.password || !user.name}>Login</button
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
      .register-form {
        display: flex;
        align-items: center;
        flex-direction: column;
      }
      </style>