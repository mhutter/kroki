<script>
  import { setGame, setName } from "../store";

  export let hideForm, name, game;

  const onSubmit = () => {
    setName(name);
    setGame(game);
    hideForm();
  };
</script>

<form on:submit|preventDefault={onSubmit} class="setName">
  <input type="text" placeholder="Name" bind:value={name} required />
  <input type="text" placeholder="Game ID (optional)" bind:value={game} />
  <input type="submit" value="Start!" />
</form>
