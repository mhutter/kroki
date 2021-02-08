<script>
  import { onMount } from "svelte";

  const defaultState = { teeth: [], lost: false };
  let state = defaultState;
  let socket;

  const onMessage = ({ data }) => {
    console.log(data);
    const { e, p } = JSON.parse(data);
    switch (e) {
      case "update":
        state = p;
        break;

      default:
        console.error("Invalid message", msg);
        break;
    }
  };

  const connect = () => {
    socket = new WebSocket(`wss://localhost:8443/ws`);
    socket.addEventListener("open", (a) => console.log("WS: open", a));
    socket.addEventListener("error", (e) => console.error("WS:", e));

    socket.addEventListener("close", (a) => {
      console.log("WS: close", a);
      state = defaultState;
      setTimeout(connect, 1000);
    });

    socket.addEventListener("message", onMessage);
  };

  const press = (id) => {
    socket.send(JSON.stringify({ e: "press", p: id }));
  };

  onMount(connect);
</script>

<main>
  <h1>Kroki!</h1>
  {#each state.teeth as pressed, i (i)}
    <button disabled={pressed} on:click={() => press(i)}>{i + 1}</button>
  {/each}
  {#if state.lost}
    <h1 color="red">OMG YOU LOST</h1>
  {/if}
</main>
