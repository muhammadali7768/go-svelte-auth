<script lang="ts">
  import { onMount } from "svelte";
  import BookForm from "./lib/BookForm.svelte";
  import BookTable from "./lib/BookTable.svelte";
  import type { Book } from "./types/book";

  const BASE_URL = "http://localhost:8888/api";
  let books: Book[] = [];
  let book: Book;
  let bookid: number;

  onMount(async () => {
    getBooks();
  });

  const getBooks = async () => {
    const res = await fetch(`${BASE_URL}/books`);
    books = await res.json();
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

    const res = await fetch(`${BASE_URL}/books/${selectedBook.bookid}`, {
      mode: "cors",
      method: "PUT",
      body: JSON.stringify(selectedBook),
    });
    if (res.ok) {
      reset_input()
      getBooks();
    }
  }

  async function deleteBook(event) {
    const res = await fetch(`${BASE_URL}/books/${event.detail.id}`, {
      mode: "cors",
      method: "DELETE",
    });
    if (res.ok) getBooks();
  }

  

  async function createBook() {
    const res = await fetch(`${BASE_URL}/books`, {
      method: "POST",
      body: JSON.stringify(book),
    });
    console.log(res);
    if (res.ok) {
      reset_input()
      getBooks();
    }
  }
</script>

<h1>Book Management App </h1>

<BookForm bind:selectedBook on:create={createBook} on:update={updateBook} />
<BookTable {books} bind:bookid on:delete={deleteBook} />
