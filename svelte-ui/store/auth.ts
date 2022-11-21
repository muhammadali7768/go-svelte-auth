import { writable } from 'svelte/store';
import type {User} from '../types/user';



const user:User={name: "",email: "", userid: undefined}

export const  auth =writable({auth: false, user: user});