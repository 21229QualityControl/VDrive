<script lang="ts">
  import { browser } from "$app/env";
  import { getKeyForCode } from "$lib/key";
  import { room, name } from "$lib/lib";

  type bind = {
    key: string, 
    value: string,
  }
  
  function makeSocket() {
    sock = new WebSocket("wss://vdrive.nv7haven.com/join");
    let pars = {
      room_name: $room,
      name: $name,
    }
    sock.onopen = () => {
      sock.send(JSON.stringify(pars));
    }
    sock.onclose = () => {
      makeSocket();
    }
    sock.onmessage = (ev) => {
      let vals: Record<string, string> = JSON.parse(ev.data);
      for (let k of Object.keys(vals).sort()) {
        binds.push({
          key: k,
          value: vals[k],
        })
      }
      binds = binds;
    }
  }

  let sock: WebSocket;
  let binds: bind[] = [];
  if (browser) {
    makeSocket();
  }

  let pressed: Record<number, boolean> = {};

  function keydown(ev: KeyboardEvent) {
    let key = getKeyForCode(ev.keyCode);
    if (key && !pressed[ev.keyCode]) {
      pressed[ev.keyCode] = true;
      sock.send(key + ":true");
    }
  }

  function keyup(ev: KeyboardEvent) {
    let key = getKeyForCode(ev.keyCode);
    if (key && pressed[ev.keyCode]) {
      pressed[ev.keyCode] = false;
      sock.send(key + ":false");
    }
  }
</script>

<svelte:window on:keydown={keydown} on:keyup={keyup}/>

<h1>Keybinds</h1>
<table class="table table-striped">
  <thead>
    <tr>
      <th scope="col">Key</th>
      <th scope="col">Description</th>
    </tr>
  </thead>
  <tbody>
    {#each binds as bind}
      <tr>
        <td>{bind.key}</td>
        <td>{bind.value}</td>
      </tr>
    {/each}
  </tbody>
</table>
