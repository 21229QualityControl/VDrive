<script>

import { goto } from "$app/navigation";

</script>
<h1>VDrive</h1>
<p class="lead">
  VDrive allows you to remotely drive an FTC robot! Learn more at our <a href="https://github.com/Technova-12611/VDrive">GitHub repository</a>!
</p>

<button on:click={() => {goto("/login")}} class="btn btn-lg btn-success">Drive!</button>