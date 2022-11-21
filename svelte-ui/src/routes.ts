import Home from './views/home.svelte'
import Login from './views/login.svelte'
import Layout from './lib/layout/Layout.svelte';
const routes = [
    {
      name: '/',
      component: Home,
      layout: Layout
    },
    { name: '/login', component: Login , layout: Layout}
 
  ]
  
  export { routes }