<script lang="ts">
  import { goto } from "$app/navigation";
  import { name, room } from "$lib/lib";

  let formEl: HTMLFormElement;
  var roomVal: string;
  var nameVal: string;

  function connect() {
    // Validation
    let stat = formEl.checkValidity();
    formEl.reportValidity();
    if (!stat) {
      return;
    }

    // Connect
    room.set(roomVal);
    name.set(nameVal);
    goto("/drive");
  }
</script>

<form class="input-group" bind:this={formEl}>
  <input type="text" class="form-control" placeholder="Room Name" required bind:value={roomVal}/>
  <input type="text" class="form-control" placeholder="Username" required bind:value={nameVal}/>
  <button type="submit" class="btn btn-success" on:click|preventDefault={connect}>Connect!</button>
</form>