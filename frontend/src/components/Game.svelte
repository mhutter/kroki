<script>
  import { onDestroy, onMount } from "svelte";
  import { gameID, setGame } from "../store";
  import Kroki from "./Kroki.svelte";

  let players = [];
  let teeth = [];
  let lost;
  // const defaultState = { teeth: [], lost: false, players: [] };

  export let id, name, showForm;
  export let game = "";
  let socket;

  const onMessage = ({ data }) => {
    console.log(data);
    const { e, p } = JSON.parse(data);
    switch (e) {
      case "update":
        players = p.players;
        teeth = p.teeth;
        lost = p.lost;
        setGame(p.id);
        break;

      case "leave":
        showForm();
        break;

      default:
        console.error("Invalid message", e);
        break;
    }
  };

  const sendJSON = (payload) => socket.send(JSON.stringify(payload));

  const connect = () => {
    console.log("mounting...");
    socket = new WebSocket(`wss://localhost:8443/ws`);

    socket.onopen = () => {
      sendJSON({ e: "setName", p: name });
      sendJSON({ e: "setPlayerID", p: id });
      sendJSON({ e: "joinGame", p: game });
    };

    socket.onerror = (e) => console.error("WS:", e);

    socket.onclose = (a) => {
      console.log("WS: close", a);
      setTimeout(connect, 1000);
    };

    socket.onmessage = onMessage;
  };

  const disconnect = () => {
    // prevent reconnection
    socket.onclose = undefined;
    socket.close(1000);
  };

  const press = (id) => sendJSON({ e: "press", p: id });
  const restart = () => sendJSON({ e: "restart" });

  const playerName = (id) => players.find((p) => p.id === id).name;

  onMount(connect);
  onDestroy(disconnect);
</script>

<div class="game">
  <p>ID: <code>{$gameID}</code></p>
  <Kroki {teeth} {press} />
  {#if lost}
    <p>{playerName(lost)} het verlore :)</p>
    <button on:click={restart}>Neu starten</button>
  {/if}

  <h2>Mitspieler</h2>
  <ul class="players">
    {#each players as { id, name } (id)}
      <li class:lost={id === lost}>{name}</li>
    {/each}
  </ul>
</div>
