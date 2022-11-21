<script lang="ts">
  import { createEventDispatcher } from "svelte";
  import type { Book } from "../../types/book";
  export let books: Book[];
  export let bookid: number;
  const dispatch = createEventDispatcher();

  const setBookId = (id) => {
    bookid = id;
  };

  const setDeletedBookId = (id) => {
    dispatch("delete", {
      id: id,
    });
  };
</script>

<table>
  <tr>
    <th>Name</th>
    <th>Price</th>
    <th>Publisher</th>
    <th>Action</th>
  </tr>
  {#if !books} <tr><td colspan="4">No Books added yet</td></tr>
  {:else}
  {#each books as book}
    <tr>
      <td>{book.name}</td>
      <td>{book.price}</td>
      <td>{book.publisher}</td>
      <td>
        <button type="button" on:click={() => setBookId(book.bookid)}
          >Edit</button
        >
        <button type="button" on:click={() => setDeletedBookId(book.bookid)}
          >Delete</button
        >
      </td>
    </tr>
  {/each}
  {/if}
</table>

<style>
  table {
    border-collapse: collapse;
    width: 90%;
    margin-left: auto;
    margin-right: auto;
    margin-top: 20px;
  }

  th,
  td {
    padding: 8px;
    text-align: left;
    border-bottom: 1px solid #ddd;
  }
  button {
    height: 30px;
    padding-top: 5px;
    padding-bottom: 5px;
    font-size: 14px;
  }
</style>
