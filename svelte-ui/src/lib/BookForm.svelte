<script lang="ts">
  import { createEventDispatcher } from "svelte";
  import type { Book } from "../../types/book";

  export let selectedBook;
  let book: Book;
  const dispatch = createEventDispatcher();
  $: book =
    selectedBook?.bookid != null
      ? selectedBook
      : { name: "", price: null, publisher: "" };

  const createBook = () => {
    dispatch("create",{book:book});
  };

  const updateBook = () => {
    dispatch("update", { book: book });
  };
</script>

<div class="book-form">
  <label><input bind:value={book.name} placeholder="name" /></label>
  <label
    ><input type="number" bind:value={book.price} placeholder="price" /></label
  >
  <label><input bind:value={book.publisher} placeholder="publisher" /></label>

  <div class="buttons">
    <button
      on:click={createBook}
      disabled={!book?.name || !book?.price || !book?.publisher}>create</button
    >

    <button on:click={updateBook} disabled={!selectedBook}>update</button>
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
  .book-form {
    display: flex;
    align-items: center;
    flex-direction: column;
  }
</style>
