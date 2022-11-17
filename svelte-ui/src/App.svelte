<script lang="ts">
  import { onMount } from "svelte";
  import Table from "./lib/Table.svelte";
  import type { Book } from "./types/book";
  let books:Book[] = [];
  let bookid:number;

  let name = "";
  let price = 0;
  let publisher = "";

  onMount(async () => {
  getBooks()
  });

  const getBooks=async()=>{
    const res = await fetch(`http://localhost:8888/api/books`);
    books = await res.json();
  }
  $: selectedBook = bookid ? books.find((s) => s.bookid === bookid) : null;
  $: reset_stock_inputs(selectedBook);

  async function updateBook() {
    selectedBook.name = name;
    selectedBook.price = price;
    selectedBook.publisher = publisher;
    const res = await fetch(
      `http://localhost:8888/api/books/${selectedBook.bookid}`,
      {
        mode: "cors",
        method: "PUT",
        body: JSON.stringify(selectedBook),
      }
    );
    if(res.ok) getBooks();
  }

  async function deleteBook(event) {
    const res = await fetch(
      `http://localhost:8888/api/books/${event.detail.id}`,
      {
        mode: "cors",
        method: "DELETE",
      }
    );
    if(res.ok) getBooks();
   }

  function reset_stock_inputs(book) {
    name = book ? book.name : "";
    price = book ? book.price : "";
    publisher = book ? book.publisher : "";
  }

  async function createBook() {
    const res = await fetch(`http://localhost:8888/api/books`, {
      method: "POST",
      body: JSON.stringify({ name: name, price: price, publisher: publisher }),
    });
    console.log(res);
    if(res.ok) getBooks();
    
  }
</script>
<div class="book-form">
<label><input bind:value={name} placeholder="name" /></label>
<label><input type="number" bind:value={price} placeholder="price" /></label>
<label><input bind:value={publisher} placeholder="publisher" /></label>

<div class="buttons">
  <button on:click={createBook} disabled={!name || !price || !publisher}
    >create</button
  >
  
  <button on:click={updateBook} disabled={!name || !price || !selectedBook}
    >update</button
  >
</div> 
</div>
<Table {books} bind:bookid on:delete={deleteBook} />

<style>
  * {
    font-family: inherit;
    font-size: inherit;
  }

  input {
    display: block;
    margin: 0 0 0.5em 0;
  }

  select {
    float: left;
    margin: 0 1em 1em 0;
    width: 14em;
  }

  button{
    width: 100px;
    height: 40px;
    text-align: center;
  }
  .buttons {
    clear: both;
  }
  .book-form{
    display: flex;
    align-items: center;
    flex-direction: column;
  }
</style>
