<script lang="ts">
    import { onMount } from "svelte";
    import BookForm from "../lib/BookForm.svelte";
    import BookTable from "../lib/BookTable.svelte";
    import type { Book } from "../../types/book";
    import {navigateTo} from 'svelte-router-spa'
  import { auth } from "../../store/auth";
  import api from "../../http/index"
    let books: Book[] = [];
    let book: Book;
    let bookid: number;
  
    let currentRoute
    onMount(async () => {
        if(!$auth.auth) navigateTo("/login")

      getBooks();
    });
  
    const getBooks = async () => {
      const res = await api.get('/books');
      books = await res.data;
      console.log("Books", books);
    };
    let selectedBook: Book = {
      bookid: null,
      name: "",
      price: null,
      publisher: "",
    };
    $: selectedBook = bookid ? books.find((s) => s.bookid === bookid) : null;
   
    function reset_input() {    
     bookid=null //By setting bookid to null the selectedBook will become null
    }
  
    async function updateBook(event) {
      let bk: Book = event.detail.book;
      selectedBook.name = bk.name;
      selectedBook.price = bk.price;
      selectedBook.publisher = bk.publisher;
  
      const res = await api.put(`/books/${selectedBook.bookid}`, {...selectedBook});
      if (res.data) {
        reset_input()
        getBooks();
      }
    }
  
    async function deleteBook(event) {
      const res = await api.delete(`/books/${event.detail.id}`);
      if (res.data) getBooks();
    }
  
    
  
    async function createBook(event) {
        const c_book={...event.detail.book, userId: $auth.user.userid}
      console.log("create book",event.detail.book)
      const res = await api.post(`/books`,  JSON.stringify(c_book));
      console.log(res);
      if (res.data) {
        reset_input()
        getBooks();
      }
    }
  </script>
  
  <h1 class="title">Book Management App</h1>
  
  <BookForm bind:selectedBook on:create={createBook} on:update={updateBook} />
  <BookTable {books} bind:bookid on:delete={deleteBook} />


  <style>
    .title{
      text-align: center;
      margin-top: 20px;
      margin-bottom: 30px;
    }
  </style>
  