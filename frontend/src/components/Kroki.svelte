<script>
  export let teeth, press;
</script>

<div class="mouth">
  {#each teeth as pressed, i (i)}
    <button class="tooth" disabled={pressed} on:click={() => press(i)}>
      {i + 1}
    </button>
  {/each}
</div>
