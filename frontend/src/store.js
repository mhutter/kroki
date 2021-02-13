import { writable } from "svelte/store";

const keyPlayer = "kroki:player";
const keyGame = "kroki:game";

//
// Helper functions
//
const randomID = () => Math.random().toString(36).substr(2);
const localStorageJSON = (key) => JSON.parse(localStorage.getItem(key));

//
// Stores
//
const player = writable(localStorageJSON(keyPlayer) || { id: randomID() });
const gameID = writable(
  (() => {
    const hash = location.hash;
    if (hash !== "") {
      return hash.substring(1);
    }
    return localStorage.getItem(keyGame);
  })()
);
export { player, gameID };

//
// Public API
//
export const setName = (name) =>
  player.update((p) => {
    p = { ...p, name };
    localStorage.setItem(keyPlayer, JSON.stringify(p));
    return p;
  });

export const setGame = (id) =>
  gameID.update(() => {
    location.hash = id;
    localStorage.setItem(keyGame, id);
    return id;
  });
