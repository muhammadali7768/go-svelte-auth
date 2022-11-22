import Home from './views/home.svelte'
import Login from './views/login.svelte'
import Layout from './lib/layout/Layout.svelte';
import Register from './lib/Register.svelte';
const routes = [
    {
      name: '/',
      component: Home,
      layout: Layout
    },
    { name: '/login', component: Login , layout: Layout},
    { name: '/register', component: Register , layout: Layout}
 
  ]
  
  export { routes }