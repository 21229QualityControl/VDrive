<script lang="ts">
  import { browser } from "$app/env";
  import { room, name } from "$lib/lib";

  type bind = {
    key: string, 
    value: string,
  }

  let sock: WebSocket;
  let binds: bind[] = [];
  if (browser) {
    sock = new WebSocket("wss://vdrive.nv7haven.com/join");
    let pars = {
      room_name: $room,
      name: $name,
    }
    sock.onopen = () => {
      sock.send(JSON.stringify(pars));
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
</script>

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
