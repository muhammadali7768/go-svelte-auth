<script lang="ts">
import {auth} from '../../../store/auth'

const logout=()=>{
    $auth={auth: false, user: {}}
}

</script>

<header class="header">
  <a href="#default" class="logo">BookApp</a>
  <div class="header-right">
    <a  href="/">Home</a>
    {#if (!$auth.auth)}
    <a  href="/login">Login</a>
    <a  href="/login">Register</a>
    {:else} 
    <button type="button" href="#" on:click={logout}>Logout</button>
    <div class="user-info">{$auth.user.name}</div>
    {/if}
  </div>
</header>

<style>
  .header {
    overflow: hidden;
    background-color: #f1f1f1;
    padding: 20px 10px;
  }

  .header-right a {
    float: left;
    color: black;
    text-align: center;
    padding: 12px;
    text-decoration: none;
    font-size: 18px;
    line-height: 25px;
    border-radius: 4px;
  }
  .user-info{
    float: left;
    padding: 12px;
    text-align: center;
    line-height: 25px;
  }

  a.logo {
    font-size: 25px;
    font-weight: bold;
  }

  .header a:hover {
    background-color: #ddd;
    color: black;
  }

  .header a.active {
    background-color: dodgerblue;
    color: white;
  }

  .header-right {
    float: right;
  }

  @media screen and (max-width: 500px) {
    .header a {
      float: none;
      display: block;
      text-align: left;
    }

    .header-right {
      float: none;
    }
  }
</style>
