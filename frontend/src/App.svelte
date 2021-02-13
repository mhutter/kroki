<script>
  import { gameID, player } from "./store";
  import Game from "./components/Game.svelte";
  import Form from "./components/Form.svelte";

  let displayForm = !$player.name || $player.name === "";

  const showForm = () => (displayForm = true);
  const hideForm = () => (displayForm = false);
</script>

<main>
  <h1>Kroki!</h1>
  {#if displayForm}
    <Form {hideForm} name={$player.name} game={$gameID} />
  {:else}
    <Game {...$player} game={$gameID} {showForm} />
    <hr />
    <button on:click={showForm}>Name/Spiel wechseln</button>
  {/if}
</main>
